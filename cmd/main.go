package main

import (
	"fmt"
	"gourlshorter/v2/configs"
	"gourlshorter/v2/internal/auth"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	} 
	fmt.Println("Starting server on :8081")
	server.ListenAndServe()
}
