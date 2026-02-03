package user

import (
	"blogAPI/domain"
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

	if err = util.ValidateUsername(req.Username); err != nil {
		util.SendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = util.ValidateEmail(req.Email); err != nil {
		util.SendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = util.ValidatePassword(req.Password); err != nil {
		util.SendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// password hashing
	req.Password = util.HashPassword(req.Password)

	usr, err := h.svc.Create(domain.User{
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
