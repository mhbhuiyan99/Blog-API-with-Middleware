package post

import (
	"blogAPI/util"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	postID := r.PathValue("id")

	pId, err := strconv.Atoi(postID)
	if err != nil {
		util.SendError(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	post, err := h.postRepo.Get(pId)
	if err != nil {
		fmt.Println(err)
		util.SendError(w, "Failed to get post", http.StatusInternalServerError)
		return
	}

	if post == nil {
		util.SendError(w, "Post not found", http.StatusNotFound)
		return
	}

	util.SendData(w, post, http.StatusOK)
}
