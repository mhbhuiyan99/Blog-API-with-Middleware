package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fmt.Println("Server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}