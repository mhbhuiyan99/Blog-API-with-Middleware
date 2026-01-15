package user

import (
	"blogAPI/repo"
	"blogAPI/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreateUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req ReqCreateUser
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		util.SendError(w, "Invalid request Data", http.StatusBadRequest)
		return
	}

	// password hashing
	req.Password = util.HashPassword(req.Password)

	usr, err := h.userRepo.Create(repo.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		util.SendError(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	util.SendData(w, usr, http.StatusCreated)
}
