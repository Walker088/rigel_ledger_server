package protected

import (
	"encoding/json"
	"net/http"

	"github.com/Walker088/rigel_ledger_server/src/golang/database/dao"
	"github.com/Walker088/rigel_ledger_server/src/golang/database/pojo"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserHandler struct {
	Dao *dao.UserDao
}

func NewUserHandler(pool *pgxpool.Pool) *UserHandler {
	return &UserHandler{
		Dao: dao.NewUserDao(pool),
	}
}

func (u *UserHandler) GetUserCompleteHandler(w http.ResponseWriter, r *http.Request) {
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

	var userInfo pojo.UpdateUserInfo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInfo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if updatedUinfo, err := u.Dao.Update(uid, &userInfo); err == nil {
		response, _ := json.Marshal(updatedUinfo)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
}
