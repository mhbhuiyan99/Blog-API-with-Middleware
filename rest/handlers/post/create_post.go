package post

import (
	"blogAPI/repo"
	"blogAPI/util"
	"encoding/json"
	"net/http"
)

type ReqCreatePost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Category string `json:"category"`
	Tags    []string `json:"tags"`
}

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	

	var req ReqCreatePost
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		util.SendError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createPost, err := h.postRepo.Create(repo.Post{
		Title: req.Title,
		Content: req.Content,
		Category: req.Category,
		Tags: req.Tags,
	})
	if err != nil {
		util.SendError(w, "Failed to create post", http.StatusInternalServerError)
		return
	}
	
	util.SendData(w, createPost, http.StatusCreated)
}
