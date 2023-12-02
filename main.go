package main

import (
	"github.com/anisurrahman75/go-stock-management/pkg/db"
	"github.com/anisurrahman75/go-stock-management/pkg/routes"
	mi "github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	server := routes.CreateNewServer()
	server.Router.Use(mi.Logger)
	if err := db.InitDatabase(server); err != nil {
		log.Printf("An error has occurred: %s", err)
		panic(err)
	}
	server.LoadAllStaticFiles()
	if err := http.ListenAndServe(":8080", server.MountHandlers()); err != nil {
		log.Printf("An error has occurred: %s", err)
		panic(err)
	}
}
