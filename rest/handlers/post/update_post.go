package post

import (
	"blogAPI/repo"
	"blogAPI/util"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/lib/pq"
)

type ReqUpdatePost struct {
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Category string   `json:"category"`
	ImgURL   string   `json:"img_url"`
	Tags     []string `json:"tags"`
	Published bool     `json:"published"`
}

func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {

	postID := r.PathValue("id")

	pId, err := strconv.Atoi(postID)
	if err != nil {
		util.SendError(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	// Get authenticated user id from context
	userID, ok := r.Context().Value("user_id").(int)
	if !ok {
		util.SendError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Check the post exists and belongs to the authenticated user
	existingPost, err := h.postRepo.Get(pId)
	if err != nil {
		util.SendError(w, "Failed to fetch post", http.StatusInternalServerError)
		return
	}
	if existingPost == nil {
		util.SendError(w, "Post not found", http.StatusNotFound)
		return
	}
	if existingPost.UserID != userID {
		util.SendError(w, "Forbidden", http.StatusForbidden)
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
		ID:        pId,
		UserID:    userID,
		Title:     req.Title,
		Content:   req.Content,
		Category:  req.Category,
		ImgURL:    req.ImgURL,
		Tags:      pq.StringArray(req.Tags),
		Published: req.Published,
	})

	if err != nil {
		util.SendError(w, "Failed to update post", http.StatusInternalServerError)
		return
	}

	util.SendData(w, "Post updated successfully", http.StatusOK)
}
