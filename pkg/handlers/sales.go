package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/anisurrahman75/go-stock-management/api/models"
	"html/template"
	"net/http"
)

func (h *Handler) NewSales(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/views/newsales.html")
	if err != nil {
		panic(err)
	}
	if r.Method == http.MethodGet {
		response := struct {
			InvoiceNo   int
			ProductList []models.Product
		}{}
		productList, err := models.Product{}.List(context.Background(), h.Client)
		if err != nil {
			fmt.Println(err)
			return
		}
		response.ProductList = productList

		err = temp.Execute(w, response)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error on executing template", http.StatusBadRequest)
		}
	} else if r.Method == http.MethodPost {
		fmt.Println("----------Hello---------")
		var formData map[string]string
		err := json.NewDecoder(r.Body).Decode(&formData)
		if err != nil {
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}
		fmt.Println(formData)
	}
}
