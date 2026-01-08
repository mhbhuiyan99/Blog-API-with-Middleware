package user

import (
	"blogAPI/util"
	"encoding/json"
	"net/http"
)

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req RequestLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		util.SendError(w, "Invalid request Data", http.StatusBadRequest)
		return
	}

	usr, err := h.userRepo.Find(req.Email, req.Password)
	if err != nil {
		util.SendError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	accessToken, err := util.CreateJWT(h.cnf.JwtSecretKey, util.Payload{
		Sub:      usr.ID,
		IsWriter: usr.IsWriter,
		Username: usr.Username,
	})

	util.SendData(w, accessToken, http.StatusOK)
}
