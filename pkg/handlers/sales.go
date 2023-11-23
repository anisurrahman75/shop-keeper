package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/anisurrahman75/go-stock-management/api/data"
	"github.com/anisurrahman75/go-stock-management/api/models"
	"html/template"
	"net/http"
)

func NewSales(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	temp, err := template.ParseFiles("./templates/views/newsales.html")
	if err != nil {
		panic(err)
	}
	if r.Method == http.MethodGet {
		response := struct {
			InvoiceNo    string
			CustomerList []models.Customer
			ProductList  []models.Product
		}{"12", data.CustomerList, data.ProductList}

		err = temp.Execute(w, response)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error on executing template", http.StatusBadRequest)
		}
	} else if r.Method == http.MethodPost {
		var formData map[string]string
		err := json.NewDecoder(r.Body).Decode(&formData)
		if err != nil {
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}
		fmt.Println(formData)
	}
}
