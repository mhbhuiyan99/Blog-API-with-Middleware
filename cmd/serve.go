package cmd

import (
	"blogAPI/config"
	"blogAPI/repo"
	"blogAPI/rest"
	"blogAPI/rest/handlers/post"
	"blogAPI/rest/handlers/user"
	middleware "blogAPI/rest/middlewares"
)

func Serve() {

	cnf :=config.GetConfig() // Load configuration

	middlewares := middleware.NewMiddlewares(cnf)

	postRepo := repo.NewPostRepo()
	userRepo := repo.NewUserRepo()

	postHandler := post.NewHandler(middlewares, postRepo)
	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(
		cnf, 
		postHandler, 
		userHandler,
	)
	server.Start() 
	
}