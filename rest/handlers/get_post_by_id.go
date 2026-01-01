package handlers

import (
	"net/http"
	"strconv"
	"blogAPI/database"
	"blogAPI/util"
)

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	postID := r.PathValue("postId")

	pId, err := strconv.Atoi(postID)
	if err != nil {
		http.Error(w, "Invalid post ID", 400)
		return
	}

	for _, post := range database.List() {
		if post.ID == pId {
			util.SendData(w, post, 200)
			return
		}
	}

	util.SendData(w, "Data not found", 404)
}