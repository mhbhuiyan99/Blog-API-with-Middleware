package rest

import (
	"blogAPI/config"
	"blogAPI/rest/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Start(cnf config.Config) {
	manager := middleware.NewManager()

	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	addr := ":" + strconv.Itoa(cnf.HttpPort)
	fmt.Println("Server is running on port ", addr)

	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server: ", err)
		os.Exit(1)
	}
}