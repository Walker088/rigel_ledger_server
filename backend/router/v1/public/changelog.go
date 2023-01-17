package public

import (
	"encoding/json"
	"net/http"
)

type home struct {
	ApiVersion  string `json:"apiVersion"`
	ChangeLogMd string `json:"changeLogMd"`
}

func ChangeLogHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	h := home{
		ApiVersion:  "0.0.1",
		ChangeLogMd: "Implementing",
	}
	response, _ := json.Marshal(h)

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
