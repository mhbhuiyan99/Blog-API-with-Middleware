package cmd

import (
	"blogAPI/config"
	"blogAPI/infra/db"
	"blogAPI/repo"
	"blogAPI/rest"
	"blogAPI/rest/handlers/post"
	"blogAPI/rest/handlers/user"
	middleware "blogAPI/rest/middlewares"
	"fmt"
	"os"
)

func Serve() {

	cnf := config.GetConfig() // Load configuration

	dbCon, err := db.NewDBConnection(cnf.DB) // Establish database connection
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	middlewares := middleware.NewMiddlewares(cnf)

	postRepo := repo.NewPostRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	postHandler := post.NewHandler(middlewares, postRepo)
	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(
		cnf,
		postHandler,
		userHandler,
	)
	server.Start()

}
