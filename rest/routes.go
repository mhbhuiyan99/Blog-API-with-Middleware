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
			middleware.AuthenticateJWT,
	))

	mux.Handle(
		"GET /posts/{id}", 
		manager.With(
			http.HandlerFunc(handlers.GetPostByID),
	))
	/*
	mux.Handle(
		"PUT /posts/{id}", 
		manager.With(
			http.HandlerFunc(handlers.UpdatePost),
			middleware.AuthenticateJWT,
	))

	mux.Handle(
		"DELETE /posts/{id}", 
		manager.With(
			http.HandlerFunc(handlers.DeletePost),
			middleware.AuthenticateJWT,
	))
	*/			
}