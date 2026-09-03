package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/url"
	"strings"
)

// runOperation invokes an operation of the live document by id: the path
// template is filled from --param, the query from --query, and the body is
// sent the way the operation declares it (form fields or JSON).
func (c *cli) runOperation(ctx context.Context, args []string) error {
	flags := newFlags("run", c.out)
	var params, query, form keyValues
	flags.Var(&params, "param", "path parameter KEY=VALUE (repeatable)")
	flags.Var(&query, "query", "query parameter KEY=VALUE (repeatable)")
	flags.Var(&form, "form", "form field KEY=VALUE (repeatable)")
	jsonBody := flags.String("json-body", "", "JSON request body")
	dryRun := flags.Bool("dry-run", false, "print the request instead of sending it")
	positional, err := parseInterleaved(flags, args)
	if err != nil {
		return err
	}
	if len(positional) != 1 {
		return errors.New("usage: gbf run OPERATION_ID [--param KEY=VALUE]... [--query KEY=VALUE]... [--form KEY=VALUE]... [--json-body BODY] [--dry-run]")
	}

	s, err := loadSession()
	if err != nil {
		return err
	}
	doc, op, err := c.find(ctx, s, positional[0])
	if err != nil {
		return err
	}

	path, err := fillPath(op, url.Values(params))
	if err != nil {
		return err
	}
	target := s.AppURL + path
	if len(query) > 0 {
		target += "?" + url.Values(query).Encode()
	}
	body, contentType, err := operationBody(doc, op, *jsonBody, url.Values(form))
	if err != nil {
		return err
	}

	if *dryRun {
		return c.printDryRun(op, target, contentType, body)
	}
	return c.send(ctx, s, op.Method, target, contentType, body)
}

// fillPath substitutes the path parameters and reports the missing ones.
func fillPath(op *operation, params url.Values) (string, error) {
	path := op.Path
	var missing []string
	for _, p := range op.Parameters {
		if p.In != "path" {
			continue
		}
		value := params.Get(p.Name)
		if value == "" {
			missing = append(missing, p.Name)
			continue
		}
		path = strings.ReplaceAll(path, "{"+p.Name+"}", url.PathEscape(value))
	}
	if len(missing) > 0 {
		return "", fmt.Errorf("%s needs --param %s", op.OperationId, strings.Join(missing, ", --param "))
	}
	return path, nil
}

// operationBody builds the body the operation declares and rejects the
// other kind, so an agent learns the contract from the error.
func operationBody(doc *document, op *operation, jsonBody string, form url.Values) (io.Reader, string, error) {
	if op.RequestBody == nil {
		if jsonBody != "" || len(form) > 0 {
			return nil, "", fmt.Errorf("%s takes no request body", op.OperationId)
		}
		return nil, "", nil
	}
	_, wantsJSON := op.RequestBody.Content["application/json"]
	var mediaType string
	var bodySchema schema
	for m, content := range op.RequestBody.Content {
		mediaType, bodySchema = m, content.Schema
		break
	}
	switch {
	case wantsJSON && len(form) > 0:
		return nil, "", fmt.Errorf("%s expects a JSON body: use --json-body", op.OperationId)
	case !wantsJSON && jsonBody != "":
		return nil, "", fmt.Errorf("%s expects form fields (%s): use --form KEY=VALUE", op.OperationId, mediaType)
	case jsonBody == "" && len(form) == 0:
		if !op.RequestBody.Required {
			return nil, "", nil
		}
		var required []string
		for _, f := range doc.fields(bodySchema) {
			if f.Required {
				required = append(required, f.Name)
			}
		}
		if wantsJSON {
			return nil, "", fmt.Errorf("%s needs a JSON body (--json-body); required fields: %s", op.OperationId, strings.Join(required, ", "))
		}
		return nil, "", fmt.Errorf("%s needs form fields; required: --form %s", op.OperationId, strings.Join(required, " --form "))
	}
	return requestBody(jsonBody, form)
}

func (c *cli) printDryRun(op *operation, target, contentType string, body io.Reader) error {
	var text string
	if body != nil {
		data, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		text = string(data)
	}
	headers := map[string]string{"Accept": "application/json", "Authorization": "Bearer <token>", "User-Agent": userAgent()}
	if contentType != "" {
		headers["Content-Type"] = contentType
	}
	if c.json {
		return c.printJSON(map[string]any{"operationId": op.OperationId, "method": op.Method, "url": target, "headers": headers, "body": text})
	}
	fmt.Fprintf(c.out, "%s %s\n", op.Method, target)
	for _, name := range []string{"Accept", "Authorization", "Content-Type", "User-Agent"} {
		if value, ok := headers[name]; ok {
			fmt.Fprintf(c.out, "%s: %s\n", name, value)
		}
	}
	if text != "" {
		fmt.Fprintf(c.out, "\n%s\n", text)
	}
	return nil
}
