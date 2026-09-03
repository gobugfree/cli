package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"golang.org/x/oauth2"
)

const testDocument = `{
  "openapi": "3.1.0",
  "info": {"title": "Platform", "version": "3.x"},
  "servers": [{"url": "https://app.example"}],
  "paths": {
    "/engagements": {"get": {
      "operationId": "getEngagements", "tags": ["engagements"],
      "summary": "Engagement directory",
      "parameters": [{"name": "kind", "in": "query", "schema": {"type": "string", "enum": ["bug_bounty", "pentest"]}}],
      "responses": {"200": {"description": "One page", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/EngagementListItemList"}}}}}
    }},
    "/engagements/{org_slug}/{engagement_slug}/submit/draft": {"post": {
      "operationId": "postEngagementsSubmitDraft", "tags": ["engagements"],
      "summary": "Save a private draft report",
      "parameters": [
        {"name": "org_slug", "in": "path", "required": true, "schema": {"type": "string"}},
        {"name": "engagement_slug", "in": "path", "required": true, "schema": {"type": "string"}}
      ],
      "requestBody": {"required": true, "content": {"application/x-www-form-urlencoded": {"schema": {
        "type": "object", "required": ["title"],
        "properties": {"title": {"type": "string", "maxLength": 200, "description": "Report title"}, "origin": {"type": "string", "enum": ["human", "ai_agent"]}}
      }}}},
      "responses": {"200": {"description": "Saved"}, "default": {"description": "Error", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/Error"}}}}}
    }}
  },
  "components": {"schemas": {
    "Error": {"type": "object", "properties": {"status": {"type": "integer"}}, "required": ["status"]},
    "EngagementListItemList": {"type": "object", "properties": {"items": {"type": "array", "items": {"type": "object"}}}, "required": ["items"]}
  }}
}`

// fakePlatform serves the document and echoes every other request.
type fakePlatform struct {
	server    *httptest.Server
	requests  []*http.Request
	bodies    []string
	documents int
}

func newFakePlatform(t *testing.T) *fakePlatform {
	t.Helper()
	f := &fakePlatform{}
	f.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f.requests = append(f.requests, r)
		data, _ := readAll(r)
		f.bodies = append(f.bodies, data)
		switch {
		case r.URL.Path == "/openapi.json":
			f.documents++
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(testDocument))
		case r.URL.Path == "/nope":
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"status":404}`))
		default:
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{"items": []any{}, "path": r.URL.Path})
		}
	}))
	t.Cleanup(f.server.Close)
	return f
}

func (f *fakePlatform) last() *http.Request { return f.requests[len(f.requests)-1] }

// loggedIn writes a session for the fake platform into a temp config dir
// and clears the environment defaults, so a developer's exported GBF_*
// variables do not leak into the tests.
func loggedIn(t *testing.T, f *fakePlatform) {
	t.Helper()
	dir := t.TempDir()
	t.Setenv("GBF_SESSION_FILE", filepath.Join(dir, "gbf", "session.json"))
	t.Setenv("GBF_APP_URL", "")
	t.Setenv("GBF_ISSUER_URL", "")
	t.Setenv("GBF_CLIENT_ID", "")
	s := &session{AppURL: f.server.URL, IssuerURL: "https://iam.example/realms/gbf",
		Token: &oauth2.Token{AccessToken: fakeJWT(t), Expiry: time.Now().Add(time.Hour)}}
	if err := s.save(); err != nil {
		t.Fatal(err)
	}
}

func fakeJWT(t *testing.T) string {
	t.Helper()
	claims, _ := json.Marshal(map[string]any{
		"preferred_username": "researcher-alpha", "sub": "sub-1", "email": "alpha@example",
		"azp": "agent", "exp": time.Now().Add(time.Hour).Unix(),
	})
	enc := base64.RawURLEncoding.EncodeToString
	return enc([]byte(`{"alg":"none"}`)) + "." + enc(claims) + ".sig"
}

func runJSON(t *testing.T, args ...string) map[string]any {
	t.Helper()
	var out bytes.Buffer
	if err := run(append([]string{"--json"}, args...), &out); err != nil {
		t.Fatalf("%v: %v", args, err)
	}
	var v map[string]any
	if err := json.Unmarshal(out.Bytes(), &v); err != nil {
		t.Fatalf("%v: not JSON: %q", args, out.String())
	}
	return v
}

func TestSessionRoundTrip(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "gbf", "session.json")
	t.Setenv("GBF_SESSION_FILE", path)
	s := &session{AppURL: "https://app.example", IssuerURL: "https://iam.example/realms/gbf",
		Token: &oauth2.Token{AccessToken: "a", RefreshToken: "r", Expiry: time.Now().Add(time.Hour)}}
	if err := s.save(); err != nil {
		t.Fatal(err)
	}
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if info.Mode().Perm() != 0o600 {
		t.Fatalf("session file mode %o, want 0600", info.Mode().Perm())
	}
	loaded, err := loadSession()
	if err != nil {
		t.Fatal(err)
	}
	if loaded.AppURL != s.AppURL || loaded.Token.RefreshToken != "r" {
		t.Fatalf("loaded %+v", loaded)
	}
	if documentPath() != filepath.Join(dir, "gbf", "openapi.json") {
		t.Fatalf("document cache at %s", documentPath())
	}
}

func TestNotLoggedIn(t *testing.T) {
	t.Setenv("GBF_SESSION_FILE", filepath.Join(t.TempDir(), "session.json"))
	var out bytes.Buffer
	err := run([]string{"ops"}, &out)
	if err == nil || !strings.Contains(err.Error(), "gbf login") {
		t.Fatalf("expected a login hint, got %v", err)
	}
}

func TestWhoami(t *testing.T) {
	loggedIn(t, newFakePlatform(t))
	who := runJSON(t, "whoami")
	if who["username"] != "researcher-alpha" || who["clientId"] != "agent" {
		t.Fatalf("identity %v", who)
	}
	if _, leaked := who["groups"]; leaked {
		t.Fatalf("identity must not list groups: %v", who)
	}
}

func TestOpsListsAndFiltersFromTheCachedDocument(t *testing.T) {
	f := newFakePlatform(t)
	loggedIn(t, f)

	all := runJSON(t, "ops")["items"].([]any)
	if len(all) != 2 {
		t.Fatalf("ops %v", all)
	}
	first := all[0].(map[string]any)
	if first["operationId"] != "getEngagements" || first["method"] != "GET" {
		t.Fatalf("first op %v", first)
	}
	tagged := runJSON(t, "ops", "--tag", "engagements")["items"].([]any)
	if len(tagged) != 2 || tagged[1].(map[string]any)["operationId"] != "postEngagementsSubmitDraft" {
		t.Fatalf("tag filter %v", tagged)
	}
	found := runJSON(t, "ops", "--search", "draft")["items"].([]any)
	if len(found) != 1 {
		t.Fatalf("search %v", found)
	}
	if f.documents != 1 {
		t.Fatalf("document fetched %d times, want once (cached)", f.documents)
	}
	runJSON(t, "ops", "--refresh")
	if f.documents != 2 {
		t.Fatalf("--refresh did not fetch again")
	}

	var out bytes.Buffer
	if err := run([]string{"ops"}, &out); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), "OPERATION") || !strings.Contains(out.String(), "getEngagements") {
		t.Fatalf("table output %q", out.String())
	}
}

func TestDescribe(t *testing.T) {
	loggedIn(t, newFakePlatform(t))
	desc := runJSON(t, "describe", "postEngagementsSubmitDraft")
	if desc["method"] != "POST" || desc["path"] != "/engagements/{org_slug}/{engagement_slug}/submit/draft" {
		t.Fatalf("describe %v", desc)
	}
	params := desc["parameters"].([]any)
	if len(params) != 2 || params[0].(map[string]any)["required"] != true {
		t.Fatalf("parameters %v", params)
	}
	body := desc["requestBody"].(map[string]any)
	if body["mediaType"] != "application/x-www-form-urlencoded" || body["required"] != true {
		t.Fatalf("body %v", body)
	}
	fields := body["fields"].([]any)
	title := fields[1].(map[string]any)
	if title["name"] != "title" || title["required"] != true || title["maxLength"] != 200.0 || title["description"] != "Report title" {
		t.Fatalf("title field %v", title)
	}
	origin := fields[0].(map[string]any)
	if enum := origin["enum"].([]any); len(enum) != 2 || enum[0] != "human" {
		t.Fatalf("origin enum %v", origin)
	}
	responses := desc["responses"].([]any)
	if len(responses) != 2 || responses[1].(map[string]any)["schema"] != "Error" {
		t.Fatalf("responses %v", responses)
	}
	if desc["example"] != "gbf run postEngagementsSubmitDraft --param org_slug=... --param engagement_slug=... --form title=..." {
		t.Fatalf("example %q", desc["example"])
	}

	var out bytes.Buffer
	if err := run([]string{"describe", "getEngagements"}, &out); err != nil {
		t.Fatal(err)
	}
	text := out.String()
	for _, want := range []string{"getEngagements: GET /engagements", "kind (query): string, one of bug_bounty, pentest", "200: One page (EngagementListItemList)"} {
		if !strings.Contains(text, want) {
			t.Fatalf("describe text lacks %q:\n%s", want, text)
		}
	}
}

func TestRunFillsThePathAndSendsTheDeclaredBody(t *testing.T) {
	f := newFakePlatform(t)
	loggedIn(t, f)
	var out bytes.Buffer
	err := run([]string{"run", "postEngagementsSubmitDraft", "--param", "org_slug=acme", "--param", "engagement_slug=web app", "--form", "title=XSS in search"}, &out)
	if err != nil {
		t.Fatal(err)
	}
	req := f.last()
	if req.Method != http.MethodPost || req.URL.Path != "/engagements/acme/web app/submit/draft" {
		t.Fatalf("request %s %s", req.Method, req.URL.Path)
	}
	if req.Header.Get("Content-Type") != "application/x-www-form-urlencoded" || f.bodies[len(f.bodies)-1] != "title=XSS+in+search" {
		t.Fatalf("body %q", f.bodies[len(f.bodies)-1])
	}
	if req.Header.Get("Authorization") == "" || req.Header.Get("User-Agent") != "gbf-cli/dev/go" || req.Header.Get("Accept") != "application/json" {
		t.Fatalf("headers %v", req.Header)
	}

	// query parameters and a bodyless operation
	if err := run([]string{"run", "getEngagements", "--query", "kind=pentest"}, &out); err != nil {
		t.Fatal(err)
	}
	if req := f.last(); req.URL.Path != "/engagements" || req.URL.Query().Get("kind") != "pentest" {
		t.Fatalf("request %s", req.URL)
	}
}

func TestRunExplainsTheContract(t *testing.T) {
	loggedIn(t, newFakePlatform(t))
	var out bytes.Buffer
	tests := []struct {
		args []string
		want string
	}{
		{[]string{"run", "postEngagementsSubmitDraft"}, "needs --param org_slug, --param engagement_slug"},
		{[]string{"run", "postEngagementsSubmitDraft", "--param", "org_slug=a", "--param", "engagement_slug=b"}, "required: --form title"},
		{[]string{"run", "postEngagementsSubmitDraft", "--param", "org_slug=a", "--param", "engagement_slug=b", "--json-body", "{}"}, "expects form fields"},
		{[]string{"run", "getEngagements", "--form", "x=y"}, "takes no request body"},
		{[]string{"run", "nothing"}, "unknown operation"},
	}
	for _, tt := range tests {
		err := run(tt.args, &out)
		if err == nil || !strings.Contains(err.Error(), tt.want) {
			t.Fatalf("%v: got %v, want %q", tt.args, err, tt.want)
		}
	}
}

func TestRunDryRunSendsNothing(t *testing.T) {
	f := newFakePlatform(t)
	loggedIn(t, f)
	dry := runJSON(t, "run", "postEngagementsSubmitDraft", "--param", "org_slug=acme", "--param", "engagement_slug=web", "--form", "title=x", "--dry-run")
	if dry["method"] != "POST" || dry["url"] != f.server.URL+"/engagements/acme/web/submit/draft" || dry["body"] != "title=x" {
		t.Fatalf("dry run %v", dry)
	}
	for _, r := range f.requests {
		if r.URL.Path != "/openapi.json" {
			t.Fatalf("dry run sent %s", r.URL.Path)
		}
	}
}

func TestSchemaAndOpenapi(t *testing.T) {
	loggedIn(t, newFakePlatform(t))
	if s := runJSON(t, "schema", "Error"); s["type"] != "object" {
		t.Fatalf("schema %v", s)
	}
	var out bytes.Buffer
	if err := run([]string{"schema", "Nope"}, &out); err == nil || !strings.Contains(err.Error(), "available: EngagementListItemList, Error") {
		t.Fatalf("unknown schema: %v", err)
	}
	if doc := runJSON(t, "openapi"); doc["openapi"] != "3.1.0" {
		t.Fatalf("openapi %v", doc)
	}
}

func TestCallSendsBearerAndPrintsJSON(t *testing.T) {
	f := newFakePlatform(t)
	loggedIn(t, f)
	var out bytes.Buffer
	if err := run([]string{"call", "post", "reports/draft", "--form", "title=Hi", "--form", "severity=high"}, &out); err != nil {
		t.Fatal(err)
	}
	req := f.last()
	if req.Method != http.MethodPost || req.URL.Path != "/reports/draft" {
		t.Fatalf("request %s %s", req.Method, req.URL.Path)
	}
	if !strings.HasPrefix(req.Header.Get("Authorization"), "Bearer ") || req.Header.Get("Accept") != "application/json" {
		t.Fatalf("headers %v", req.Header)
	}
	if req.Header.Get("User-Agent") != "gbf-cli/dev/go" {
		t.Fatalf("user agent %q", req.Header.Get("User-Agent"))
	}
	if f.bodies[len(f.bodies)-1] != "severity=high&title=Hi" {
		t.Fatalf("form body %q", f.bodies[len(f.bodies)-1])
	}
	if !strings.Contains(out.String(), `"path": "/reports/draft"`) {
		t.Fatalf("output %q", out.String())
	}
}

func TestCallReportsErrors(t *testing.T) {
	loggedIn(t, newFakePlatform(t))
	var out bytes.Buffer
	err := run([]string{"call", "GET", "/nope"}, &out)
	if err == nil || !strings.Contains(err.Error(), "404") {
		t.Fatalf("expected a 404 error, got %v", err)
	}
	if !strings.Contains(out.String(), `"status": 404`) {
		t.Fatalf("body not printed: %q", out.String())
	}
}

// The skill ships inside the binary: printing and installing it needs no
// session and no network.
func TestSkillIsEmbeddedAndInstallable(t *testing.T) {
	t.Setenv("GBF_SESSION_FILE", filepath.Join(t.TempDir(), "none.json"))
	var out bytes.Buffer
	if err := run([]string{"skill"}, &out); err != nil || !strings.HasPrefix(out.String(), "---\nname: gbf\n") {
		t.Fatalf("skill: %v %q", err, out.String())
	}
	dir := t.TempDir()
	out.Reset()
	if err := run([]string{"skill", "--install", "--dir", dir}, &out); err != nil {
		t.Fatalf("install: %v", err)
	}
	installed, err := os.ReadFile(filepath.Join(dir, "gbf", "SKILL.md"))
	if err != nil || !strings.HasPrefix(string(installed), "---\nname: gbf\n") {
		t.Fatalf("installed skill: %v %q", err, installed)
	}
	if !strings.Contains(out.String(), "version dev+") {
		t.Fatalf("install output %q", out.String())
	}
	if s := runJSON(t, "skill"); s["name"] != "gbf" || !strings.HasPrefix(s["content"].(string), "---") {
		t.Fatalf("skill json %v", s)
	}
}

// --agent picks the user-level skills directory of the named agent under
// the home directory; --dir wins over it, and an unknown agent is refused.
func TestSkillInstallsForTheNamedAgent(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	for agent, dir := range map[string]string{"": ".claude", "claude": ".claude", "cursor": ".cursor", "codex": ".agents", "lmstudio": ".lmstudio"} {
		args := []string{"skill", "--install"}
		if agent != "" {
			args = append(args, "--agent", agent)
		}
		var out bytes.Buffer
		if err := run(args, &out); err != nil {
			t.Fatalf("%v: %v", args, err)
		}
		target := filepath.Join(home, dir, "skills", "gbf", "SKILL.md")
		if _, err := os.Stat(target); err != nil {
			t.Fatalf("%v: %v", args, err)
		}
		if !strings.Contains(out.String(), target) {
			t.Fatalf("%v: output %q lacks %s", args, out.String(), target)
		}
	}
	explicit := t.TempDir()
	var out bytes.Buffer
	if err := run([]string{"skill", "--install", "--agent", "cursor", "--dir", explicit}, &out); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(explicit, "gbf", "SKILL.md")); err != nil {
		t.Fatalf("--dir ignored: %v", err)
	}
	err := run([]string{"skill", "--install", "--agent", "copilot"}, &out)
	if err == nil || !strings.Contains(err.Error(), "unknown agent") {
		t.Fatalf("unknown agent: %v", err)
	}
}

func TestCloudflareAccessHeadersOnEveryRequest(t *testing.T) {
	f := newFakePlatform(t)
	loggedIn(t, f)
	t.Setenv("CF_ACCESS_CLIENT_ID", "cf-id")
	t.Setenv("CF_ACCESS_CLIENT_SECRET", "cf-secret")

	var out bytes.Buffer
	if err := run([]string{"call", "GET", "/engagements"}, &out); err != nil {
		t.Fatal(err)
	}
	req := f.last()
	if req.Header.Get("CF-Access-Client-Id") != "cf-id" || req.Header.Get("CF-Access-Client-Secret") != "cf-secret" {
		t.Fatalf("access headers missing on an API call: %v", req.Header)
	}
	if !strings.HasPrefix(req.Header.Get("Authorization"), "Bearer ") {
		t.Fatalf("bearer missing next to the access headers: %v", req.Header)
	}

	// a half-configured token is rejected before any request
	t.Setenv("CF_ACCESS_CLIENT_SECRET", "")
	err := run([]string{"call", "GET", "/engagements"}, &out)
	if err == nil || !strings.Contains(err.Error(), "must be set together") {
		t.Fatalf("expected the pairing error, got %v", err)
	}
}

func readAll(r *http.Request) (string, error) {
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r.Body)
	return buf.String(), err
}

// __complete answers the shells from the cached document: commands, flags,
// operation ids, parameter and field names, enum values, paths, schemas.
func TestCompleteFromTheCachedDocument(t *testing.T) {
	f := newFakePlatform(t)
	loggedIn(t, f)
	runJSON(t, "ops") // caches the document
	fetched := f.documents

	complete := func(words ...string) string {
		var out bytes.Buffer
		if err := run(append([]string{"__complete", "--"}, words...), &out); err != nil {
			t.Fatalf("%v: %v", words, err)
		}
		return out.String()
	}
	values := func(words ...string) []string {
		var vals []string
		for _, line := range strings.Split(strings.TrimSpace(complete(words...)), "\n") {
			if !strings.HasPrefix(line, ":") {
				vals = append(vals, strings.SplitN(line, "\t", 2)[0])
			}
		}
		return vals
	}
	tests := []struct {
		words []string
		want  []string
	}{
		{[]string{"r"}, []string{"run"}},
		{[]string{"-"}, []string{"--json"}},
		{[]string{"--json", "ru"}, []string{"run"}},
		{[]string{"run", ""}, []string{"getEngagements", "postEngagementsSubmitDraft"}},
		{[]string{"describe", "post"}, []string{"postEngagementsSubmitDraft"}},
		{[]string{"run", "postEngagementsSubmitDraft", "--"}, []string{"--dry-run", "--param", "--form"}},
		{[]string{"run", "getEngagements", "--"}, []string{"--dry-run", "--query"}},
		{[]string{"run", "postEngagementsSubmitDraft", "--param", ""}, []string{"org_slug=", "engagement_slug="}},
		{[]string{"run", "postEngagementsSubmitDraft", "--param", "org_slug=acme", "--param", ""}, []string{"engagement_slug="}},
		{[]string{"run", "postEngagementsSubmitDraft", "--form", ""}, []string{"origin=", "title="}},
		{[]string{"run", "postEngagementsSubmitDraft", "--form", "origin="}, []string{"origin=human", "origin=ai_agent"}},
		{[]string{"run", "postEngagementsSubmitDraft", "--form=ori"}, []string{"--form=origin="}},
		{[]string{"run", "getEngagements", "--query", "kind="}, []string{"kind=bug_bounty", "kind=pentest"}},
		{[]string{"ops", "--tag", ""}, []string{"engagements"}},
		{[]string{"call", "p"}, nil},
		{[]string{"call", "P"}, []string{"POST", "PUT", "PATCH"}},
		{[]string{"call", "get", "/"}, []string{"/engagements"}},
		{[]string{"schema", "E"}, []string{"EngagementListItemList", "Error"}},
		{[]string{"skill", "--agent", ""}, []string{"claude", "cursor", "codex", "lmstudio"}},
		{[]string{"completion", "z"}, []string{"zsh"}},
	}
	for _, tt := range tests {
		if got := values(tt.words...); strings.Join(got, " ") != strings.Join(tt.want, " ") {
			t.Errorf("%v: got %v, want %v", tt.words, got, tt.want)
		}
	}
	if out := complete("skill", "--dir", ""); out != ":dirs\n" {
		t.Errorf("--dir: %q", out)
	}
	if out := complete("run", "getEngagements", "--"); !strings.HasSuffix(out, ":nofile\n") || !strings.Contains(out, "--dry-run\tprint the request") {
		t.Errorf("descriptions and directive: %q", out)
	}
	if f.documents != fetched {
		t.Fatalf("completion fetched the document")
	}

	// without a cache, static words still complete and nothing fails
	t.Setenv("GBF_SESSION_FILE", filepath.Join(t.TempDir(), "session.json"))
	if got := values("run", ""); got != nil {
		t.Errorf("no cache: %v", got)
	}
	if got := values("sk"); strings.Join(got, " ") != "skill" {
		t.Errorf("no cache commands: %v", got)
	}
}

func TestCompletionScripts(t *testing.T) {
	for _, shell := range []string{"bash", "zsh", "fish", "powershell"} {
		var out bytes.Buffer
		if err := run([]string{"completion", shell}, &out); err != nil || !strings.Contains(out.String(), "__complete") {
			t.Fatalf("%s: %v %q", shell, err, out.String())
		}
	}
	var out bytes.Buffer
	if err := run([]string{"completion", "ksh"}, &out); err == nil || !strings.Contains(err.Error(), "unknown shell") {
		t.Fatalf("ksh: %v", err)
	}
}

// --install writes the script where the shell loads it by itself.
func TestCompletionInstall(t *testing.T) {
	data, config := t.TempDir(), t.TempDir()
	t.Setenv("XDG_DATA_HOME", data)
	t.Setenv("XDG_CONFIG_HOME", config)
	for shell, target := range map[string]string{
		"bash": filepath.Join(data, "bash-completion", "completions", "gbf"),
		"zsh":  filepath.Join(data, "zsh", "site-functions", "_gbf"),
		"fish": filepath.Join(config, "fish", "completions", "gbf.fish"),
	} {
		var out bytes.Buffer
		if err := run([]string{"completion", shell, "--install"}, &out); err != nil {
			t.Fatalf("%s: %v", shell, err)
		}
		installed, err := os.ReadFile(target)
		if err != nil || !strings.Contains(string(installed), "__complete") {
			t.Fatalf("%s: %v %q", shell, err, installed)
		}
		if !strings.Contains(out.String(), target) {
			t.Fatalf("%s: output %q", shell, out.String())
		}
	}
	var out bytes.Buffer
	if err := run([]string{"completion", "powershell", "--install"}, &out); err == nil || !strings.Contains(err.Error(), "$PROFILE") {
		t.Fatalf("powershell: %v", err)
	}
}
