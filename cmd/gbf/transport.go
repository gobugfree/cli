package main

import (
	"errors"
	"net/http"
	"os"
)

// Cloudflare Access service tokens let gbf reach a platform (and its
// identity provider) behind Zero Trust: when both variables are set, every
// request carries the two headers Access expects.
const (
	envAccessClientId     = "CF_ACCESS_CLIENT_ID"
	envAccessClientSecret = "CF_ACCESS_CLIENT_SECRET"
)

type accessTransport struct {
	base         http.RoundTripper
	clientId     string
	clientSecret string
}

func (t accessTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req = req.Clone(req.Context())
	req.Header.Set("CF-Access-Client-Id", t.clientId)
	req.Header.Set("CF-Access-Client-Secret", t.clientSecret)
	return t.base.RoundTrip(req)
}

// httpClient is the client behind every request gbf makes; it adds the
// Cloudflare Access headers when the service token is configured.
func httpClient() *http.Client {
	id, secret := os.Getenv(envAccessClientId), os.Getenv(envAccessClientSecret)
	if id == "" || secret == "" {
		return &http.Client{}
	}
	return &http.Client{Transport: accessTransport{base: http.DefaultTransport, clientId: id, clientSecret: secret}}
}

// checkAccessEnv rejects a half-configured service token early, instead of
// letting Access answer every request with its login page.
func checkAccessEnv() error {
	id, secret := os.Getenv(envAccessClientId), os.Getenv(envAccessClientSecret)
	if (id == "") != (secret == "") {
		return errors.New(envAccessClientId + " and " + envAccessClientSecret + " must be set together")
	}
	return nil
}
