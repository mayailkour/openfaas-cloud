package handlers

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// OpenFaaSCloudClaims extends standard claims
type OpenFaaSCloudClaims struct {
	// Name is the full name of the user for OIDC
	Name string `json:"name"`

	// AccessToken for use with the GitHub Profile API
	AccessToken string `json:"access_token"`

	// Inherit from standard claims
	jwt.StandardClaims
}

// GitHubAccessToken as issued by GitHub
type GitHubAccessToken struct {
	AccessToken string `json:"access_token"`
}

func buildGitHubURL(config *Config, resource string, scope string) *url.URL {
	authURL := "https://github.com/login/oauth/authorize"
	u, _ := url.Parse(authURL)
	q := u.Query()

	q.Set("scope", scope)
	q.Set("allow_signup", "0")
	q.Set("state", fmt.Sprintf("%d", time.Now().Unix()))
	q.Set("client_id", config.ClientID)

	redirectURI := combineURL(config.ExternalRedirectDomain, "/oauth2/authorized")

	q.Set("redirect_uri", redirectURI)

	u.RawQuery = q.Encode()
	return u
}

func combineURL(a, b string) string {
	if !strings.HasSuffix(a, "/") {
		a = a + "/"
	}
	if strings.HasPrefix(b, "/") {
		b = strings.TrimPrefix(b, "/")
	}

	return a + b
}
