package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Shell completion. gbf completion SHELL prints a script that asks
// gbf __complete for candidates on every tab. The hidden command receives
// the words of the command line after gbf, the last one being the word
// under the cursor (empty when the cursor follows a space), and prints one
// candidate per line as VALUE<tab>DESCRIPTION, then a directive line:
// :nofile (the default), :dirs (the shell should offer directories), or
// :files. Candidates come from the cached document only; without a login
// or a cache, static words are still completed. It never fails and never
// touches the network, so a tab is always quiet and fast.

type candidate struct {
	value, description string
}

type completion struct {
	candidates []candidate
	directive  string
}

func (c *cli) complete(args []string) error {
	words := args
	if len(words) == 0 {
		words = []string{""}
	}
	cur := words[len(words)-1]
	result := c.completeWords(words[:len(words)-1], cur)
	if result.directive == "" {
		result.directive = ":nofile"
	}
	for _, cand := range result.candidates {
		if !strings.HasPrefix(cand.value, cur) {
			continue
		}
		fmt.Fprintf(c.out, "%s\t%s\n", cand.value, cand.description)
	}
	fmt.Fprintln(c.out, result.directive)
	return nil
}

var commands = []candidate{
	{"login", "sign in with the device flow"},
	{"logout", "forget the session on this machine"},
	{"status", "session, platform, and token state"},
	{"whoami", "who the token says you are"},
	{"ops", "the operations you may call"},
	{"describe", "one operation's parameters, body, and responses"},
	{"run", "invoke an operation by id"},
	{"call", "send any request"},
	{"schema", "one component schema"},
	{"openapi", "the whole document"},
	{"skill", "the agent skill"},
	{"completion", "shell completion script"},
	{"version", "print the version"},
	{"help", "usage"},
}

// valueFlags lists, per command, the flags that take a value; every other
// flag is a boolean switch.
var valueFlags = map[string][]string{
	"login": {"app-url", "issuer-url", "client-id"},
	"ops":   {"tag", "search"},
	"run":   {"param", "query", "form", "json-body"},
	"call":  {"json-body", "form"},
	"skill": {"agent", "dir"},
}

func (c *cli) completeWords(prev []string, cur string) completion {
	if len(prev) > 0 && prev[0] == "--json" {
		prev = prev[1:]
	} else if len(prev) == 0 && strings.HasPrefix(cur, "-") {
		return completion{candidates: []candidate{{"--json", "JSON output"}}}
	}
	if len(prev) == 0 {
		return completion{candidates: commands}
	}
	command, rest := prev[0], prev[1:]

	// the value of a flag: --flag VALUE, or --flag=VALUE in one word
	flag, prefix, value, isValue := flagValue(command, rest, cur)
	if isValue {
		result := c.completeFlagValue(command, flag, rest, value)
		for i := range result.candidates {
			result.candidates[i].value = prefix + result.candidates[i].value
		}
		return result
	}
	if strings.HasPrefix(cur, "-") {
		return completion{candidates: c.flagsFor(command, rest)}
	}
	return c.completePositional(command, positionals(command, rest), cur)
}

// flagValue reports whether cur is the value of a flag: either the previous
// word is a value-taking flag, or cur itself is --flag=VALUE. It returns
// the flag name, the prefix to keep in front of every candidate, and the
// value typed so far.
func flagValue(command string, prev []string, cur string) (flag, prefix, value string, ok bool) {
	if strings.HasPrefix(cur, "-") {
		name, val, found := strings.Cut(strings.TrimLeft(cur, "-"), "=")
		if found && contains(valueFlags[command], name) {
			return name, cur[:len(cur)-len(val)], val, true
		}
		return "", "", "", false
	}
	if len(prev) == 0 {
		return "", "", "", false
	}
	last := prev[len(prev)-1]
	if strings.HasPrefix(last, "-") && !strings.Contains(last, "=") {
		if name := strings.TrimLeft(last, "-"); contains(valueFlags[command], name) {
			return name, "", cur, true
		}
	}
	return "", "", "", false
}

// positionals returns the positional words among the arguments so far.
func positionals(command string, words []string) []string {
	var out []string
	for i := 0; i < len(words); i++ {
		w := words[i]
		if !strings.HasPrefix(w, "-") {
			out = append(out, w)
			continue
		}
		if name := strings.TrimLeft(w, "-"); !strings.Contains(w, "=") && contains(valueFlags[command], name) {
			i++
		}
	}
	return out
}

func (c *cli) flagsFor(command string, rest []string) []candidate {
	switch command {
	case "login":
		return []candidate{
			{"--app-url", "platform base URL"},
			{"--issuer-url", "identity provider realm URL"},
			{"--client-id", "device-flow client id"},
		}
	case "ops":
		return []candidate{
			{"--tag", "only one area"},
			{"--search", "only ids, paths, or summaries containing a text"},
			{"--refresh", "fetch the document again"},
		}
	case "run":
		flags := []candidate{{"--dry-run", "print the request instead of sending it"}}
		op := c.operationFor(positionals(command, rest))
		if op == nil {
			return append(flags,
				candidate{"--param", "path parameter KEY=VALUE"},
				candidate{"--query", "query parameter KEY=VALUE"},
				candidate{"--form", "form field KEY=VALUE"},
				candidate{"--json-body", "JSON request body"},
			)
		}
		if len(c.params(op, "path")) > 0 {
			flags = append(flags, candidate{"--param", "path parameter KEY=VALUE"})
		}
		if len(c.params(op, "query")) > 0 {
			flags = append(flags, candidate{"--query", "query parameter KEY=VALUE"})
		}
		if op.RequestBody != nil {
			if _, json := op.RequestBody.Content["application/json"]; json {
				flags = append(flags, candidate{"--json-body", "JSON request body"})
			} else {
				flags = append(flags, candidate{"--form", "form field KEY=VALUE"})
			}
		}
		return flags
	case "call":
		return []candidate{{"--json-body", "JSON request body"}, {"--form", "form field KEY=VALUE"}}
	case "openapi":
		return []candidate{{"--refresh", "fetch the document again"}}
	case "completion":
		return []candidate{{"--install", "write the script where the shell loads it"}}
	case "skill":
		return []candidate{
			{"--install", "write the skill where the agent loads it"},
			{"--agent", "claude, cursor, codex, or lmstudio"},
			{"--dir", "skills directory instead of the agent's default"},
		}
	}
	return nil
}

func (c *cli) completeFlagValue(command, flag string, rest []string, value string) completion {
	switch command + " --" + flag {
	case "ops --tag":
		return completion{candidates: c.tags()}
	case "skill --agent":
		return completion{candidates: []candidate{
			{"claude", "~/.claude/skills"}, {"cursor", "~/.cursor/skills"},
			{"codex", "~/.agents/skills"}, {"lmstudio", "~/.lmstudio/skills"},
		}}
	case "skill --dir":
		return completion{directive: ":dirs"}
	case "run --param":
		return c.keyValues(rest, value, func(op *operation) []field { return c.params(op, "path") })
	case "run --query":
		return c.keyValues(rest, value, func(op *operation) []field { return c.params(op, "query") })
	case "run --form":
		return c.keyValues(rest, value, c.formFields)
	}
	return completion{}
}

// keyValues completes KEY= for the operation's fields not given yet, and
// KEY=VALUE once the key is typed and the field is an enumeration.
func (c *cli) keyValues(rest []string, value string, fields func(*operation) []field) completion {
	op := c.operationFor(positionals("run", rest))
	if op == nil {
		return completion{}
	}
	key, _, hasKey := strings.Cut(value, "=")
	var cands []candidate
	for _, f := range fields(op) {
		switch {
		case hasKey && f.Name == key:
			for _, v := range f.Enum {
				cands = append(cands, candidate{key + "=" + v, f.Description})
			}
			return completion{candidates: cands}
		case !hasKey && !given(rest, f.Name):
			desc := f.Type
			if f.Required {
				desc += ", required"
			}
			if f.Description != "" {
				desc += ": " + f.Description
			}
			cands = append(cands, candidate{f.Name + "=", desc})
		}
	}
	return completion{candidates: cands}
}

// given reports whether a KEY=VALUE for key already appears in the words.
func given(words []string, key string) bool {
	for _, w := range words {
		w = strings.TrimPrefix(strings.TrimPrefix(w, "--param="), "--query=")
		w = strings.TrimPrefix(w, "--form=")
		if strings.HasPrefix(w, key+"=") {
			return true
		}
	}
	return false
}

func (c *cli) completePositional(command string, before []string, cur string) completion {
	switch command {
	case "describe":
		if len(before) == 0 {
			return completion{candidates: c.operations()}
		}
	case "run":
		if len(before) == 0 {
			return completion{candidates: c.operations()}
		}
	case "call":
		switch len(before) {
		case 0:
			return completion{candidates: []candidate{{"GET", ""}, {"POST", ""}, {"PUT", ""}, {"PATCH", ""}, {"DELETE", ""}}}
		case 1:
			return completion{candidates: c.paths(strings.ToUpper(before[0]))}
		}
	case "schema":
		if len(before) == 0 {
			return completion{candidates: c.schemas()}
		}
	case "completion":
		if len(before) == 0 {
			return completion{candidates: []candidate{{"bash", ""}, {"zsh", ""}, {"fish", ""}, {"powershell", ""}}}
		}
	}
	return completion{}
}

// cachedDocument is the session's cached document, or nil: completion
// never fetches.
func (c *cli) cachedDocument() *document {
	data, err := os.ReadFile(documentPath())
	if err != nil {
		return nil
	}
	doc, err := parseDocument(data)
	if err != nil {
		return nil
	}
	return doc
}

func (c *cli) operationFor(positional []string) *operation {
	if len(positional) == 0 {
		return nil
	}
	doc := c.cachedDocument()
	if doc == nil {
		return nil
	}
	return doc.byId(positional[0])
}

func (c *cli) operations() []candidate {
	doc := c.cachedDocument()
	if doc == nil {
		return nil
	}
	var cands []candidate
	for _, op := range doc.operations() {
		cands = append(cands, candidate{op.OperationId, op.Method + " " + op.Path + "  " + truncate(op.Summary, 50)})
	}
	return cands
}

func (c *cli) tags() []candidate {
	doc := c.cachedDocument()
	if doc == nil {
		return nil
	}
	count := map[string]int{}
	for _, op := range doc.operations() {
		for _, tag := range op.Tags {
			count[tag]++
		}
	}
	names := make([]string, 0, len(count))
	for name := range count {
		names = append(names, name)
	}
	sort.Strings(names)
	cands := make([]candidate, 0, len(names))
	for _, name := range names {
		cands = append(cands, candidate{name, fmt.Sprintf("%d operations", count[name])})
	}
	return cands
}

func (c *cli) paths(method string) []candidate {
	doc := c.cachedDocument()
	if doc == nil {
		return nil
	}
	var cands []candidate
	for _, op := range doc.operations() {
		if method == "" || op.Method == method {
			cands = append(cands, candidate{op.Path, op.OperationId})
		}
	}
	return cands
}

func (c *cli) schemas() []candidate {
	doc := c.cachedDocument()
	if doc == nil {
		return nil
	}
	names := make([]string, 0, len(doc.Components.Schemas))
	for name := range doc.Components.Schemas {
		names = append(names, name)
	}
	sort.Strings(names)
	cands := make([]candidate, 0, len(names))
	for _, name := range names {
		cands = append(cands, candidate{name, ""})
	}
	return cands
}

// params returns the operation's parameters of one location as fields.
func (c *cli) params(op *operation, in string) []field {
	doc := c.cachedDocument()
	if doc == nil {
		return nil
	}
	var fields []field
	for _, p := range op.Parameters {
		if p.In != in {
			continue
		}
		f := doc.field(p.Name, p.Schema, p.Required)
		f.Description = p.Description
		fields = append(fields, f)
	}
	return fields
}

func (c *cli) formFields(op *operation) []field {
	if op.RequestBody == nil {
		return nil
	}
	doc := c.cachedDocument()
	if doc == nil {
		return nil
	}
	for mediaType, content := range op.RequestBody.Content {
		if mediaType == "application/json" {
			return nil
		}
		return doc.fields(content.Schema)
	}
	return nil
}

// completionScript prints the completion script of one shell, or with
// --install writes it where that shell loads it by itself.
func (c *cli) completionScript(args []string) error {
	flags := newFlags("completion", c.out)
	install := flags.Bool("install", false, "write the script where the shell loads it, instead of printing it")
	positional, err := parseInterleaved(flags, args)
	if err != nil {
		return err
	}
	if len(positional) != 1 {
		return errors.New("usage: gbf completion bash|zsh|fish|powershell [--install]")
	}
	shell := positional[0]
	script, ok := completionScripts[shell]
	if !ok {
		return fmt.Errorf("unknown shell %q: use bash, zsh, fish, or powershell", shell)
	}
	if !*install {
		_, err := fmt.Fprint(c.out, script)
		return err
	}
	target, note, err := completionTarget(shell)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(target, []byte(script), 0o644); err != nil {
		return err
	}
	fmt.Fprintf(c.out, "Installed the %s completion at %s\n%s\n", shell, target, note)
	return nil
}

// completionTarget is the file a shell loads completions from by itself,
// with what the user still has to do once, if anything.
func completionTarget(shell string) (target, note string, err error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", "", err
	}
	dataHome := os.Getenv("XDG_DATA_HOME")
	if dataHome == "" {
		dataHome = filepath.Join(home, ".local", "share")
	}
	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome == "" {
		configHome = filepath.Join(home, ".config")
	}
	switch shell {
	case "bash":
		return filepath.Join(dataHome, "bash-completion", "completions", "gbf"),
			"bash-completion loads it on the first tab in every new shell. Without the bash-completion package, add to ~/.bashrc instead:\n  eval \"$(gbf completion bash)\"", nil
	case "zsh":
		dir := filepath.Join(dataHome, "zsh", "site-functions")
		return filepath.Join(dir, "_gbf"),
			"Make sure ~/.zshrc lists the directory before compinit, then open a new shell:\n  fpath+=(" + dir + ")", nil
	case "fish":
		return filepath.Join(configHome, "fish", "completions", "gbf.fish"),
			"fish loads it in every new shell.", nil
	case "powershell":
		return "", "", errors.New("PowerShell has no completion directory; add to $PROFILE instead:\n  gbf completion powershell | Out-String | Invoke-Expression")
	}
	return "", "", fmt.Errorf("unknown shell %q", shell)
}

var completionScripts = map[string]string{
	"bash": `# bash completion for gbf: eval "$(gbf completion bash)"
_gbf() {
    local cur words cword
    if declare -F _init_completion >/dev/null 2>&1; then
        _init_completion -n "=:" || return
    else
        words=("${COMP_WORDS[@]}"); cword=$COMP_CWORD; cur=${COMP_WORDS[COMP_CWORD]}
    fi
    local IFS=$'\n' line directive="" value
    local -a values=()
    while read -r line; do
        case "$line" in
            :*) directive=$line ;;
            *) value=${line%%$'\t'*}; values+=("$value") ;;
        esac
    done < <("${words[0]}" __complete -- "${words[@]:1:cword}" 2>/dev/null)
    case "$directive" in
        :dirs) compopt -o dirnames 2>/dev/null; COMPREPLY=(); return ;;
        :files) compopt -o default 2>/dev/null; COMPREPLY=(); return ;;
    esac
    # bash splits words at = and : unless bash-completion joined them; keep
    # only what follows the last such break in the current word
    local cur_bash=${COMP_WORDS[COMP_CWORD]} prefix=""
    if [[ "$cur_bash" == "=" || "$cur_bash" == ":" ]]; then
        prefix=$cur
    elif [[ "$cur" != "$cur_bash" && "$cur" == *"$cur_bash" ]]; then
        prefix=${cur%"$cur_bash"}
    fi
    COMPREPLY=()
    for value in "${values[@]}"; do
        [[ "$value" == *= ]] && compopt -o nospace 2>/dev/null
        COMPREPLY+=("${value#"$prefix"}")
    done
}
complete -o nosort -F _gbf gbf 2>/dev/null || complete -F _gbf gbf
`,
	"zsh": `#compdef gbf -P */gbf
# zsh completion for gbf: eval "$(gbf completion zsh)"
# Registered for gbf and for any path ending in /gbf, such as ./gbf.
_gbf() {
    local -a lines plain nospace
    local line value description directive=""
    lines=("${(@f)$(${words[1]} __complete -- "${(@)words[2,CURRENT]}" 2>/dev/null)}")
    for line in "${lines[@]}"; do
        case "$line" in
            :*) directive=$line ;;
            "") ;;
            *)
                value=${line%%$'\t'*}
                description=${line#*$'\t'}
                value=${value//:/\\:}
                if [[ "$value" == *= ]]; then
                    nospace+=("${value}:${description}")
                else
                    plain+=("${value}:${description}")
                fi
                ;;
        esac
    done
    case "$directive" in
        :dirs) _files -/; return ;;
        :files) _files; return ;;
    esac
    (( ${#nospace} )) && _describe -t gbf-values 'gbf' nospace -S ''
    (( ${#plain} )) && _describe -t gbf-values 'gbf' plain
    return 0
}
if [[ "${zsh_eval_context[-1]}" == loadautofunc ]]; then
    _gbf "$@"
else
    compdef _gbf gbf -P '*/gbf'
fi
`,
	"fish": `# fish completion for gbf: gbf completion fish | source
function __gbf_complete
    set -l tokens (commandline -cpo)
    set -l cur (commandline -ct)
    set -l lines ($tokens[1] __complete -- $tokens[2..-1] $cur 2>/dev/null)
    for line in $lines
        switch $line
            case ':dirs'
                __fish_complete_directories $cur
            case ':files'
                __fish_complete_path $cur
            case ':*'
            case '*'
                echo $line
        end
    end
end
complete -c gbf -f -a '(__gbf_complete)'
`,
	"powershell": `# PowerShell completion for gbf: gbf completion powershell | Out-String | Invoke-Expression
Register-ArgumentCompleter -Native -CommandName gbf -ScriptBlock {
    param($wordToComplete, $commandAst, $cursorPosition)
    $elements = $commandAst.CommandElements
    $words = @()
    for ($i = 1; $i -lt $elements.Count; $i++) {
        if ($elements[$i].Extent.StartOffset -ge $cursorPosition) { break }
        $words += $elements[$i].Extent.Text
    }
    if ($wordToComplete -eq '' -or $words.Count -eq 0 -or $words[-1] -ne $wordToComplete) { $words += $wordToComplete }
    $lines = & $elements[0].Extent.Text __complete -- @words 2>$null
    foreach ($line in $lines) {
        if ($line.StartsWith(':')) { continue }
        $parts = $line -split "` + "`t" + `", 2
        $description = if ($parts.Count -gt 1 -and $parts[1]) { $parts[1] } else { $parts[0] }
        [System.Management.Automation.CompletionResult]::new($parts[0], $parts[0], 'ParameterValue', $description)
    }
}
`,
}
