package post

import (
	"blogAPI/repo"
	"blogAPI/util"
	"encoding/json"
	"net/http"
	"strconv"
)

type ReqUpdatePost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Category string `json:"category"`
	Tags    []string `json:"tags"`
}

func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {

	postID := r.PathValue("id")

	pId, err := strconv.Atoi(postID)
	if err != nil {
		util.SendError(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	var req ReqUpdatePost
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		util.SendError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	_, err = h.postRepo.Update(repo.Post{
		ID:      pId,
		Title:   req.Title,
		Content: req.Content,
		Category: req.Category,
		Tags:    req.Tags,
	})

	if err != nil {
		util.SendError(w, "Failed to update post", http.StatusInternalServerError)
		return
	}

	util.SendData(w, "Post updated successfully", http.StatusOK)
}