package rest

import (
	"blogAPI/config"
	"blogAPI/rest/handlers/post"
	"blogAPI/rest/handlers/user"
	middleware "blogAPI/rest/middlewares"
	"context"
	"fmt"
	"net/http"
	"strconv"
)

type Server struct {
	cnf         *config.Config
	postHandler *post.Handler
	userHandler *user.Handler
	httpServer  *http.Server
}

func NewServer(
	cnf *config.Config,
	postHandler *post.Handler,
	userHandler *user.Handler,
) *Server {
	return &Server{
		cnf:         cnf,
		postHandler: postHandler,
		userHandler: userHandler,
	}
}

func (server *Server) Start() error {
	manager := middleware.NewManager()

	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	//initRoutes
	server.postHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(server.cnf.HttpPort)
	fmt.Println("Server is running on port ", addr)

	server.httpServer = &http.Server{
		Addr:    addr,
		Handler: wrappedMux,
	}

	err := server.httpServer.ListenAndServe()
	return err
}

func (server *Server) Shutdown(ctx context.Context) error {
	return server.httpServer.Shutdown(ctx)
}
