package cmd

import (
	"blogAPI/middleware"
	"fmt"
	"net/http"
)

func Serve() {

	manager := middleware.NewManager()

	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	fmt.Println("Server is running on port 4000")
	err := http.ListenAndServe(":4000", wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server: ", err)
	}
}