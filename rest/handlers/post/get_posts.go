package post

import (
	"blogAPI/domain"
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

	cnt, err := h.svc.TotalCount(domain.CountOptions{}, true)
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

func (h *Handler) GetUserPosts(w http.ResponseWriter, r *http.Request) {
	reqQuery := r.URL.Query()
	pageAsStr := reqQuery.Get("page")
	limitAsStr := reqQuery.Get("limit")
	userId := r.PathValue("userId")

	page, _ := strconv.ParseInt(pageAsStr, 10, 32)
	limit, _ := strconv.ParseInt(limitAsStr, 10, 32)
	uId, err := strconv.Atoi(userId)

	if err != nil {
		util.SendError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	cnt, err := h.svc.TotalCount(
		domain.CountOptions{
			UserID: uId, 
		}, true,
	)

	if err != nil {
		util.SendError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	postList, err := h.svc.Published(int64(uId), page, limit)
	if err != nil {
		util.SendError(w, "Failed to get posts", http.StatusInternalServerError)
		return
	}

	util.SendPage(w, postList, page, limit, cnt)
}

func (h *Handler) GetByCategory(w http.ResponseWriter, r *http.Request) {
	reqQuery := r.URL.Query()
	pageAsStr := reqQuery.Get("page")
	limitAsStr := reqQuery.Get("limit")
	category := r.PathValue("category")

	page, _ := strconv.ParseInt(pageAsStr, 10, 32)
	limit, _ := strconv.ParseInt(limitAsStr, 10, 32)

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	cnt, err := h.svc.TotalCount(
		domain.CountOptions{
			Category: category,
		}, true,
	)

	if err != nil {
		util.SendError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	postList, err := h.svc.GetByCategory(category, page, limit)
	if err != nil {
		util.SendError(w, "Failed to get posts", http.StatusInternalServerError)
		return
	}

	util.SendPage(w, postList, page, limit, cnt)
}

func (h *Handler) GetByTag(w http.ResponseWriter, r *http.Request) {
	reqQuery := r.URL.Query()
	pageAsStr := reqQuery.Get("page")
	limitAsStr := reqQuery.Get("limit")
	tag := r.PathValue("tag")

	page, _ := strconv.ParseInt(pageAsStr, 10, 32)
	limit, _ := strconv.ParseInt(limitAsStr, 10, 32)

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	cnt, err := h.svc.TotalCount(
		domain.CountOptions{
			Tag: tag,
		}, true,
	)

	if err != nil {
		util.SendError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	postList, err := h.svc.GetByTag(tag, page, limit)
	if err != nil {
		util.SendError(w, "Failed to get posts", http.StatusInternalServerError)
		return
	}

	util.SendPage(w, postList, page, limit, cnt)
}

func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	reqQuery := r.URL.Query()
	pageAsStr := reqQuery.Get("page")
	limitAsStr := reqQuery.Get("limit")
	keyword := reqQuery.Get("q")
	keyword = util.SanitizeSearchQuery(keyword)

	page, _ := strconv.ParseInt(pageAsStr, 10, 32)
	limit, _ := strconv.ParseInt(limitAsStr, 10, 32)

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	cnt, err := h.svc.TotalCount(
		domain.CountOptions{
			Keyword: keyword,
		}, true,
	)

	if err != nil {
		util.SendError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	postList, err := h.svc.Search(keyword, page, limit)

	if err != nil {
		util.SendError(w, "Failed to get posts", http.StatusInternalServerError)
		return
	}

	util.SendPage(w, postList, page, limit, cnt)
}