package cmd

import (
	"blogAPI/config"
	"blogAPI/infra/db"
	"blogAPI/post"
	"blogAPI/repo"
	"blogAPI/rest"
	pstHandler "blogAPI/rest/handlers/post"
	usrHandler "blogAPI/rest/handlers/user"
	middleware "blogAPI/rest/middlewares"
	"blogAPI/user"
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

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	middlewares := middleware.NewMiddlewares(cnf)

	// repos
	postRepo := repo.NewPostRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	// domains
	usrSvc := user.NewService(userRepo)
	pstSvc := post.NewService(postRepo)

	// handlers
	postHandler := pstHandler.NewHandler(middlewares, pstSvc)
	userHandler := usrHandler.NewHandler(cnf, usrSvc)

	server := rest.NewServer(
		cnf,
		postHandler,
		userHandler,
	)
	server.Start()

}
