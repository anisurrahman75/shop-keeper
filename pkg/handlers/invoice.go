package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/anisurrahman75/go-stock-management/api/models"
	"log"
	"net/http"
	"text/template"
)

func (h *Handler) InvoicePrint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method)
	tem, err := template.ParseFiles("./templates/views/invoice.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	// Parse JSON
	var invoiceData models.InvoiceData
	err = json.NewDecoder(r.Body).Decode(&invoiceData)
	if err != nil {
		fmt.Println(err)
		http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
		return
	}
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error on executing template", http.StatusBadRequest)
	}
	for i := range invoiceData.ProductsInfo {
		invoiceData.ProductsInfo[i].Index = i + 1
	}
	//Execute the template without passing any data.
	err = tem.Execute(w, invoiceData)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error on executing template", http.StatusBadRequest)
	}

}

func (h *Handler) InvoiceAddIntoRecord(w http.ResponseWriter, r *http.Request) {
	var invoiceData models.InvoiceData
	err := json.NewDecoder(r.Body).Decode(&invoiceData)
	if err != nil {
		fmt.Println(err)
		http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
		return
	}
	response := struct {
		AddSuccess bool `json:"AddSuccess"`
	}{true}

	if err := invoiceData.Add(context.Background(), h.Client); err != nil {
		log.Printf("Error adding Invoice: %v", err)
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
