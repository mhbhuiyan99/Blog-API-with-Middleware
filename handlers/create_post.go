package handlers

import (
	"encoding/json"
	"net/http"
	"blogAPI/database"
	"blogAPI/util"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var newPost database.Post
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newPost)
	if err != nil {
		http.Error(w, "Invalid request payload", 400)
		return
	}

	newPost.ID = len(database.PostList) + 1
	database.PostList = append(database.PostList, newPost)
	util.SendData(w, newPost, 201)
}