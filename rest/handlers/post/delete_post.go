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

	err = h.postRepo.Delete(pId)
	if err != nil {
		fmt.Println(err)
		util.SendError(w, "Failed to delete post", http.StatusInternalServerError)
		return
	}
	util.SendData(w, "Successfully deleted post", http.StatusOK)
}
