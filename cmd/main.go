package main

import (
	"fmt"
	"gourlshorter/v2/configs"
	"gourlshorter/v2/internal/auth"
	"gourlshorter/v2/internal/link"
	"gourlshorter/v2/internal/stat"
	"gourlshorter/v2/internal/user"
	"gourlshorter/v2/pkg/db"
	"gourlshorter/v2/pkg/event"
	"gourlshorter/v2/pkg/middleware"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)

	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})
	
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		Config: conf,
		EventBus: eventBus,		
	})

	stat.NewStatHandler(router, stat.StatHandlerDeps{
		StatRepository: statRepository,
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
	go statService.AddClick() 
	
	// Start the stat service to listen for events

	fmt.Println("Starting server on :8081")
	server.ListenAndServe()
}
