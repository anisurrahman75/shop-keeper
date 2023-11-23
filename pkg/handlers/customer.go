package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/anisurrahman75/go-stock-management/api/data"
	"github.com/anisurrahman75/go-stock-management/api/models"
	"github.com/go-chi/chi/v5"
	"html/template"
	"net/http"
)

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	shopName := chi.URLParam(r, "shopName")
	fmt.Println("Shop Name : ", shopName)
	customerInfo := getCustomerInfoFromShopName(shopName)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(customerInfo); err != nil {
		panic(err)
	}
}

func getCustomerInfoFromShopName(name string) models.Customer {
	for _, customer := range data.CustomerList {
		if name == customer.ShopName {
			return customer
		}
	}
	return models.Customer{}
}

func AddCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	temp, err := template.ParseFiles("./templates/views/customeradd.html")
	if err != nil {
		panic(err)
	}
	if r.Method == http.MethodGet {
		err = temp.Execute(w, nil)
		if err != nil {
			http.Error(w, "Error on executing template", http.StatusBadRequest)
		}

	} else if r.Method == http.MethodPost {

		var cusomerData models.Customer
		err := json.NewDecoder(r.Body).Decode(&cusomerData)
		if err != nil {
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}
		response := struct {
			IsCustomerAddSuccessful bool `json:"isCustomerAddSuccessful"`
		}{false}
		yes := addIntoCustomerList(cusomerData)
		if yes {
			response.IsCustomerAddSuccessful = true
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

func ListCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	temp, err := template.ParseFiles("./templates/views/customeradd.html")
	if err != nil {
		panic(err)
	}
	if r.Method == http.MethodGet {
		err = temp.Execute(w, nil)
		if err != nil {
			http.Error(w, "Error on executing template", http.StatusBadRequest)
		}

	} else if r.Method == http.MethodPost {

		var cusomerData models.Customer
		err := json.NewDecoder(r.Body).Decode(&cusomerData)
		if err != nil {
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}
		response := struct {
			IsCustomerAddSuccessful bool `json:"isCustomerAddSuccessful"`
		}{false}
		yes := addIntoCustomerList(cusomerData)
		if yes {
			response.IsCustomerAddSuccessful = true
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

func addIntoCustomerList(customer models.Customer) bool {
	data.CustomerList = append(data.CustomerList, customer)
	return true
}
