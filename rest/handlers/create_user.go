package handlers

import (
	"blogAPI/database"
	"blogAPI/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser database.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request Data", http.StatusBadRequest)
		return
	}

	createUser := newUser.Store()
	util.SendData(w, createUser, http.StatusCreated)
}