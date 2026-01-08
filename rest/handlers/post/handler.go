package post

import (
	middleware "blogAPI/rest/middlewares"
	"blogAPI/repo"
)

type Handler struct {
	middlewares *middleware.Middlewares
	postRepo repo.PostRepo
}

func NewHandler(
	middlewares *middleware.Middlewares,
	postRepo repo.PostRepo,
	) *Handler {
	return &Handler{
		middlewares: middlewares,
		postRepo: postRepo,
	}
}