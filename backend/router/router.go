package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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

func New() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	return r
}
