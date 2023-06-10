package public

import (
	"encoding/json"
	"net/http"
)

type HomeInfo struct {
	ApiVersion  string `json:"apiVersion"`
	ChangeLogMd string `json:"changeLogMd"`
}

func NewHomeInfo() *HomeInfo {
	return &HomeInfo{
		ApiVersion:  "0.0.1",
		ChangeLogMd: "Implementing",
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
