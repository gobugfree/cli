package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gobugfree/cli/v3/api"
	"golang.org/x/oauth2"
)

// session is what login stores: where the platform is and the user's
// token. It is private to the user (directory 0700, file 0600).
type session struct {
	AppURL    string        `json:"appUrl"`
	IssuerURL string        `json:"issuerUrl"`
	ClientID  string        `json:"clientId,omitempty"`
	Token     *oauth2.Token `json:"token"`
}

func sessionPath() string {
	if path := os.Getenv("GBF_SESSION_FILE"); path != "" {
		return path
	}
	dir := os.Getenv("XDG_CONFIG_HOME")
	if dir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			home = "."
		}
		dir = filepath.Join(home, ".config")
	}
	return filepath.Join(dir, "gbf", "session.json")
}

func loadSession() (*session, error) {
	data, err := os.ReadFile(sessionPath())
	if errors.Is(err, os.ErrNotExist) {
		return nil, errors.New("not logged in: run gbf login")
	}
	if err != nil {
		return nil, err
	}
	var s session
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, fmt.Errorf("session file %s: %w", sessionPath(), err)
	}
	if s.AppURL == "" || s.IssuerURL == "" || s.Token == nil {
		return nil, fmt.Errorf("session file %s is incomplete: run gbf login", sessionPath())
	}
	return &s, nil
}

func (s *session) save() error {
	path := sessionPath()
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o600)
}

func (s *session) agent() api.Agent {
	return api.Agent{AppURL: s.AppURL, IssuerURL: s.IssuerURL, ClientID: s.ClientID, HTTPClient: httpClient()}
}

func (s *session) tokenSource(ctx context.Context) oauth2.TokenSource {
	return s.agent().TokenSource(ctx, s.Token)
}

// persist stores the token the source holds now, so a refresh survives the
// process. Failures are reported, not fatal: the call itself succeeded.
func (s *session) persist(source oauth2.TokenSource) {
	token, err := source.Token()
	if err != nil || token.AccessToken == s.Token.AccessToken {
		return
	}
	s.Token = token
	if err := s.save(); err != nil {
		fmt.Fprintln(os.Stderr, "gbf: could not store the refreshed token:", err)
	}
}
