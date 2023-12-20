package routes

import (
	"cloud.google.com/go/firestore"
	"github.com/anisurrahman75/go-stock-management/config"
	"github.com/anisurrahman75/go-stock-management/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type Server struct {
	Router *chi.Mux
	Client *firestore.Client
}

func CreateNewServer() *Server {
	s := &Server{
		Router: chi.NewRouter(),
	}
	var err error
	s.Client, err = config.DBConnect()
	if err != nil {
		log.Print(err)
	}
	return s
}

func (s *Server) MountHandlers() http.Handler {
	handler := handlers.New(s.Client)
	s.Router.Group(func(r chi.Router) {
		r.HandleFunc("/signup", handler.SignUp)
		r.HandleFunc("/signin", handler.SignIn)
	})
	//s.Router.Group(func(r chi.Router) {
	s.Router.Group(func(r chi.Router) {
		//r.Use(auth.Verify)
		r.HandleFunc("/dashboard", handler.Dashboard)
		r.HandleFunc("/signout", handler.SignOut)

		r.Route("/get", func(r chi.Router) {
			r.HandleFunc("/customer/list", handler.GetCustomerList)
		})

		r.Route("/product", func(r chi.Router) {
			r.HandleFunc("/add", handler.ProductAdd)
			r.HandleFunc("/list", handler.ProductList)
			r.HandleFunc("/brand-unit", handler.BrandUnit)

		})

		r.Route("/customer", func(r chi.Router) {
			r.HandleFunc("/add", handler.AddCustomer)
			r.HandleFunc("/list", handler.ListCustomer)
		})

		r.HandleFunc("/customer/details/{shop}/{owner}", handler.DetailsCustomer)

		r.Route("/sales", func(r chi.Router) {
			r.HandleFunc("/new", handler.NewSales)
			r.HandleFunc("/invoice", handler.InvoicePrint)
			r.HandleFunc("/invoice-add", handler.InvoiceAddIntoRecord)
		})
	})
	return s.Router
}

func (s *Server) LoadAllStaticFiles() {
	frameworkFiles := http.FileServer(http.Dir("./templates/assets"))
	s.Router.Handle("/assets/*", http.StripPrefix("/assets/", frameworkFiles))

	myJsFiles := http.FileServer(http.Dir("./templates/views-js"))
	s.Router.Handle("/views-js/*", http.StripPrefix("/views-js/", myJsFiles))
}
