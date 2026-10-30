package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"golang.org/x/oauth2"
)

// fakeIssuer serves the two Keycloak endpoints the device flow needs.
func fakeIssuer(t *testing.T) (*httptest.Server, *[]string) {
	t.Helper()
	var grants []string
	mux := http.NewServeMux()
	mux.HandleFunc("POST /realms/gobugfree/protocol/openid-connect/auth/device", func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("client_id") != "agent" {
			http.Error(w, `{"error":"invalid_client"}`, http.StatusUnauthorized)
			return
		}
		writeJSON(w, map[string]any{
			"device_code":               "device-1",
			"user_code":                 "ABCD-EFGH",
			"verification_uri":          "https://iam.example/device",
			"verification_uri_complete": "https://iam.example/device?user_code=ABCD-EFGH",
			"expires_in":                600,
			"interval":                  1,
		})
	})
	mux.HandleFunc("POST /realms/gobugfree/protocol/openid-connect/token", func(w http.ResponseWriter, r *http.Request) {
		grant := r.FormValue("grant_type")
		grants = append(grants, grant)
		switch grant {
		case "urn:ietf:params:oauth:grant-type:device_code":
			if r.FormValue("device_code") != "device-1" {
				http.Error(w, `{"error":"invalid_grant"}`, http.StatusBadRequest)
				return
			}
			writeJSON(w, map[string]any{"access_token": "access-1", "refresh_token": "refresh-1", "token_type": "Bearer", "expires_in": 3600})
		case "refresh_token":
			if r.FormValue("refresh_token") != "refresh-1" {
				http.Error(w, `{"error":"invalid_grant"}`, http.StatusBadRequest)
				return
			}
			writeJSON(w, map[string]any{"access_token": "access-2", "refresh_token": "refresh-2", "token_type": "Bearer", "expires_in": 3600})
		default:
			http.Error(w, `{"error":"unsupported_grant_type"}`, http.StatusBadRequest)
		}
	})
	server := httptest.NewServer(mux)
	t.Cleanup(server.Close)
	return server, &grants
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}

func TestLogin(t *testing.T) {
	issuer, grants := fakeIssuer(t)
	agent := Agent{IssuerURL: issuer.URL + "/realms/gobugfree/"}

	var shown *oauth2.DeviceAuthResponse
	token, err := agent.Login(context.Background(), func(r *oauth2.DeviceAuthResponse) { shown = r })
	if err != nil {
		t.Fatalf("login: %v", err)
	}
	if shown == nil || shown.UserCode != "ABCD-EFGH" || shown.VerificationURIComplete == "" {
		t.Fatalf("verification details not shown: %+v", shown)
	}
	if token.AccessToken != "access-1" || token.RefreshToken != "refresh-1" {
		t.Fatalf("unexpected token %+v", token)
	}
	if len(*grants) != 1 || (*grants)[0] != "urn:ietf:params:oauth:grant-type:device_code" {
		t.Fatalf("unexpected grants %v", *grants)
	}
}

func TestNewClientSendsBearerAndRefreshes(t *testing.T) {
	issuer, grants := fakeIssuer(t)

	var authorization, accept string
	app := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization = r.Header.Get("Authorization")
		accept = r.Header.Get("Accept")
		writeJSON(w, map[string]any{"items": []any{}})
	}))
	t.Cleanup(app.Close)

	agent := Agent{AppURL: app.URL + "/", IssuerURL: issuer.URL + "/realms/gobugfree"}
	expired := &oauth2.Token{AccessToken: "access-1", RefreshToken: "refresh-1", Expiry: time.Now().Add(-time.Minute)}
	client := agent.NewClient(context.Background(), expired)

	configuration := client.GetConfig()
	if got := configuration.Servers[0].URL; got != app.URL {
		t.Fatalf("server url %q, want %q", got, app.URL)
	}
	if configuration.DefaultHeader["Accept"] != "application/json" {
		t.Fatalf("Accept default header missing: %v", configuration.DefaultHeader)
	}

	response, err := configuration.HTTPClient.Get(app.URL + "/engagements")
	if err != nil {
		t.Fatalf("request: %v", err)
	}
	response.Body.Close()
	if authorization != "Bearer access-2" {
		t.Fatalf("authorization %q, want the refreshed bearer", authorization)
	}
	if accept != "" {
		t.Fatalf("the raw HTTP client sets no Accept header, got %q", accept)
	}
	if len(*grants) != 1 || (*grants)[0] != "refresh_token" {
		t.Fatalf("expected one refresh, got %v", *grants)
	}
}
