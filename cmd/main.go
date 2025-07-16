package main

import (
	"fmt"
	"gourlshorter/v2/internal/auth"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	auth.NewAuthHandler(router)
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Starting server on :8081")
	server.ListenAndServe()
}
