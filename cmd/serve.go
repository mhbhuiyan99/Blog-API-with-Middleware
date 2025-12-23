package cmd

import (
	"blogAPI/global_router"
	"blogAPI/handlers"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /posts", http.HandlerFunc(handlers.GetPosts))
	mux.Handle("POST /posts", http.HandlerFunc(handlers.CreatePost))
	mux.Handle("GET /posts/{postId}", http.HandlerFunc(handlers.GetPostByID))

	fmt.Println("Server is running on port 4000")
	err := http.ListenAndServe(":4000", global_router.GlobalRouter(mux))
	if err != nil {
		fmt.Println("Error starting the server: ", err)
	}
}