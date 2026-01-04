package post

import (
	"blogAPI/database"
	"blogAPI/util"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	

	var newPost database.Post
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newPost)
	if err != nil {
		http.Error(w, "Invalid request payload", 400)
		return
	}

	newPost.ID = len(database.List()) + 1
	database.Store(newPost)
	util.SendData(w, newPost, 201)
}
