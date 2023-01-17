package protect

import (
	"net/http"
)

func UserHomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	w.WriteHeader(http.StatusOK)
}
