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
		))
	
	mux.Handle(
		"GET /posts",
		manager.With(
			http.HandlerFunc(h.GetPosts),
		))

	mux.Handle(
		"GET /posts/{id}",
		manager.With(
			http.HandlerFunc(h.GetPost),
		))

	mux.Handle(
		"GET /my-posts/drafts",
		manager.With(
			http.HandlerFunc(h.GetDrafts),
			h.middlewares.AuthenticateJWT,
		))
	
	mux.Handle(
		"GET /my-posts/published",
		manager.With(
			http.HandlerFunc(h.GetPublished),
			h.middlewares.AuthenticateJWT,
		))
	
	mux.Handle(
		"GET /users/{userId}/posts",
		manager.With(
			http.HandlerFunc(h.GetUserPosts),
		))

	mux.Handle(
		"PUT /posts/{id}",
		manager.With(
			http.HandlerFunc(h.UpdatePost),
			h.middlewares.AuthenticateJWT,
		))

	mux.Handle(
		"DELETE /posts/{id}",
		manager.With(
			http.HandlerFunc(h.DeletePost),
			h.middlewares.AuthenticateJWT,
		))

}
