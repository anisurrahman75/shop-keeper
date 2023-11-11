package routes

import (
	"github.com/anisurrahman75/go-stock-management/api/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Server struct {
	Router *chi.Mux
}

func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}

func (s *Server) MountHandlers() http.Handler {
	s.Router.Group(func(r chi.Router) {
		r.HandleFunc("/signup", handlers.SignUp)
		r.HandleFunc("/signin", handlers.SignIn)
	})

	//s.Router.Group(func(r chi.Router) {
	s.Router.Group(func(r chi.Router) {
		//r.Use(auth.Verify)
		r.HandleFunc("/dashboard", handlers.Dashboard)
		r.HandleFunc("/signout", handlers.SignOut)
		r.HandleFunc("/productadd", handlers.ProductAdd)
		r.HandleFunc("/productlist", handlers.ProductList)

	})
	return s.Router
}
