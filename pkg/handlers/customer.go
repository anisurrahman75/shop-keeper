package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/anisurrahman75/go-stock-management/api/data"
	"github.com/anisurrahman75/go-stock-management/api/models"
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
	"net/http"
)

func (h *Handler) AddCustomer(w http.ResponseWriter, r *http.Request) {
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
			AddSuccess bool `json:"AddSuccess"`
		}{true}

		if err := cusomerData.Add(context.Background(), h.Client); err != nil {
			log.Printf("Error adding Customer: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
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

func (h *Handler) ListCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	temp, err := template.ParseFiles("./templates/views/customerlist.html")
	if err != nil {
		panic(err)
	}
	if r.Method == http.MethodGet {
		customerList, err := models.Customer{}.List(context.Background(), h.Client)
		if err != nil {
			log.Printf("Error getting customer list: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		fmt.Println(customerList)
		err = temp.Execute(w, customerList)
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

type CustomerData struct {
	Shop string `json:"Shop"`
	Name string `json:"Name"`
}

func (h *Handler) DetailsCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	shop := chi.URLParam(r, "shop")
	owner := chi.URLParam(r, "owner")

	fmt.Println("---" + shop + owner + "---")

	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("./templates/views/customerdetails.html")
		if err != nil {
			panic(err)
		}
		if r.Method == http.MethodGet {
			if err != nil {
				panic(err)
			}
			response := struct {
				Info   models.Customer
				Record models.CustomerRecord
			}{}
			err := temp.Execute(w, response)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
		}
	} else {
		var monthYear = struct {
			Month string `json:"Month"`
			Year  string `json:"Year"`
		}{}
		err := json.NewDecoder(r.Body).Decode(&monthYear)
		if err != nil {
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}

		monthYearStr := monthYear.Month + "-" + monthYear.Year
		invoiceData := models.InvoiceData{
			CustomerInfo: models.Customer{
				Shop:  shop,
				Owner: owner,
			},
		}
		invoiceData.CustomerInfo.ShopNameOwner = invoiceData.CustomerInfo.GetShopNameOwner()
		invoiceList, err := invoiceData.List(context.Background(), h.Client, monthYearStr)
		if err != nil {
			log.Printf("Error adding Customer: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		response := struct {
			GetSuccess  bool                 `json:"GetSuccess"`
			InvoiceList []models.InvoiceData `json:"InvoiceList"`
		}{true, invoiceList}

		fmt.Println(response)

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

func (h *Handler) GetCustomerList(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == http.MethodGet {
		customerList, err := models.Customer{}.List(context.Background(), h.Client)
		if err != nil {
			fmt.Println(err)
			return
		}
		for i := 0; i < len(customerList); i++ {
			customerList[i].ShopNameOwner = customerList[i].GetShopNameOwner()
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(customerList)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}
}
