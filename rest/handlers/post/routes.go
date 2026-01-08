package post

import (
	middleware "blogAPI/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"GET /posts", 
		manager.With(
			http.HandlerFunc(h.GetPosts),
	))

	mux.Handle(
		"POST /posts", 
		manager.With(
			http.HandlerFunc(h.CreatePost),
			h.middlewares.AuthenticateJWT,
	))

	mux.Handle(
		"GET /posts/{id}", 
		manager.With(
			http.HandlerFunc(h.GetPost),
	))
	/*
	mux.Handle(
		"PUT /posts/{id}", 
		manager.With(
			http.HandlerFunc(handlers.UpdatePost),
			h.middlewares.AuthenticateJWT,
	))

	mux.Handle(
		"DELETE /posts/{id}", 
		manager.With(
			http.HandlerFunc(handlers.DeletePost),
			h.middlewares.AuthenticateJWT,
	))
	*/			
}