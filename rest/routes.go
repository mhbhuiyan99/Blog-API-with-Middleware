package rest

import (
	"blogAPI/rest/handlers"
	"blogAPI/rest/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"GET /posts", 
		manager.With(
			http.HandlerFunc(handlers.GetPosts),
	))

	mux.Handle(
		"POST /posts", 
		manager.With(
			http.HandlerFunc(handlers.CreatePost),
	))

	mux.Handle(
		"GET /posts/{postId}", 
		manager.With(
			http.HandlerFunc(handlers.GetPostByID),
	))
}