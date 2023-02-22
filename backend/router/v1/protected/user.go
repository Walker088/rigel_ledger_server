package protected

import (
	"encoding/json"
	"net/http"

	"github.com/Walker088/rigel_ledger_server/backend/database/dao"
	"github.com/Walker088/rigel_ledger_server/backend/database/pojo"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	Dao *dao.UserDao
}

func (u *UserHandler) GetUserCompleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	uid := chi.URLParam(r, "userId")
	userInfo, err := u.Dao.GetComplete(uid)
	if err == nil {
		response, _ := json.Marshal(userInfo)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}

	w.WriteHeader(http.StatusInternalServerError)
}

func (u *UserHandler) GetUserBasicHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	uid := chi.URLParam(r, "userId")
	if userInfo, err := u.Dao.GetBasic(uid); err == nil {
		response, _ := json.Marshal(userInfo)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}

	w.WriteHeader(http.StatusInternalServerError)
}

func (u *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "userId")

	var userInfo pojo.UserInfo
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userInfo)
	if err == nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if updatedUinfo, err := u.Dao.Update(uid, &userInfo); err == nil {
		response, _ := json.Marshal(updatedUinfo)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
	w.WriteHeader(http.StatusInternalServerError)
}
