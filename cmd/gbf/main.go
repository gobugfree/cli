// Command gbf is the GObugfree command-line client for people and agents.
// It logs in with the OAuth 2.0 device flow and then drives the platform
// through its live OpenAPI document: the document is scoped to the
// logged-in user, so what it lists is exactly what the user may call.
package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"time"

	gbfcli "github.com/gobugfree/cli/v3"
	"github.com/gobugfree/cli/v3/api"
	"golang.org/x/oauth2"
)

// version is set at build time (-X main.version=...); a go install build
// reports its module version instead.
var version = "dev"

func init() {
	if version != "dev" {
		return
	}
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "" && info.Main.Version != "(devel)" {
		version = info.Main.Version
	}
}

// userAgent identifies the command to the platform, on the generated
// client's calls and on generic ones alike.
func userAgent() string {
	return "gbf-cli/" + version + "/go"
}

const usage = `gbf: GObugfree command-line client for people and agents

Usage:
  gbf [--json] COMMAND ...

  gbf login [--app-url URL] [--issuer-url URL] [--client-id ID]
  gbf logout
  gbf status
  gbf whoami
  gbf ops [--tag TAG] [--search TEXT] [--refresh]
  gbf describe OPERATION_ID
  gbf run OPERATION_ID [--param KEY=VALUE]... [--query KEY=VALUE]...
                       [--form KEY=VALUE]... [--json-body BODY] [--dry-run]
  gbf call METHOD PATH [--json-body BODY] [--form KEY=VALUE]...
  gbf schema NAME
  gbf openapi [--refresh]
  gbf skill [--install] [--dir DIR]
  gbf version

The platform's live OpenAPI document, scoped to the logged-in user, is the
source of truth: ops lists the operations this user may call, describe shows
one operation's parameters and body, run invokes it by id, and call sends
any request. Output is JSON when stdout is not a terminal or with --json as
the first argument. login starts the device flow: open the printed URL, sign
in, approve, and gbf continues; the URLs default to GBF_APP_URL and
GBF_ISSUER_URL. skill prints the agent skill that ships with gbf (an
installable SKILL.md); --install writes it to DIR/gbf/SKILL.md, DIR
defaulting to ~/.claude/skills. Behind Cloudflare Access (Zero Trust), set
CF_ACCESS_CLIENT_ID and CF_ACCESS_CLIENT_SECRET to a service token: every
request then carries the CF-Access-Client-Id and CF-Access-Client-Secret
headers. The session (URLs, token, cached document) lives in
$XDG_CONFIG_HOME/gbf (GBF_SESSION_FILE overrides the session file).
`

type cli struct {
	out  io.Writer
	json bool
}

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "gbf:", err)
		os.Exit(1)
	}
}

func run(args []string, out io.Writer) error {
	c := &cli{out: out, json: stdoutIsPipe(out)}
	if len(args) > 0 && args[0] == "--json" {
		c.json = true
		args = args[1:]
	}
	if len(args) == 0 {
		fmt.Fprint(out, usage)
		return nil
	}
	if err := checkAccessEnv(); err != nil {
		return err
	}
	ctx := context.Background()
	switch args[0] {
	case "login":
		return c.login(ctx, args[1:])
	case "logout":
		return c.logout()
	case "status":
		return c.status()
	case "whoami":
		return c.whoami()
	case "ops":
		return c.ops(ctx, args[1:])
	case "describe":
		return c.describe(ctx, args[1:])
	case "run":
		return c.runOperation(ctx, args[1:])
	case "call":
		return c.call(ctx, args[1:])
	case "schema":
		return c.schema(ctx, args[1:])
	case "openapi":
		return c.openapi(ctx, args[1:])
	case "skill":
		return c.skill(args[1:])
	case "version":
		fmt.Fprintln(out, version)
		return nil
	case "help", "-h", "--help":
		fmt.Fprint(out, usage)
		return nil
	default:
		return fmt.Errorf("unknown command %q\n\n%s", args[0], usage)
	}
}

// stdoutIsPipe reports whether out is the real stdout and it is not a
// terminal, in which case an agent or a pipe is reading and JSON is the
// better default.
func stdoutIsPipe(out io.Writer) bool {
	if out != io.Writer(os.Stdout) {
		return false
	}
	info, err := os.Stdout.Stat()
	return err == nil && info.Mode()&os.ModeCharDevice == 0
}

func (c *cli) login(ctx context.Context, args []string) error {
	flags := newFlags("login", c.out)
	appURL := flags.String("app-url", os.Getenv("GBF_APP_URL"), "platform base URL")
	issuerURL := flags.String("issuer-url", os.Getenv("GBF_ISSUER_URL"), "Keycloak realm URL")
	clientID := flags.String("client-id", os.Getenv("GBF_CLIENT_ID"), "device-flow client id (default agent)")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *appURL == "" || *issuerURL == "" {
		return errors.New("login needs --app-url and --issuer-url (or GBF_APP_URL and GBF_ISSUER_URL)")
	}

	s := &session{AppURL: strings.TrimSuffix(*appURL, "/"), IssuerURL: strings.TrimSuffix(*issuerURL, "/"), ClientID: *clientID}
	token, err := s.agent().Login(ctx, func(r *oauth2.DeviceAuthResponse) {
		fmt.Fprintf(c.out, "Open %s\nSign in and approve the login (code %s). Waiting...\n", r.VerificationURIComplete, r.UserCode)
	})
	if err != nil {
		return err
	}
	s.Token = token
	if err := s.save(); err != nil {
		return err
	}
	_ = os.Remove(documentPath())
	fmt.Fprintf(c.out, "Logged in to %s; session stored in %s\n", s.AppURL, sessionPath())
	return nil
}

func (c *cli) logout() error {
	for _, path := range []string{sessionPath(), documentPath()} {
		if err := os.Remove(path); err != nil && !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}
	fmt.Fprintln(c.out, "Logged out")
	return nil
}

func (c *cli) status() error {
	s, err := loadSession()
	if err != nil {
		return err
	}
	clientID := s.agent().ClientID
	if clientID == "" {
		clientID = api.DefaultClientID
	}
	state := "no expiry recorded"
	switch {
	case s.Token.Expiry.IsZero():
	case s.Token.Valid():
		state = "valid for " + time.Until(s.Token.Expiry).Round(time.Second).String()
	case s.Token.RefreshToken != "":
		state = "expired, refreshes on the next call"
	default:
		state = "expired, run gbf login"
	}
	if c.json {
		return c.printJSON(map[string]any{
			"appUrl": s.AppURL, "issuerUrl": s.IssuerURL, "clientId": clientID,
			"tokenValid": s.Token.Valid(), "tokenExpiry": s.Token.Expiry, "token": state,
		})
	}
	fmt.Fprintf(c.out, "Platform: %s\nIssuer:   %s\nClient:   %s\nToken:    %s\n", s.AppURL, s.IssuerURL, clientID, state)
	return nil
}

// skill prints the agent skill that ships with gbf, or installs it where the
// agent loads skills from. It needs no session and no network.
func (c *cli) skill(args []string) error {
	flags := newFlags("skill", c.out)
	install := flags.Bool("install", false, "write the skill to DIR/gbf/SKILL.md instead of printing it")
	dir := flags.String("dir", "", "skills directory for --install (default ~/.claude/skills)")
	if err := flags.Parse(args); err != nil {
		return err
	}
	v := gbfcli.SkillVersion(version)
	if !*install {
		if c.json {
			return c.printJSON(map[string]string{"name": gbfcli.SkillName, "version": v, "content": gbfcli.Skill})
		}
		_, err := io.WriteString(c.out, gbfcli.Skill)
		return err
	}
	if *dir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		*dir = filepath.Join(home, ".claude", "skills")
	}
	target := filepath.Join(*dir, gbfcli.SkillName, "SKILL.md")
	if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(target, []byte(gbfcli.Skill), 0o644); err != nil {
		return err
	}
	fmt.Fprintf(c.out, "Installed the skill (version %s) at %s\n", v, target)
	return nil
}
