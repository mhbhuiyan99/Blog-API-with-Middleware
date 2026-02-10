package post

import (
	middleware "blogAPI/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	
	mux.Handle(
		"POST /posts",
		manager.With(
			http.HandlerFunc(h.CreatePost),
			h.middlewares.AuthenticateJWT,
			h.middlewares.RateLimit,
		))
	
	mux.Handle(
		"GET /posts",
		manager.With(
			http.HandlerFunc(h.GetPosts),
			h.middlewares.RateLimit,
		))

	mux.Handle(
		"GET /posts/{id}",
		manager.With(
			http.HandlerFunc(h.GetPost),
			h.middlewares.RateLimit,
		))

	mux.Handle(
		"GET /my-posts/drafts",
		manager.With(
			http.HandlerFunc(h.GetDrafts),
			h.middlewares.AuthenticateJWT,
			h.middlewares.RateLimit,
		))
	
	mux.Handle(
		"GET /my-posts/published",
		manager.With(
			http.HandlerFunc(h.GetPublished),
			h.middlewares.AuthenticateJWT,
			h.middlewares.RateLimit,
		))
	
	mux.Handle(
		"GET /users/{userId}/posts",
		manager.With(
			http.HandlerFunc(h.GetUserPosts),
			h.middlewares.RateLimit,
		))

	mux.Handle(
		"PUT /posts/{id}",
		manager.With(
			http.HandlerFunc(h.UpdatePost),
			h.middlewares.AuthenticateJWT,
			h.middlewares.RateLimit,
		))

	mux.Handle(
		"DELETE /posts/{id}",
		manager.With(
			http.HandlerFunc(h.DeletePost),
			h.middlewares.AuthenticateJWT,
			h.middlewares.RateLimit,
		))

}
