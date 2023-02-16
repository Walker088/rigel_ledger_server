package protect

import (
	"net/http"
)

func UserHomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	//uid := chi.URLParam(r, "userId")

	w.WriteHeader(http.StatusOK)
}
