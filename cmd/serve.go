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
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	userHandler := usrHandler.NewHandler(cnf, usrSvc, middlewares)

	server := rest.NewServer(
		cnf,
		postHandler,
		userHandler,
	)


	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Error starting the server: ", err)
		}
	}()

	// block until Ctrl+C or Kill signal is received
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	
	fmt.Println("Shutting down server...")

	// give in-flight requests 5 seconds to complete before shutting down
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Error shutting down server: ", err)
	}

	if err := dbCon.Close(); err != nil {
		fmt.Println("Error closing database connection: ", err)
	}
	
	fmt.Println("Server gracefully stopped")
}
