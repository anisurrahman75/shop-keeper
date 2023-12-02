package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/anisurrahman75/go-stock-management/api/data"
	"github.com/anisurrahman75/go-stock-management/api/models"
	"html/template"
	"log"
	"net/http"
)

func (h *Handler) ProductAdd(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	temp, err := template.ParseFiles("./templates/views/productadd.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodGet {
		brandList, err := models.Brand{}.List(context.Background(), h.Client)
		if err != nil {
			log.Printf("Error getting brand list: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		unitList, err := models.Unit{}.List(context.Background(), h.Client)
		if err != nil {
			log.Printf("Error getting unit list: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		response := struct {
			BrandList []models.Brand
			UnitList  []models.Unit
		}{brandList, unitList}

		err = temp.Execute(w, response)
		if err != nil {
			log.Printf("Error executing template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	} else if r.Method == http.MethodPost {
		var formData map[string]string
		err := json.NewDecoder(r.Body).Decode(&formData)
		if err != nil {
			log.Printf("Error parsing JSON: %v", err)
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}

		response := struct {
			IsProductAddSuccessfully bool `json:"is_product_add_successfully"`
		}{false}

		if addIntoProductList(formData) {
			response.IsProductAddSuccessfully = true
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Printf("Error encoding JSON response: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
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

func (h *Handler) ProductList(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	temp, err := template.ParseFiles("./templates/views/productlist.html")
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
