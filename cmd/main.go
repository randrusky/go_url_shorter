package main

import (
	"fmt"
	"gourlshorter/v2/configs"
	"gourlshorter/v2/internal/auth"
	"gourlshorter/v2/pkg/db"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
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
