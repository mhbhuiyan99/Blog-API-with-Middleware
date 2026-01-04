package cmd

import (
	"blogAPI/config"
	"blogAPI/rest"
	"blogAPI/rest/handlers/post"
	"blogAPI/rest/handlers/user"
	middleware "blogAPI/rest/middlewares"
)

func Serve() {

	cnf :=config.GetConfig() // Load configuration

	middlewares := middleware.NewMiddlewares(cnf)

	postHandler := post.NewHandler(middlewares)
	userHandler := user.NewHandler(middlewares)

	server := rest.NewServer(
		cnf, 
		postHandler, 
		userHandler,
	)
	server.Start() 
	
}