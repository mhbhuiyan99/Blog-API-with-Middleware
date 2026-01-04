package user

import (
	"blogAPI/config"
	"blogAPI/database"
	"blogAPI/util"
	"encoding/json"
	"net/http"
)

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin RequestLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)
	if err != nil {
		http.Error(w, "Invalid request Data", http.StatusBadRequest)
		return
	}

	usr := database.Find(reqLogin.Email, reqLogin.Password)
	if usr == nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	cnf := config.GetConfig()

	accessToken, err := util.CreateJWT(cnf.JwtSecretKey, util.Payload{
		Sub:      usr.ID,
		IsWriter: usr.IsWriter,
		Username: usr.Username,
	})

	util.SendData(w, accessToken, http.StatusOK)
}
