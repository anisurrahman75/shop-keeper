package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/anisurrahman75/go-stock-management/api/data"
	"github.com/anisurrahman75/go-stock-management/api/models"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func GetCustomerInfo(w http.ResponseWriter, r *http.Request) {
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
