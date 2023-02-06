package oauth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Walker088/rigel_ledger_server/backend/jwt"
	"go.uber.org/zap"
)

const (
	//authUrl  = "https://github.com/login/oauth/authorize"
	tokenURL = "https://github.com/login/oauth/access_token"
)

type GithubOAuth struct {
	ClientID     string
	ClientSecret string
	//AuthURL      string
	TokenURL string

	jwtEngine *jwt.JwtEngine
	Logger    *zap.SugaredLogger
}

type githubAccessTokenResp struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type githubUserInfoResp struct {
	GhUid       int    `json:"id"`
	GhUserHome  string `json:"html_url"`
	UserId      string `json:"login"`
	UserMail    string `json:"email"`
	UserName    string `json:"name"`
	UserCompany string `json:"company"`
	AvatarUrl   string `json:"avatar_url"`
}

func New(clientId string, clientSecret string, logger *zap.SugaredLogger, jwtEngine *jwt.JwtEngine) *GithubOAuth {
	return &GithubOAuth{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		//AuthURL:      authUrl,
		TokenURL:  tokenURL,
		Logger:    logger,
		jwtEngine: jwtEngine,
	}
}

func (g *GithubOAuth) GithubLogin(w http.ResponseWriter, r *http.Request) {
	c := r.URL.Query().Get("code")
	ghresp, err := g.getAccessToken(c)
	if err != nil {
		g.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ghUserInfo, err := g.getUserInfo(ghresp)
	if err != nil {
		g.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tokens, err := g.jwtEngine.GenTokens(ghUserInfo.UserId)
	if err != nil {
		g.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	marshaledTokens, _ := json.Marshal(tokens)
	w.WriteHeader(http.StatusOK)
	w.Write(marshaledTokens)
}

func (g *GithubOAuth) getUserInfo(ghresp *githubAccessTokenResp) (*githubUserInfoResp, error) {
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
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		g.Logger.Errorf("Get Github UserInfo error, status: %d, body: %s", resp.StatusCode, string(body))
		return nil, fmt.Errorf("GET github UserInfo Error, status code: %d", resp.StatusCode)
	}
	var u githubUserInfoResp
	if err := json.NewDecoder(resp.Body).Decode(&u); err != nil {
		g.Logger.Errorf("Could not parse JSON response: %v", err)
		return nil, fmt.Errorf("ERROR: can not parse JSON response: %v", err)
	}
	return &u, nil
}

func (g *GithubOAuth) getAccessToken(code string) (*githubAccessTokenResp, error) {
	rqstBodyMap := map[string]string{
		"client_id":     g.ClientID,
		"client_secret": g.ClientSecret,
		"code":          code,
	}
	rqstJson, _ := json.Marshal(rqstBodyMap)
	req, err := http.NewRequest(
		http.MethodPost,
		g.TokenURL,
		bytes.NewBuffer(rqstJson),
	)
	if err != nil {
		g.Logger.Error("Github Oauth Request creation failed")
		return nil, errors.New("ERROR: Github Oauth Request creation failed")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		g.Logger.Error("Github Oauth Request failed")
		return nil, errors.New("ERROR: Github Oauth Request failed")
	}
	defer resp.Body.Close()
	var ghresp githubAccessTokenResp
	if err := json.NewDecoder(resp.Body).Decode(&ghresp); err != nil {
		g.Logger.Errorf("Could not parse JSON response: %v", err)
		return nil, fmt.Errorf("ERROR: Could not parse JSON response: %v", err)
	}
	return &ghresp, nil
}
