package user

import (
	middleware "blogAPI/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle(
		"POST /users",
		manager.With(
			http.HandlerFunc(h.CreateUser),
			h.middlewares.RateLimit,
		))

	mux.Handle(
		"POST /users/login",
		manager.With(
			http.HandlerFunc(h.Login),
			h.middlewares.RateLimit,
		))
}
