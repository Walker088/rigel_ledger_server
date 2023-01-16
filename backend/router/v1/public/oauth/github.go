package oauth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

const (
	authUrl  = "https://github.com/login/oauth/authorize"
	tokenURL = "https://github.com/login/oauth/access_token"
)

type GithubOAuth struct {
	ClientID     string
	ClientSecret string
	AuthURL      string
	TokenURL     string

	Logger *zap.SugaredLogger
}

type githubAccessTokenResp struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

func New(clientId string, clientSecret string, logger *zap.SugaredLogger) *GithubOAuth {
	return &GithubOAuth{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		AuthURL:      authUrl,
		TokenURL:     tokenURL,
		Logger:       logger,
	}
}

func (g *GithubOAuth) LoginHandler(w http.ResponseWriter, r *http.Request) {
	redirectURL := fmt.Sprintf(
		"%s?client_id=%s&redirect_uri=%s&scope=%s",
		g.AuthURL,
		g.ClientID,
		"http://localhost:8000/login/github/callback",
		"user:email",
	)
	http.Redirect(w, r, redirectURL, http.StatusMovedPermanently)
}

func (g *GithubOAuth) CallbackHandler(w http.ResponseWriter, r *http.Request) {
	c := r.URL.Query().Get("code")

	rqstBodyMap := map[string]string{
		"client_id":     g.ClientID,
		"client_secret": g.ClientSecret,
		"code":          c,
	}
	rqstJson, _ := json.Marshal(rqstBodyMap)
	req, err := http.NewRequest(
		http.MethodPost,
		g.TokenURL,
		bytes.NewBuffer(rqstJson),
	)
	if err != nil {
		g.Logger.Error("Github Oauth Request creation failed")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		g.Logger.Error("Github Oauth Request failed")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	var ghresp githubAccessTokenResp
	if err := json.NewDecoder(resp.Body).Decode(&ghresp); err != nil {
		g.Logger.Errorf("Could not parse JSON response: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	final, err := g.getUserInfo(ghresp)
	if err != nil {
		g.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(final)
}

func (g *GithubOAuth) getUserInfo(ghresp githubAccessTokenResp) ([]byte, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		"https://api.github.com/user",
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("GET GITHUB USERINFO ERR: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", ghresp.AccessToken))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GET GITHUB USERINFO ERR:  %v", err)
	}
	defer resp.Body.Close()
	return []byte(""), nil
}
