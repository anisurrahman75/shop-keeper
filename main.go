package main

import (
	"github.com/anisurrahman75/go-stock-management/pkg/routes"
	mi "github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	server := routes.CreateNewServer()
	server.Router.Use(mi.Logger)
	fs := http.FileServer(http.Dir("./templates/"))
	server.Router.Handle("/assets/*", http.StripPrefix("/assets/", fs))

	//http.Handle("/", http.FileServer(http.Dir("./templates")))

	// Serve CSS, JS, and other assets from the "assets" directory
	//http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./templates/assets"))))

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./templates/assets"))))

	err := http.ListenAndServe(":8000", server.MountHandlers())
	if err != nil {
		panic(err)
	}
}
