package post

import (
	"blogAPI/util"
	"net/http"
)

func (h *Handler) GetPosts(w http.ResponseWriter, r *http.Request) {
	postList, err := h.postRepo.List()
	if err != nil {
		util.SendError(w, "Failed to get posts", http.StatusInternalServerError)
		return
	}

	util.SendData(w, postList, http.StatusOK)
}