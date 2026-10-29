package api

import (
	"context"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
)

// DefaultClientID is the public Keycloak client the platform provisions for
// agents; it only allows the device authorization grant.
const DefaultClientID = "agent"

// DefaultAppURL and DefaultIssuerURL are the production platform and its
// Keycloak realm.
const (
	DefaultAppURL    = "https://app.gobugfree.com"
	DefaultIssuerURL = "https://id.gobugfree.com/realms/gbf"
)

// DefaultScopes are the scopes the platform expects on an agent token. The
// token names the user only; the platform resolves their permissions
// itself.
var DefaultScopes = []string{"openid", "profile"}

// Agent authenticates with the OAuth 2.0 device authorization grant (RFC
// 8628) against the platform's Keycloak realm. The token is issued to the
// human who approves the login in a browser: the client then acts as that
// user, within that user's permissions, and every action is recorded under
// their identity.
type Agent struct {
	// AppURL is the platform base URL, e.g. https://app.gobugfree.com.
	AppURL string
	// IssuerURL is the Keycloak realm, e.g. https://id.gobugfree.com/realms/gbf.
	IssuerURL string
	// ClientID defaults to DefaultClientID.
	ClientID string
	// Scopes default to DefaultScopes.
	Scopes []string
	// HTTPClient performs the token requests and carries the API calls; nil
	// uses http.DefaultClient.
	HTTPClient *http.Client
}

func (a Agent) config() *oauth2.Config {
	issuer := strings.TrimSuffix(a.IssuerURL, "/")
	clientID := a.ClientID
	if clientID == "" {
		clientID = DefaultClientID
	}
	scopes := a.Scopes
	if scopes == nil {
		scopes = DefaultScopes
	}
	return &oauth2.Config{
		ClientID: clientID,
		Scopes:   scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:       issuer + "/protocol/openid-connect/auth",
			DeviceAuthURL: issuer + "/protocol/openid-connect/auth/device",
			TokenURL:      issuer + "/protocol/openid-connect/token",
			AuthStyle:     oauth2.AuthStyleInParams,
		},
	}
}

func (a Agent) context(ctx context.Context) context.Context {
	if a.HTTPClient == nil {
		return ctx
	}
	return context.WithValue(ctx, oauth2.HTTPClient, a.HTTPClient)
}

// Login starts the device flow, hands the verification URL and user code to
// show (display them and ask the user to approve in a browser), and blocks
// until the user approves, the device code expires, or ctx ends. Persist the
// returned token to reuse the login across runs.
func (a Agent) Login(ctx context.Context, show func(*oauth2.DeviceAuthResponse)) (*oauth2.Token, error) {
	ctx = a.context(ctx)
	config := a.config()
	response, err := config.DeviceAuth(ctx)
	if err != nil {
		return nil, err
	}
	if show != nil {
		show(response)
	}
	return config.DeviceAccessToken(ctx, response)
}

// TokenSource returns a source that refreshes token when it expires. The
// refresh token is bound to the user's Keycloak session; once that ends the
// source fails and Login is needed again.
func (a Agent) TokenSource(ctx context.Context, token *oauth2.Token) oauth2.TokenSource {
	return a.config().TokenSource(a.context(ctx), token)
}

// NewClient builds an APIClient for AppURL that sends the user's bearer
// token, refreshes it when needed, and asks for JSON on every route.
func (a Agent) NewClient(ctx context.Context, token *oauth2.Token) *APIClient {
	return a.NewClientWithTokenSource(ctx, a.TokenSource(ctx, token))
}

// NewClientWithTokenSource is NewClient for a caller-managed token source,
// for example one that persists refreshed tokens.
func (a Agent) NewClientWithTokenSource(ctx context.Context, source oauth2.TokenSource) *APIClient {
	configuration := NewConfiguration()
	configuration.Servers = ServerConfigurations{{URL: strings.TrimSuffix(a.AppURL, "/")}}
	configuration.HTTPClient = oauth2.NewClient(a.context(ctx), source)
	// every route serves JSON to a bearer caller that accepts it, including
	// routes whose generated operation declares no JSON response
	configuration.AddDefaultHeader("Accept", "application/json")
	return NewAPIClient(configuration)
}
