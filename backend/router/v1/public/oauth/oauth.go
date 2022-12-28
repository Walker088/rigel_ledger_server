package oauth

import "net/http"

type OAuthInterface interface {
	CallbackHandler(w http.ResponseWriter, r *http.Request)
}

type OAuthProp struct {
	ClientID     string
	ClientSecret string
	AuthURL      string
	TokenURL     string

	RedirectURL string
	Scopes      []string
}
