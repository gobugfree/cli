# gbf: the GObugfree command-line client

`gbf` lets security researchers and their AI agents work the GObugfree bug
bounty platform from a terminal. It logs in with the OAuth 2.0 device flow
as the researcher, then drives the platform through that user's live OpenAPI
document: whatever the document lists is exactly what the user may call.
The agent skill that teaches an AI agent how to use it ships inside the
binary.

## Install

Every [release](https://github.com/gobugfree/cli/releases/latest) carries
one binary per platform plus `checksums.txt`. The repository is internal,
so downloading needs a signed-in GitHub account with access to it.

| Platform | File |
|---|---|
| Linux, x86-64 | `gbf_cli_linux_amd64` |
| Linux, ARM64 | `gbf_cli_linux_arm64` |
| macOS, Apple silicon | `gbf_cli_darwin_arm64` |
| macOS, Intel | `gbf_cli_darwin_amd64` |
| Windows, x86-64 | `gbf_cli_windows_amd64.exe` |

Download the file for your platform, verify it against `checksums.txt`,
make it executable, and put it on your PATH as `gbf`:

```bash
curl -fsSLO https://github.com/gobugfree/cli/releases/latest/download/gbf_cli_linux_amd64
sha256sum -c --ignore-missing <(curl -fsSL https://github.com/gobugfree/cli/releases/latest/download/checksums.txt)
install -m 0755 gbf_cli_linux_amd64 ~/.local/bin/gbf
```

On macOS use `shasum -a 256 -c` for the check. The binary is not notarized:
if you downloaded it with a browser, run
`xattr -d com.apple.quarantine <path-to-gbf>` once, otherwise macOS refuses
to start it (`curl` does not set that flag). On Windows, put `gbf.exe` in a
directory on your PATH.

Go users can build from source instead:

```bash
GOPRIVATE=github.com/gobugfree go install github.com/gobugfree/cli/v3/cmd/gbf@latest
```

## Use

```bash
gbf login                                    # app.gobugfree.com; --app-url and --issuer-url for another platform
gbf whoami                                   # user, client, token expiry
gbf ops --tag reports                        # operations, by tag or --search TEXT
gbf describe postEngagementsRulesAccept      # parameters, body fields, responses, example
gbf run postEngagementsRulesAccept --param org_slug=acme --param engagement_slug=web-app --dry-run
gbf run getReports --query status=validated
gbf call GET /reports                        # any request, for what run does not cover
gbf schema Error                             # one component schema
gbf openapi --refresh                        # the whole document
```

`login` prints the device-flow URL and waits for the approval; the session
(URLs, token, and the cached document) is stored under `$XDG_CONFIG_HOME/gbf`
(`~/.config/gbf`) with mode 0600, `GBF_SESSION_FILE` overrides the file, and
the token is refreshed on use. `run` fills the path from `--param`, the query
from `--query`, and sends the body the way the operation declares it, form
fields or `--json-body`; `--dry-run` prints the exact request instead of
sending it. Output is JSON whenever stdout is not a terminal, or with
`--json` as the first argument; errors keep the platform's JSON error body
(with `requestId`) on stdout and exit non-zero.

Every `GBF_HEADER_*` environment variable is sent as a request header, with
underscores as dashes, on every request including the device flow:

```bash
export GBF_HEADER_CF_ACCESS_CLIENT_ID=...      # CF-Access-Client-Id
export GBF_HEADER_CF_ACCESS_CLIENT_SECRET=...  # CF-Access-Client-Secret
export GBF_HEADER_X_TENANT=acme                # X-Tenant
```

That is how a platform behind Cloudflare Access (Zero Trust), an API
gateway, or any other header-checking proxy is reached. An empty value
configures no header, and nothing of these headers is written to the
session file.

## The agent skill

`SKILL.md` teaches an AI agent the method: login, `whoami`, `ops`,
`describe`, `run`, and `--dry-run` before every binding action. `gbf skill`
prints it; `gbf skill --install` writes it to the agent's user-level skills
directory, no login or network needed:

| Agent | Command | Installs to |
|---|---|---|
| Claude Code (default) | `gbf skill --install` | `~/.claude/skills/gbf/SKILL.md` |
| Cursor | `gbf skill --install --agent cursor` | `~/.cursor/skills/gbf/SKILL.md` |
| Codex | `gbf skill --install --agent codex` | `~/.agents/skills/gbf/SKILL.md` |
| LM Studio | `gbf skill --install --agent lmstudio` | `~/.lmstudio/skills/gbf/SKILL.md` |

`--dir DIR` installs to `DIR/gbf/SKILL.md` for any other agent or for a
project-level skills directory. Re-run the install after updating `gbf`; the
skill is stamped with the version and a content hash.

## Shell completion

Tab completes commands, flags, operation ids, path and query parameters,
form fields, enum values, request paths, tags, and schema names, all from
the document cached at login, so it is instant and never touches the
network. Install it once for your shell:

```bash
gbf completion bash --install    # or zsh, fish
```

That writes the script where the shell loads it by itself: bash-completion's
user directory (`~/.local/share/bash-completion/completions/gbf`, picked up
on the first tab in every new shell), `~/.local/share/zsh/site-functions/_gbf`
(add that directory to `fpath` before `compinit`), or
`~/.config/fish/completions/gbf.fish`. Without `--install` the script is
printed, for `eval "$(gbf completion bash)"` in a shell without
bash-completion, or, in PowerShell,
`gbf completion powershell | Out-String | Invoke-Expression` in `$PROFILE`.

```
$ gbf run postRetestsVerdict --form <TAB>
notes=    outcome=
$ gbf run postRetestsVerdict --form outcome=<TAB>
passed  failed
```

## Develop

`openapi.json` is the researcher-scoped OpenAPI document the platform
delivers; the platform's `OpenAPI delivery` workflow refreshes it on every
push to the platform's `main`, regenerates the `api` package, commits, and
tags this repository `v<platform version>` after a platform version bump.
The `Release` workflow then publishes the binaries. Hand-written sources
are `cmd/gbf`, `api/agent.go` with its test (the device flow), `SKILL.md`,
and the configuration; everything else under `api` is generated.

```bash
make generate    # regenerate api from openapi.json (Docker, official openapi-generator image)
make check       # gofmt, go vet, go test
make build       # .build/gbf
make install     # put it on the PATH (BINDIR, default $(go env GOPATH)/bin)
make build-all   # the release binaries into dist/
```

CI runs `make drift`, which fails when the committed `api` package differs
from a fresh generation. A test checks that every operation id the skill
names exists in `openapi.json`.
