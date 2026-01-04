package rest

import (
	"blogAPI/config"
	"blogAPI/rest/handlers/post"
	"blogAPI/rest/handlers/user"
	middleware "blogAPI/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	cnf 		*config.Config
	postHandler *post.Handler
	userHandler *user.Handler
}

func NewServer(
	cnf 		*config.Config,
	postHandler *post.Handler, 
	userHandler *user.Handler,
) *Server {
	return &Server{
		cnf: 		 cnf,
		postHandler: postHandler,
		userHandler: userHandler,
	}
}

func (server *Server) Start() {
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

	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server: ", err)
		os.Exit(1)
	}
}