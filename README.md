# gbf: the GObugfree command-line client

`gbf` lets people and AI agents work the GObugfree bug bounty platform from a
terminal. It logs in with the OAuth 2.0 device flow as a platform user, then
drives the platform through that user's live OpenAPI document: whatever the
document lists is exactly what the user may call, whatever their role. The
agent skill that teaches an AI agent how to use it ships inside the binary.

## Install

Every release on the [releases page](https://github.com/gobugfree/cli/releases/latest)
carries one binary per platform plus `checksums.txt`:

| Platform | File |
|---|---|
| Linux, x86-64 | `gbf_cli_linux_amd64` |
| Linux, ARM64 | `gbf_cli_linux_arm64` |
| macOS, Apple silicon | `gbf_cli_darwin_arm64` |
| macOS, Intel | `gbf_cli_darwin_amd64` |
| Windows, x86-64 | `gbf_cli_windows_amd64.exe` |

Verify the download against `checksums.txt`, then put it on your PATH under
the name `gbf`. The repository is internal, so downloading needs a signed-in
organization member (the commands below assume the browser or `gh` did
that; `gh release download -R gobugfree/cli -p 'gbf_cli_*'` works too).

### Linux

```bash
curl -fsSLO https://github.com/gobugfree/cli/releases/latest/download/gbf_cli_linux_amd64
sha256sum -c --ignore-missing <(curl -fsSL https://github.com/gobugfree/cli/releases/latest/download/checksums.txt)
install -m 0755 gbf_cli_linux_amd64 ~/.local/bin/gbf
```

### macOS

```bash
curl -fsSLO https://github.com/gobugfree/cli/releases/latest/download/gbf_cli_darwin_arm64   # Intel: gbf_cli_darwin_amd64
shasum -a 256 -c --ignore-missing <(curl -fsSL https://github.com/gobugfree/cli/releases/latest/download/checksums.txt)
install -m 0755 gbf_cli_darwin_arm64 /usr/local/bin/gbf
```

The binary is not notarized. A file downloaded with a browser carries the
quarantine flag and macOS refuses to run it; `curl` does not set the flag.
If you downloaded it with a browser, clear the flag once:

```bash
xattr -d com.apple.quarantine /usr/local/bin/gbf
```

### Windows

In PowerShell:

```powershell
$dir = "$env:LOCALAPPDATA\Programs\gbf"
New-Item -ItemType Directory -Force $dir | Out-Null
Invoke-WebRequest https://github.com/gobugfree/cli/releases/latest/download/gbf_cli_windows_amd64.exe -OutFile "$dir\gbf.exe"
(Get-FileHash "$dir\gbf.exe" -Algorithm SHA256).Hash   # compare with checksums.txt
[Environment]::SetEnvironmentVariable("Path", "$env:Path;$dir", "User")
```

Open a new terminal afterwards so the PATH change applies. The session file
lives in `%USERPROFILE%\.config\gbf\session.json` unless
`XDG_CONFIG_HOME` or `GBF_SESSION_FILE` says otherwise.

### From source (Go)

Go users can build from source instead. An authenticated organization
member and `GOPRIVATE` are needed:

```bash
export GOPRIVATE=github.com/gobugfree
go install github.com/gobugfree/cli/v3/cmd/gbf@latest
```

The binary lands in `$(go env GOBIN)`, or `$(go env GOPATH)/bin` when `GOBIN`
is unset; make sure that directory is on your PATH.

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
