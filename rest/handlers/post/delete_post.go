package post

import (
	"blogAPI/util"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	postID := r.PathValue("id")

	pId, err := strconv.Atoi(postID)
	if err != nil {
		fmt.Println(err)
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
	existingPost, err := h.svc.Get(pId)
	if err != nil {
		fmt.Println(err)
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

	err = h.svc.Delete(pId)
	if err != nil {
		fmt.Println(err)
		util.SendError(w, "Failed to delete post", http.StatusInternalServerError)
		return
	}
	util.SendData(w, "Successfully deleted post", http.StatusOK)
}
