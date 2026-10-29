package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

// envHeaderPrefix marks the environment variables gbf sends as request
// headers: the name after the prefix becomes the header name with
// underscores as dashes, so GBF_HEADER_CF_ACCESS_CLIENT_ID sends
// CF-Access-Client-Id. That reaches a platform (and its identity provider)
// behind Cloudflare Access, an API gateway, or any other proxy that wants a
// header, without a flag on every command.
const envHeaderPrefix = "GBF_HEADER_"

type headerTransport struct {
	base   http.RoundTripper
	header http.Header
}

func (t headerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req = req.Clone(req.Context())
	for name, values := range t.header {
		req.Header[name] = values
	}
	return t.base.RoundTrip(req)
}

// httpClient is the client behind every request gbf makes; it adds the
// headers configured through the environment.
func httpClient() *http.Client {
	header := http.Header{}
	for _, h := range envHeaders() {
		header.Set(h.name, h.value)
	}
	if len(header) == 0 {
		return &http.Client{}
	}
	return &http.Client{Transport: headerTransport{base: http.DefaultTransport, header: header}}
}

// envHeader is one GBF_HEADER_* variable: the variable it came from, the
// header name it sends, and its value.
type envHeader struct{ variable, name, value string }

// envHeaders reads the GBF_HEADER_* variables, in the environment's order.
// An empty value configures no header, so a wrapper can drop one by
// exporting it empty.
func envHeaders() []envHeader {
	var headers []envHeader
	for _, entry := range os.Environ() {
		variable, value, ok := strings.Cut(entry, "=")
		if !ok || !strings.HasPrefix(variable, envHeaderPrefix) || value == "" {
			continue
		}
		name := strings.ReplaceAll(strings.TrimPrefix(variable, envHeaderPrefix), "_", "-")
		if name == "" {
			continue
		}
		headers = append(headers, envHeader{variable: variable, name: name, value: value})
	}
	return headers
}

// checkHeaderEnv rejects a value the transport could not send, instead of
// failing every request with a write error from deep inside net/http.
func checkHeaderEnv() error {
	for _, h := range envHeaders() {
		if strings.ContainsAny(h.value, "\r\n\x00") {
			return fmt.Errorf("%s holds a newline, which cannot be sent as the %s header", h.variable, h.name)
		}
	}
	return nil
}
