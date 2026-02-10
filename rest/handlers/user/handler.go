package user

import (
	"blogAPI/config"
	middleware "blogAPI/rest/middlewares"
)

type Handler struct {
	cnf         *config.Config
	svc         Service
	middlewares *middleware.Middlewares
}

func NewHandler(
	cnf *config.Config,
	svc Service,
	middlewares *middleware.Middlewares,
) *Handler {
	return &Handler{
		cnf:         cnf,
		svc:         svc,
		middlewares: middlewares,
	}
}
