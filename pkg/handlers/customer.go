package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/anisurrahman75/go-stock-management/api/data"
	"github.com/anisurrahman75/go-stock-management/api/models"
	"github.com/go-chi/chi/v5"
	"html/template"
	"net/http"
	"strconv"
)

func (h *Handler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	idStr := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Url Customer Id: ", idInt)
	customerInfo := getCustomerInfoFromId(idInt)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(customerInfo); err != nil {
		panic(err)
	}
}

func getCustomerInfoFromId(id int) models.Customer {
	for _, customer := range data.CustomerList {
		if id == customer.ID {
			return customer
		}
	}
	return models.Customer{}
}

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

func (h *Handler) ListCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	temp, err := template.ParseFiles("./templates/views/customerlist.html")
	if err != nil {
		panic(err)
	}
	if r.Method == http.MethodGet {
		err = temp.Execute(w, data.CustomerList)
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

func (h *Handler) DetailsCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)

	idStr := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	customerInfo := getCustomerInfoFromId(idInt)
	customerRecord := getCustomerRecordFromId(idInt)
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
		}{customerInfo, customerRecord}

		err := temp.Execute(w, response)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}
}

func getCustomerRecordFromId(idInt int) models.CustomerRecord {
	for _, record := range data.CustomerRecordList {
		if record.CustomerID == idInt {
			return record
		}
	}
	return models.CustomerRecord{}
}

func addIntoCustomerList(customer models.Customer) bool {
	data.CustomerList = append(data.CustomerList, customer)
	return true
}

func (h *Handler) GetCustomerList(w http.ResponseWriter, r *http.Request) {
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
