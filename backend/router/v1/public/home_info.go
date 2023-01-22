package public

import (
	"encoding/json"
	"net/http"
)

type HomeInfo struct {
	ApiVersion  string `json:"apiVersion"`
	ChangeLogMd string `json:"changeLogMd"`

	OauthOtps map[string]struct {
		Title string `json:"title"`
		Link  string `json:"link"`
	} `json:"oauthOtps"`
}

func NewHomeInfo(ghLink string) *HomeInfo {
	return &HomeInfo{
		ApiVersion:  "0.0.1",
		ChangeLogMd: "Implementing",
		OauthOtps: map[string]struct {
			Title string `json:"title"`
			Link  string `json:"link"`
		}{
			"github": {Title: "Github", Link: ghLink},
		},
	}
}

func (hi *HomeInfo) HomeInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	response, _ := json.Marshal(hi)

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
