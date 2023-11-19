package routes

import (
	"github.com/anisurrahman75/go-stock-management/pkg/handlers"
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

		r.Route("/product", func(r chi.Router) {
			r.HandleFunc("/add", handlers.ProductAdd)
			r.HandleFunc("/list", handlers.ProductList)

		})

		r.Route("/customer", func(r chi.Router) {
			r.HandleFunc("/add", handlers.AddCustomer)
			r.HandleFunc("/{shopName}", handlers.GetCustomer)
		})

		r.Route("/sales", func(r chi.Router) {
			r.HandleFunc("/new", handlers.NewSales)
			r.HandleFunc("/invoice", handlers.InvoicePrint)
		})

	})
	return s.Router
}
