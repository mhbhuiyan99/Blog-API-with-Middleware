package post

import (
	"blogAPI/util"
	"net/http"
	"strconv"
)


func (h *Handler) GetPosts(w http.ResponseWriter, r *http.Request) {

	reqQuery := r.URL.Query()
	pageAsStr := reqQuery.Get("page")
	limitAsStr := reqQuery.Get("limit")

	page, _ := strconv.ParseInt(pageAsStr, 10, 32)
	limit, _ := strconv.ParseInt(limitAsStr, 10, 32)

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	cnt, err := h.svc.Count()
	if err != nil {
		util.SendData(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	postList, err := h.svc.List(page, limit)
	if err != nil {
		util.SendError(w, "Failed to get posts", http.StatusInternalServerError)
		return
	}

	util.SendPage(w, postList, page, limit, cnt)
}
