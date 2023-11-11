package main

import (
	"github.com/anisurrahman75/go-stock-management/api/routes"
	mi "github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	server := routes.CreateNewServer()
	server.Router.Use(mi.Logger)
	fs := http.FileServer(http.Dir("./api/templates/assets"))
	server.Router.Handle("/assets/*", http.StripPrefix("/assets/", fs))

	err := http.ListenAndServe(":8000", server.MountHandlers())
	if err != nil {
		panic(err)
	}
}
