package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/tabwriter"
)

// document is the slice of the platform's OpenAPI document the command
// works with. It is scoped to the logged-in user by the server, so the
// operations it lists are exactly the ones this user may call.
type document struct {
	raw        []byte
	Info       struct{ Version string }            `json:"info"`
	Paths      map[string]map[string]operation     `json:"paths"`
	Components struct{ Schemas map[string]schema } `json:"components"`
	Servers    []struct {
		URL string `json:"url"`
	} `json:"servers"`
}

type operation struct {
	// Method and Path are filled from the document's structure.
	Method string `json:"-"`
	Path   string `json:"-"`

	OperationId string           `json:"operationId"`
	Summary     string           `json:"summary,omitempty"`
	Description string           `json:"description,omitempty"`
	Tags        []string         `json:"tags,omitempty"`
	Parameters  []parameter      `json:"parameters,omitempty"`
	RequestBody *requestBodySpec `json:"requestBody,omitempty"`
	Responses   map[string]struct {
		Description string             `json:"description"`
		Content     map[string]content `json:"content,omitempty"`
	} `json:"responses"`
}

type parameter struct {
	Name        string `json:"name"`
	In          string `json:"in"`
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Schema      schema `json:"schema"`
}

type requestBodySpec struct {
	Description string             `json:"description,omitempty"`
	Required    bool               `json:"required,omitempty"`
	Content     map[string]content `json:"content"`
}

type content struct {
	Schema schema `json:"schema"`
}

// schema keeps the JSON schema as loosely typed data; only the parts the
// command renders are read from it.
type schema map[string]any

func documentPath() string {
	return filepath.Join(filepath.Dir(sessionPath()), "openapi.json")
}

func parseDocument(data []byte) (*document, error) {
	var doc document
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, fmt.Errorf("parse the OpenAPI document: %w", err)
	}
	doc.raw = data
	for path, item := range doc.Paths {
		for method, op := range item {
			op.Method, op.Path = strings.ToUpper(method), path
			item[method] = op
		}
	}
	return &doc, nil
}

// loadDocument returns the user's document, from the session cache unless
// refresh is set or nothing is cached yet.
func (c *cli) loadDocument(ctx context.Context, s *session, refresh bool) (*document, error) {
	if !refresh {
		if data, err := os.ReadFile(documentPath()); err == nil {
			if doc, err := parseDocument(data); err == nil {
				return doc, nil
			}
		}
	}
	res, err := s.request(ctx, http.MethodGet, s.AppURL+"/openapi.json", "application/json", "", nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET /openapi.json: %s", res.Status)
	}
	doc, err := parseDocument(data)
	if err != nil {
		return nil, err
	}
	if err := os.WriteFile(documentPath(), data, 0o600); err != nil {
		fmt.Fprintln(os.Stderr, "gbf: could not cache the OpenAPI document:", err)
	}
	return doc, nil
}

// operations returns every operation sorted by tag, path, and method.
func (d *document) operations() []operation {
	var ops []operation
	for _, item := range d.Paths {
		for _, op := range item {
			ops = append(ops, op)
		}
	}
	sort.Slice(ops, func(i, j int) bool {
		a, b := ops[i], ops[j]
		if ta, tb := firstTag(a), firstTag(b); ta != tb {
			return ta < tb
		}
		if a.Path != b.Path {
			return a.Path < b.Path
		}
		return a.Method < b.Method
	})
	return ops
}

func firstTag(op operation) string {
	if len(op.Tags) > 0 {
		return op.Tags[0]
	}
	return ""
}

// find looks an operation up by id, refreshing the cached document once
// when the id is unknown: the routes may have changed since it was cached.
func (c *cli) find(ctx context.Context, s *session, id string) (*document, *operation, error) {
	doc, err := c.loadDocument(ctx, s, false)
	if err != nil {
		return nil, nil, err
	}
	if op := doc.byId(id); op != nil {
		return doc, op, nil
	}
	if doc, err = c.loadDocument(ctx, s, true); err != nil {
		return nil, nil, err
	}
	if op := doc.byId(id); op != nil {
		return doc, op, nil
	}
	return nil, nil, fmt.Errorf("unknown operation %q for this user; gbf ops lists the available ones", id)
}

func (d *document) byId(id string) *operation {
	for _, item := range d.Paths {
		for _, op := range item {
			if op.OperationId == id {
				return &op
			}
		}
	}
	return nil
}

func (c *cli) ops(ctx context.Context, args []string) error {
	flags := newFlags("ops", c.out)
	tag := flags.String("tag", "", "only operations with this tag")
	search := flags.String("search", "", "only operations whose id, path, or summary contains this text")
	refresh := flags.Bool("refresh", false, "fetch the document again instead of using the cached one")
	if err := flags.Parse(args); err != nil {
		return err
	}
	s, err := loadSession()
	if err != nil {
		return err
	}
	doc, err := c.loadDocument(ctx, s, *refresh)
	if err != nil {
		return err
	}
	needle := strings.ToLower(*search)
	var ops []operation
	for _, op := range doc.operations() {
		if *tag != "" && !contains(op.Tags, *tag) {
			continue
		}
		if needle != "" && !strings.Contains(strings.ToLower(op.OperationId+" "+op.Path+" "+op.Summary), needle) {
			continue
		}
		ops = append(ops, op)
	}
	if c.json {
		type item struct {
			OperationId string   `json:"operationId"`
			Method      string   `json:"method"`
			Path        string   `json:"path"`
			Tags        []string `json:"tags"`
			Summary     string   `json:"summary"`
		}
		items := make([]item, 0, len(ops))
		for _, op := range ops {
			items = append(items, item{op.OperationId, op.Method, op.Path, op.Tags, op.Summary})
		}
		return c.printJSON(map[string]any{"items": items})
	}
	w := tabwriter.NewWriter(c.out, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "OPERATION\tMETHOD\tPATH\tSUMMARY")
	for _, op := range ops {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", op.OperationId, op.Method, op.Path, truncate(op.Summary, 72))
	}
	return w.Flush()
}

func (c *cli) describe(ctx context.Context, args []string) error {
	if len(args) != 1 {
		return errors.New("usage: gbf describe OPERATION_ID")
	}
	s, err := loadSession()
	if err != nil {
		return err
	}
	doc, op, err := c.find(ctx, s, args[0])
	if err != nil {
		return err
	}
	desc := doc.describe(op)
	if c.json {
		return c.printJSON(desc)
	}
	fmt.Fprintf(c.out, "%s: %s %s\n", op.OperationId, op.Method, op.Path)
	if op.Summary != "" {
		fmt.Fprintf(c.out, "  %s\n", op.Summary)
	}
	if op.Description != "" {
		fmt.Fprintf(c.out, "  %s\n", op.Description)
	}
	if len(desc.Parameters) > 0 {
		fmt.Fprintln(c.out, "\nParameters:")
		for _, p := range desc.Parameters {
			fmt.Fprintf(c.out, "  %s (%s%s)%s\n", p.Name, p.In, requiredMark(p.Required), fieldDetails(p.field))
		}
	}
	if desc.RequestBody != nil {
		fmt.Fprintf(c.out, "\nRequest body (%s%s):\n", desc.RequestBody.MediaType, requiredMark(desc.RequestBody.Required))
		for _, f := range desc.RequestBody.Fields {
			fmt.Fprintf(c.out, "  %s%s%s\n", f.Name, requiredMark(f.Required), fieldDetails(f))
		}
	}
	fmt.Fprintln(c.out, "\nResponses:")
	for _, r := range desc.Responses {
		line := "  " + r.Status + ": " + r.Description
		if r.Schema != "" {
			line += " (" + r.Schema + ")"
		}
		fmt.Fprintln(c.out, line)
	}
	fmt.Fprintf(c.out, "\nExample:\n  %s\n", desc.Example)
	return nil
}

// description is the agent-facing contract of one operation.
type description struct {
	OperationId string         `json:"operationId"`
	Method      string         `json:"method"`
	Path        string         `json:"path"`
	Summary     string         `json:"summary,omitempty"`
	Description string         `json:"description,omitempty"`
	Tags        []string       `json:"tags,omitempty"`
	Parameters  []paramDesc    `json:"parameters"`
	RequestBody *bodyDesc      `json:"requestBody,omitempty"`
	Responses   []responseDesc `json:"responses"`
	Example     string         `json:"example"`
}

type paramDesc struct {
	Name     string `json:"name"`
	In       string `json:"in"`
	Required bool   `json:"required"`
	field
}

type bodyDesc struct {
	MediaType string  `json:"mediaType"`
	Required  bool    `json:"required"`
	Fields    []field `json:"fields"`
}

type responseDesc struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Schema      string `json:"schema,omitempty"`
}

// field is one property of a schema, flattened to what a caller needs to
// know to fill it in.
type field struct {
	Name        string   `json:"name,omitempty"`
	Type        string   `json:"type"`
	Required    bool     `json:"required"`
	Description string   `json:"description,omitempty"`
	Format      string   `json:"format,omitempty"`
	Enum        []string `json:"enum,omitempty"`
	Pattern     string   `json:"pattern,omitempty"`
	MinLength   *float64 `json:"minLength,omitempty"`
	MaxLength   *float64 `json:"maxLength,omitempty"`
	Minimum     *float64 `json:"minimum,omitempty"`
	Maximum     *float64 `json:"maximum,omitempty"`
	MaxItems    *float64 `json:"maxItems,omitempty"`
}

func (d *document) describe(op *operation) description {
	desc := description{
		OperationId: op.OperationId, Method: op.Method, Path: op.Path,
		Summary: op.Summary, Description: op.Description, Tags: op.Tags,
		Parameters: []paramDesc{}, Responses: []responseDesc{},
	}
	for _, p := range op.Parameters {
		f := d.field("", p.Schema, p.Required)
		f.Description = p.Description
		desc.Parameters = append(desc.Parameters, paramDesc{Name: p.Name, In: p.In, Required: p.Required, field: f})
	}
	if op.RequestBody != nil {
		for mediaType, content := range op.RequestBody.Content {
			body := &bodyDesc{MediaType: mediaType, Required: op.RequestBody.Required, Fields: d.fields(content.Schema)}
			desc.RequestBody = body
			break
		}
	}
	statuses := make([]string, 0, len(op.Responses))
	for status := range op.Responses {
		statuses = append(statuses, status)
	}
	sort.Strings(statuses)
	for _, status := range statuses {
		r := op.Responses[status]
		rd := responseDesc{Status: status, Description: r.Description}
		for _, content := range r.Content {
			rd.Schema = d.schemaName(content.Schema)
			break
		}
		desc.Responses = append(desc.Responses, rd)
	}
	desc.Example = example(op, desc)
	return desc
}

// fields flattens an object schema (resolving a $ref) into its properties.
func (d *document) fields(s schema) []field {
	s = d.resolve(s)
	required := map[string]bool{}
	if list, ok := s["required"].([]any); ok {
		for _, name := range list {
			required[fmt.Sprint(name)] = true
		}
	}
	properties, _ := s["properties"].(map[string]any)
	names := make([]string, 0, len(properties))
	for name := range properties {
		names = append(names, name)
	}
	sort.Strings(names)
	fields := make([]field, 0, len(names))
	for _, name := range names {
		prop, _ := properties[name].(map[string]any)
		fields = append(fields, d.field(name, schema(prop), required[name]))
	}
	return fields
}

func (d *document) field(name string, s schema, required bool) field {
	s = d.resolve(s)
	f := field{Name: name, Required: required, Type: fmt.Sprint(s["type"])}
	if f.Type == "<nil>" {
		f.Type = "object"
	}
	if f.Type == "array" {
		if items, ok := s["items"].(map[string]any); ok {
			f.Type = "array of " + d.field("", schema(items), false).Type
			if enum := stringList(items["enum"]); len(enum) > 0 {
				f.Enum = enum
			}
		}
	}
	f.Description, _ = s["description"].(string)
	f.Format, _ = s["format"].(string)
	f.Pattern, _ = s["pattern"].(string)
	if enum := stringList(s["enum"]); len(enum) > 0 {
		f.Enum = enum
	}
	f.MinLength = number(s["minLength"])
	f.MaxLength = number(s["maxLength"])
	f.Minimum = number(s["minimum"])
	f.Maximum = number(s["maximum"])
	f.MaxItems = number(s["maxItems"])
	return f
}

// resolve follows a $ref into components.schemas.
func (d *document) resolve(s schema) schema {
	ref, ok := s["$ref"].(string)
	if !ok {
		return s
	}
	if target, ok := d.Components.Schemas[strings.TrimPrefix(ref, "#/components/schemas/")]; ok {
		return target
	}
	return s
}

func (d *document) schemaName(s schema) string {
	if ref, ok := s["$ref"].(string); ok {
		return strings.TrimPrefix(ref, "#/components/schemas/")
	}
	if t, ok := s["type"].(string); ok {
		return t
	}
	return ""
}

func (c *cli) schema(ctx context.Context, args []string) error {
	if len(args) != 1 {
		return errors.New("usage: gbf schema NAME")
	}
	s, err := loadSession()
	if err != nil {
		return err
	}
	doc, err := c.loadDocument(ctx, s, false)
	if err != nil {
		return err
	}
	target, ok := doc.Components.Schemas[args[0]]
	if !ok {
		names := make([]string, 0, len(doc.Components.Schemas))
		for name := range doc.Components.Schemas {
			names = append(names, name)
		}
		sort.Strings(names)
		return fmt.Errorf("unknown schema %q; available: %s", args[0], strings.Join(names, ", "))
	}
	return c.printJSON(target)
}

func (c *cli) openapi(ctx context.Context, args []string) error {
	flags := newFlags("openapi", c.out)
	refresh := flags.Bool("refresh", false, "fetch the document again instead of using the cached one")
	if err := flags.Parse(args); err != nil {
		return err
	}
	s, err := loadSession()
	if err != nil {
		return err
	}
	doc, err := c.loadDocument(ctx, s, *refresh)
	if err != nil {
		return err
	}
	var v any
	if err := json.Unmarshal(doc.raw, &v); err != nil {
		return err
	}
	return c.printJSON(v)
}

func example(op *operation, desc description) string {
	parts := []string{"gbf run " + op.OperationId}
	for _, p := range desc.Parameters {
		if p.In == "path" {
			parts = append(parts, "--param "+p.Name+"=...")
		}
	}
	if desc.RequestBody != nil {
		if desc.RequestBody.MediaType == "application/json" {
			parts = append(parts, "--json-body '{...}'")
		} else {
			for _, f := range desc.RequestBody.Fields {
				if f.Required {
					parts = append(parts, "--form "+f.Name+"=...")
				}
			}
		}
	}
	return strings.Join(parts, " ")
}

func fieldDetails(f field) string {
	details := []string{f.Type}
	if f.Format != "" {
		details = append(details, "format "+f.Format)
	}
	if len(f.Enum) > 0 {
		details = append(details, "one of "+strings.Join(f.Enum, ", "))
	}
	if f.Pattern != "" {
		details = append(details, "pattern "+f.Pattern)
	}
	if f.MinLength != nil || f.MaxLength != nil {
		details = append(details, "length "+rangeText(f.MinLength, f.MaxLength))
	}
	if f.Minimum != nil || f.Maximum != nil {
		details = append(details, "range "+rangeText(f.Minimum, f.Maximum))
	}
	if f.MaxItems != nil {
		details = append(details, fmt.Sprintf("at most %v items", *f.MaxItems))
	}
	text := ": " + strings.Join(details, ", ")
	if f.Description != "" {
		text += ". " + f.Description
	}
	return text
}

func rangeText(min, max *float64) string {
	switch {
	case min != nil && max != nil:
		return fmt.Sprintf("%v-%v", *min, *max)
	case min != nil:
		return fmt.Sprintf(">= %v", *min)
	default:
		return fmt.Sprintf("<= %v", *max)
	}
}

func requiredMark(required bool) string {
	if required {
		return ", required"
	}
	return ""
}

func stringList(v any) []string {
	list, ok := v.([]any)
	if !ok {
		return nil
	}
	out := make([]string, 0, len(list))
	for _, item := range list {
		out = append(out, fmt.Sprint(item))
	}
	return out
}

func number(v any) *float64 {
	n, ok := v.(float64)
	if !ok {
		return nil
	}
	return &n
}

func contains(list []string, value string) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-1] + "…"
}
