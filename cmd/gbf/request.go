package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/oauth2"
)

// request sends one request as the logged-in user, refreshing and storing
// the token as needed. accept and contentType may be empty.
func (s *session) request(ctx context.Context, method, target, accept, contentType string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, target, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent())
	if accept != "" {
		req.Header.Set("Accept", accept)
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	source := s.tokenSource(ctx)
	// the oauth2 client wraps the transport of the context's HTTP client,
	// so the configured headers ride along with the bearer
	res, err := oauth2.NewClient(context.WithValue(ctx, oauth2.HTTPClient, httpClient()), source).Do(req)
	s.persist(source)
	if err != nil {
		return nil, loginHint(err)
	}
	return res, nil
}

// loginHint points at login when the token could not be refreshed, which
// means the Keycloak session behind it ended.
func loginHint(err error) error {
	var retrieve *oauth2.RetrieveError
	if errors.As(err, &retrieve) {
		return fmt.Errorf("%w; the login has expired, run gbf login", err)
	}
	return err
}

// send performs a JSON API request and prints the response body; a 4xx or
// 5xx status is also an error, so scripts and agents notice.
func (c *cli) send(ctx context.Context, s *session, method, target, contentType string, body io.Reader) error {
	res, err := s.request(ctx, method, target, "application/json", contentType, body)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var pretty bytes.Buffer
	if json.Indent(&pretty, data, "", "  ") == nil {
		data = append(pretty.Bytes(), '\n')
	}
	if _, err := c.out.Write(data); err != nil {
		return err
	}
	if res.StatusCode >= 400 {
		return fmt.Errorf("%s %s: %s", method, strings.TrimPrefix(target, s.AppURL), res.Status)
	}
	return nil
}

// call sends one request to a route as the logged-in user; forms are sent
// url-encoded like the UI does.
func (c *cli) call(ctx context.Context, args []string) error {
	flags := newFlags("call", c.out)
	jsonBody := flags.String("json-body", "", "JSON request body")
	var form keyValues
	flags.Var(&form, "form", "form field KEY=VALUE (repeatable)")
	positional, err := parseInterleaved(flags, args)
	if err != nil {
		return err
	}
	if len(positional) != 2 {
		return errors.New("usage: gbf call METHOD PATH [--json-body BODY] [--form KEY=VALUE]...")
	}
	method := strings.ToUpper(positional[0])
	path := positional[1]
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	body, contentType, err := requestBody(*jsonBody, url.Values(form))
	if err != nil {
		return err
	}
	s, err := loadSession()
	if err != nil {
		return err
	}
	return c.send(ctx, s, method, s.AppURL+path, contentType, body)
}

func requestBody(jsonBody string, form url.Values) (io.Reader, string, error) {
	switch {
	case jsonBody != "" && len(form) > 0:
		return nil, "", errors.New("use either --json-body or --form, not both")
	case jsonBody != "":
		if !json.Valid([]byte(jsonBody)) {
			return nil, "", errors.New("--json-body is not valid JSON")
		}
		return strings.NewReader(jsonBody), "application/json", nil
	case len(form) > 0:
		return strings.NewReader(form.Encode()), "application/x-www-form-urlencoded", nil
	}
	return nil, "", nil
}

func (c *cli) printJSON(v any) error {
	encoder := json.NewEncoder(c.out)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)
	return encoder.Encode(v)
}
