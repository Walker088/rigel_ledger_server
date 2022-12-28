package oauth

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	authUrl  = "https://github.com/login/oauth/authorize"
	tokenURL = "https://github.com/login/oauth/access_token"

	clientID     = ""
	clientSecret = ""
)

type GithubOAuth struct {
	ClientID     string
	ClientSecret string
	AuthURL      string
	TokenURL     string

	RedirectURL string
	Scopes      []string
}

type githubAccessTokenResp struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

func (g GithubOAuth) CallbackHandler(w http.ResponseWriter, r *http.Request) {
	c := r.URL.Query().Get("code")

	// Use the code which is provided by the github oauth server to get the access token
	rqstBodyMap := map[string]string{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"code":          c,
	}
	rqstJson, _ := json.Marshal(rqstBodyMap)
	req, err := http.NewRequest(
		"POST",
		"https://github.com/login/oauth/access_token",
		bytes.NewBuffer(rqstJson),
	)
	if err != nil {
		log.Panic("Request creation failed")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	// Get the response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panic("Request failed")
	}
	respbody, _ := ioutil.ReadAll(resp.Body)
	var ghresp githubAccessTokenResp
	json.Unmarshal(respbody, &ghresp)
}
