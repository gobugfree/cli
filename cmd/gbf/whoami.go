package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

// identity is what the access token says about the user. The claims are
// read, not verified: the platform verifies them on every request, this
// only tells the agent whose permissions it has.
type identity struct {
	Username string    `json:"username"`
	Subject  string    `json:"subject"`
	Email    string    `json:"email,omitempty"`
	ClientId string    `json:"clientId,omitempty"`
	Expires  time.Time `json:"expires"`
}

func identityFromToken(accessToken string) (*identity, error) {
	parts := strings.Split(accessToken, ".")
	if len(parts) != 3 {
		return nil, errors.New("the access token is not a JWT")
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("decode token claims: %w", err)
	}
	var claims struct {
		Username string `json:"preferred_username"`
		Subject  string `json:"sub"`
		Email    string `json:"email"`
		ClientId string `json:"azp"`
		Expires  int64  `json:"exp"`
	}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return nil, fmt.Errorf("decode token claims: %w", err)
	}
	return &identity{
		Username: claims.Username, Subject: claims.Subject, Email: claims.Email,
		ClientId: claims.ClientId, Expires: time.Unix(claims.Expires, 0).UTC(),
	}, nil
}

func (c *cli) whoami() error {
	s, err := loadSession()
	if err != nil {
		return err
	}
	who, err := identityFromToken(s.Token.AccessToken)
	if err != nil {
		return err
	}
	if c.json {
		return c.printJSON(who)
	}
	fmt.Fprintf(c.out, "User:    %s (%s)\n", who.Username, who.Subject)
	if who.Email != "" {
		fmt.Fprintf(c.out, "Email:   %s\n", who.Email)
	}
	fmt.Fprintf(c.out, "Client:  %s\nExpires: %s\n", who.ClientId, who.Expires.Format(time.RFC3339))
	return nil
}
