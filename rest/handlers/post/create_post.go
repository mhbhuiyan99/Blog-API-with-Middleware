package post

import (
	"blogAPI/repo"
	"blogAPI/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreatePost struct {
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Category  string   `json:"category"`
	Tags      []string `json:"tags"`
	Published bool     `json:"published"`
}

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {

	// Get the user ID from the context (set by the authentication middleware)
	userID, ok := r.Context().Value("user_id").(int)
	if !ok {
		util.SendError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req ReqCreatePost
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		util.SendError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createPost, err := h.postRepo.Create(repo.Post{
		UserID:   userID,
		Title:    req.Title,
		Content:  req.Content,
		Category: req.Category,
		Tags:     req.Tags,
		Published: req.Published,
	})
	
	if err != nil {
		fmt.Println("Failed to create post:", err)
		util.SendError(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	util.SendData(w, createPost, http.StatusCreated)
}
