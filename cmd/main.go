package main

import (
	"fmt"
	"net/http"

	"github.com/n-korel/shortcut-api/configs"
	"github.com/n-korel/shortcut-api/internal/auth"
	"github.com/n-korel/shortcut-api/internal/link"
	"github.com/n-korel/shortcut-api/internal/stat"
	"github.com/n-korel/shortcut-api/internal/user"
	"github.com/n-korel/shortcut-api/pkg/db"
	"github.com/n-korel/shortcut-api/pkg/event"
	"github.com/n-korel/shortcut-api/pkg/middleware"
)


func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	//Repositories
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)

	//Services
	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus: eventBus,
		StatRepository: statRepository,
	})

	//Handlers
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

	//Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr: "127.0.0.1:8081",
		Handler: stack(router),
	}

	go statService.AddClick()

	fmt.Println("Server is listening on port 8081")

	server.ListenAndServe()
}

