package post

import (
	"blogAPI/repo"
	middleware "blogAPI/rest/middlewares"
)

type Handler struct {
	middlewares *middleware.Middlewares
	postRepo    repo.PostRepo
}

func NewHandler(
	middlewares *middleware.Middlewares,
	postRepo repo.PostRepo,
) *Handler {
	return &Handler{
		middlewares: middlewares,
		postRepo:    postRepo,
	}
}
