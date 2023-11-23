package main

import (
	"github.com/anisurrahman75/go-stock-management/pkg/routes"
	mi "github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	server := routes.CreateNewServer()
	server.Router.Use(mi.Logger)
	server.LoadAllStaticFiles()
	err := http.ListenAndServe(":8000", server.MountHandlers())
	if err != nil {
		panic(err)
	}
}
