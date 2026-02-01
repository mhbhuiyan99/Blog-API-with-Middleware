package post

import (
	"blogAPI/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetDrafts(w http.ResponseWriter, r *http.Request) {
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

	userId := r.Context().Value("user_id").(int)

	cnt, err := h.svc.CountDrafts(userId)
	if err != nil {
		util.SendError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	draftList, err := h.svc.Draft(int64(userId), page, limit)
	if err != nil {
		util.SendError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	util.SendPage(w, draftList, page, limit, cnt)
}