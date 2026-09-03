# gbf: the GObugfree command-line client

`gbf` lets people and AI agents work the GObugfree bug bounty platform from a
terminal. It logs in with the OAuth 2.0 device flow as a platform user, then
drives the platform through that user's live OpenAPI document: whatever the
document lists is exactly what the user may call, whatever their role. The
agent skill that teaches an AI agent how to use it ships inside the binary.

## Install

Download `gbf_cli_<os>_<arch>` for your platform from the
[releases](https://github.com/gobugfree/cli/releases/latest) (Linux and macOS
on amd64 and arm64, Windows on amd64 as `.exe`), verify it against
`checksums.txt`, make it executable, and put it on your PATH as `gbf`:

```bash
curl -fsSLO https://github.com/gobugfree/cli/releases/latest/download/gbf_cli_linux_amd64
install -m 0755 gbf_cli_linux_amd64 ~/.local/bin/gbf
```

Go users can build from source instead. The repository is internal, so an
authenticated organization member and `GOPRIVATE` are needed:

```bash
export GOPRIVATE=github.com/gobugfree
go install github.com/gobugfree/cli/v3/cmd/gbf@latest
```

## Use

```bash
gbf login --app-url https://<platform> --issuer-url https://<keycloak>/realms/gbf
gbf whoami                                   # user, groups (areas), token expiry
gbf ops --tag engagements                    # operations, by tag or --search TEXT
gbf describe postEngagementsRulesAccept      # parameters, body fields, responses, example
gbf run postEngagementsRulesAccept --param org_slug=acme --param engagement_slug=web-app --dry-run
gbf run postReportsDraft --form title=... --form severity=high
gbf run getEngagements --query kind=pentest
gbf call GET /reports                        # any request, for what run does not cover
gbf schema Error                             # one component schema
gbf openapi --refresh                        # the whole document
gbf skill --install                          # install the agent skill into ~/.claude/skills/gbf
```

`login` prints the device-flow URL and waits for the approval; the session
(URLs, token, and the cached document) is stored under `$XDG_CONFIG_HOME/gbf`
(`~/.config/gbf`) with mode 0600, `GBF_SESSION_FILE` overrides the file, and
the token is refreshed on use. `run` fills the path from `--param`, the query
from `--query`, and sends the body the way the operation declares it, form
fields or `--json-body`; a missing parameter or the wrong kind of body is
reported with what the operation needs, and `--dry-run` prints the exact
request instead of sending it, which is how an agent shows its user a
binding action before performing it. Output is JSON whenever stdout is not
a terminal, or with `--json` as the first argument; errors keep the
platform's JSON error body (with `requestId`) on stdout and exit non-zero.
The document is cached per session and refreshed with `--refresh` or
automatically when an operation id is unknown.

Behind Cloudflare Access (Zero Trust), export a service token as
`CF_ACCESS_CLIENT_ID` and `CF_ACCESS_CLIENT_SECRET`: every request gbf makes,
to the platform and to the identity provider, then carries the
`CF-Access-Client-Id` and `CF-Access-Client-Secret` headers. The values stay
in the environment; nothing of them is written to the session file.

## The agent skill

`SKILL.md` is an installable agent skill (frontmatter with the name `gbf`
and a description, then the instructions). `gbf skill` prints it and
`gbf skill --install` writes it to `~/.claude/skills/gbf/SKILL.md` (or
`--dir DIR/gbf/SKILL.md`); it needs no login and no network. It teaches the
method (login, `whoami`, `ops`, `describe`, `run`, `--dry-run` before binding
actions) and the ground rules, and names a few operations by id as examples;
a test checks those ids against `openapi.json`.

## How this repository is fed

`openapi.json` is the platform's full OpenAPI document. The platform's
`OpenAPI delivery` workflow exports it from the live route table on every
push to its `main`, regenerates the `api` package here with `make generate`,
runs `make check`, commits the result to `main`, and, after a platform version
bump, tags this repository `v<platform version>`. The `Release` workflow then
publishes the binaries. Nothing here is edited by hand except the sources:
`cmd/gbf`, `api/agent.go` and its test (the OAuth 2.0 device flow that
openapi-generator does not provide, protected by `api/.openapi-generator-ignore`),
`SKILL.md`, and the configuration.

## Develop

```bash
make generate    # regenerate api from openapi.json (Docker, official openapi-generator image)
make check       # gofmt, go vet, go test
make build       # .build/gbf
make install     # put it on the PATH (BINDIR, default $(go env GOPATH)/bin)
make build-all   # the release binaries into dist/
```

CI runs `make drift`, which fails when the committed `api` package differs
from a fresh generation, so regenerate and commit after changing the
generator configuration. To try an API change before the platform delivers
it, copy the platform's `.build/openapi.json` (from `make openapi` there)
over `openapi.json` and run `make generate check`.
