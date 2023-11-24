package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/anisurrahman75/go-stock-management/api/data"
	"github.com/anisurrahman75/go-stock-management/api/models"
	"html/template"
	"net/http"
)

func ProductAdd(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	temp, err := template.ParseFiles("./api/templates/views/productadd.html")
	if err != nil {
		panic(err)
	}
	if r.Method == http.MethodGet {
		response := struct {
			BrandList []models.Brand
			UnitList  []models.Unit
		}{data.BrandList, data.UnitList}
		err = temp.Execute(w, response)
		if err != nil {
			http.Error(w, "Error on executing template", http.StatusBadRequest)
		}

	} else if r.Method == http.MethodPost {
		var formData map[string]string
		err := json.NewDecoder(r.Body).Decode(&formData)
		if err != nil {
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}
		response := struct {
			IsPoductAddSuccessfully bool `json:"is_poduct_add_successfully"`
		}{false}
		yes := addIntoProductList(formData)
		if yes {
			response.IsPoductAddSuccessfully = true
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			panic(err)
		}
	}
}

func addIntoProductList(formData map[string]string) bool {
	product := models.Product{
		Brand:       models.Brand{Name: formData["brand"]},
		Name:        formData["product_name"],
		Grade:       formData["product_grade"],
		Unit:        models.Unit{Name: formData["unit"]},
		Description: formData["description"],
		Qty:         0,
	}
	data.ProductList = append(data.ProductList, product)
	return true
}

func GetCustomerList(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(data.CustomerList)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}
}

func ProductList(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	temp, err := template.ParseFiles("./api/templates/views/productlist.html")
	if err != nil {
		panic(err)
	}
	if r.Method == http.MethodGet {
		response := struct {
			ProductList []models.Product
		}{data.ProductList}

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
	}
}
