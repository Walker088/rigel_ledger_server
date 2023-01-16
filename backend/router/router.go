package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/Walker088/rigel_ledger_server/backend/config"
	"github.com/Walker088/rigel_ledger_server/backend/router/v1/public/oauth"
)

type home struct {
	ApiVersion  string `json:"apiVersion"`
	ChangeLogMd string `json:"changeLogMd"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	h := home{
		ApiVersion:  "0.0.1",
		ChangeLogMd: "Implementing",
	}
	response, _ := json.Marshal(h)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func New(c *config.GithubOAuthConfig, logger *zap.SugaredLogger) *mux.Router {
	var auth = oauth.New(c.OauthGithubClientId, c.OauthGithubClientSecret, logger)

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	r.HandleFunc("/oauth/github/login", auth.LoginHandler)
	r.HandleFunc("/oauth/github/callback", auth.CallbackHandler)
	return r
}
