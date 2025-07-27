package main

import (
	"fmt"
	"gourlshorter/v2/configs"
	"gourlshorter/v2/internal/auth"
	"gourlshorter/v2/internal/link"
	"gourlshorter/v2/internal/user"
	"gourlshorter/v2/pkg/db"
	"gourlshorter/v2/pkg/middleware"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	authService := auth.NewAuthService(userRepository)
	
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		Config: conf,		
	})

	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	} 
	fmt.Println("Starting server on :8081")
	server.ListenAndServe()
}
