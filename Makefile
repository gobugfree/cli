.PHONY: generate check drift build build-all install clean

# The api package is generated from openapi.json, the document the platform
# delivers (workflow "OpenAPI delivery" in gobugfree/platform). The generator
# runs from its official image, pinned by tag and digest.
OPENAPI_GENERATOR_IMAGE := openapitools/openapi-generator-cli:v7.25.0@sha256:2ab0a9680222de65dc9d3baf861aa02b99e1b80c211d8221ebf3ae8f8a102524
# Rootful Docker needs -u so the generated files belong to the caller;
# rootless Docker already maps the container root to the caller, and -u
# would make the bind mount read-only for it.
DOCKER_USER = $(shell docker info --format '{{.SecurityOptions}}' 2>/dev/null | grep -q rootless || echo -u $$(id -u):$$(id -g))

# a release build carries its tag; anything else is dev
VERSION ?= $(shell git describe --tags --exact-match 2>/dev/null || echo dev)
LDFLAGS := -s -w -X main.version=$(VERSION)
TARGETS := linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64
BINDIR ?= $(shell go env GOPATH)/bin

# regenerate api from openapi.json; files of removed routes go first so
# nothing stale survives
generate:
	cd api && if [ -f .openapi-generator/FILES ]; then xargs rm -f < .openapi-generator/FILES; fi
	docker run --rm $(DOCKER_USER) -v "$(CURDIR):/local" $(OPENAPI_GENERATOR_IMAGE) \
		generate -g go -i /local/openapi.json -o /local/api -c /local/openapi-generator.yaml
	go mod tidy

check:
	test -z "$$(gofmt -l cmd *.go api/agent.go api/agent_test.go)"
	go vet ./...
	go test ./...

# CI: the committed api package must equal a fresh generation
drift: generate
	git add -A
	git diff --cached --exit-code --stat

build:
	CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o .build/gbf ./cmd/gbf

# the release binaries, one per target, plus their checksums
build-all:
	rm -rf dist && mkdir -p dist
	for target in $(TARGETS); do \
		os=$${target%/*}; arch=$${target#*/}; \
		name="gbf_cli_$${os}_$${arch}"; [ "$$os" = windows ] && name="$$name.exe"; \
		CGO_ENABLED=0 GOOS=$$os GOARCH=$$arch go build -trimpath -ldflags "$(LDFLAGS)" -o "dist/$$name" ./cmd/gbf || exit 1; \
	done
	cd dist && sha256sum -- * > checksums.txt && cat checksums.txt

install: build
	install -m 0755 .build/gbf "$(BINDIR)/gbf"
	@echo "gbf installed at $(BINDIR)/gbf"

clean:
	rm -rf .build dist
