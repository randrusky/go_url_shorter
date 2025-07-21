package main

import (
	"fmt"
	"gourlshorter/v2/configs"
	"gourlshorter/v2/internal/auth"
	"gourlshorter/v2/internal/link"
	"gourlshorter/v2/pkg/db"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	linkRepository := link.NewLinkRepository(db)
	
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	} 
	fmt.Println("Starting server on :8081")
	server.ListenAndServe()
}
