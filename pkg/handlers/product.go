package handlers

import (
	"context"
	"encoding/json"
	"fmt"
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
		var product models.Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			log.Printf("Error parsing JSON: %v", err)
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}
		fmt.Println(product)

		response := struct {
			AddSuccess bool `json:"AddSuccess"`
		}{true}
		if err := product.Add(context.TODO(), h.Client); err != nil {
			log.Printf("Error Adding Product: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
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

func (h *Handler) ProductList(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/views/productlist.html")
	if err != nil {
		panic(err)
	}
	if r.Method == http.MethodGet {
		productList, err := models.Product{}.List(context.Background(), h.Client)
		if err != nil {
			log.Printf("Error getting Product list: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		response := struct {
			ProductList []models.Product
		}{productList}

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
	} else if r.Method == http.MethodDelete {
		var formData = struct {
			Type string `json:"type"`
			Name string `json:"name"`
		}{}

		err := json.NewDecoder(r.Body).Decode(&formData)
		if err != nil {
			log.Printf("Error parsing JSON: %v", err)
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}
		fmt.Println("hello", formData.Name, "hello")

		product := models.Product{
			Name: formData.Name,
		}
		if err := product.Delete(context.Background(), h.Client); err != nil {
			log.Printf("Error parsing JSON: %v", err)
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}

		delResponse := struct {
			DeleteSuccessful bool `json:"DeleteSuccessful"`
		}{true}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err = json.NewEncoder(w).Encode(delResponse)
		if err != nil {
			log.Printf("Error parsing JSON: %v", err)
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}
	}
}

func (h *Handler) BrandUnit(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	temp, err := template.ParseFiles("./templates/views/brand,unit-manage.html")
	if err != nil {
		panic(err)
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

		fmt.Println(response)

		err = temp.Execute(w, response)
		if err != nil {
			log.Printf("Error executing template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

	} else if r.Method == http.MethodPost {
		postRes := struct {
			AddSuccessful bool `json:"AddSuccessful"`
		}{true}

		var formData = struct {
			Type string `json:"type"`
			Name string `json:"name"`
		}{}

		err := json.NewDecoder(r.Body).Decode(&formData)
		if err != nil {
			log.Printf("Error parsing JSON: %v", err)
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}
		switch formData.Type {
		case "brand":
			brand := models.Brand{
				Name: formData.Name,
			}
			if err := brand.Add(context.Background(), h.Client); err != nil {
				log.Printf("Error parsing JSON: %v", err)
				http.Error(w, "Error parsing JSON", http.StatusBadRequest)
				return
			}

		case "unit":
			unit := models.Unit{
				Name: formData.Name,
			}
			if err := unit.Add(context.Background(), h.Client); err != nil {
				log.Printf("Error parsing JSON: %v", err)
				http.Error(w, "Error parsing JSON", http.StatusBadRequest)
				return
			}
		default:
			http.Error(w, "Invalid form type field", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(postRes)
		if err != nil {
			log.Printf("Error parsing JSON: %v", err)
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}
	} else if r.Method == http.MethodDelete {
		var formData = struct {
			Type string `json:"type"`
			Name string `json:"name"`
		}{}

		err := json.NewDecoder(r.Body).Decode(&formData)
		if err != nil {
			log.Printf("Error parsing JSON: %v", err)
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}
		fmt.Println(formData)
		switch formData.Type {
		case "brand":
			brand := models.Brand{
				Name: formData.Name,
			}
			if err := brand.Delete(context.Background(), h.Client); err != nil {
				log.Printf("Error parsing JSON: %v", err)
				http.Error(w, "Error parsing JSON", http.StatusBadRequest)
				return
			}

		case "unit":
			unit := models.Unit{
				Name: formData.Name,
			}
			if err := unit.Delete(context.Background(), h.Client); err != nil {
				log.Printf("Error parsing JSON: %v", err)
				http.Error(w, "Error parsing JSON", http.StatusBadRequest)
				return
			}
		default:
			http.Error(w, "Invalid form type field", http.StatusBadRequest)
			return
		}

		delResponse := struct {
			DeleteSuccessful bool `json:"DeleteSuccessful"`
		}{true}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err = json.NewEncoder(w).Encode(delResponse)
		if err != nil {
			log.Printf("Error parsing JSON: %v", err)
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}

	}
}
