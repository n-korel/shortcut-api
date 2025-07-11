package main

import (
	"fmt"
	"net/http"

	"github.com/n-korel/shortcut-api/configs"
	"github.com/n-korel/shortcut-api/internal/auth"
	"github.com/n-korel/shortcut-api/internal/link"
	"github.com/n-korel/shortcut-api/pkg/db"
)



func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	//Repositories
	linkRepository := link.NewLinkRepository(db)

	//Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	server := http.Server{
		Addr: "127.0.0.1:8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")

	server.ListenAndServe()
}

