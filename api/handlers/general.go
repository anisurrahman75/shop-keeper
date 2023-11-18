package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/anisurrahman75/go-stock-management/api/data"
	"github.com/anisurrahman75/go-stock-management/api/models"
	"html/template"
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method

	temp, err := template.ParseFiles("./api/templates/views/dashboard.html")
	if err != nil {
		panic(err)
	}
	err = temp.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func InvoicePrint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method)
	tem, err := template.ParseFiles("./api/templates/views/invoice.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	type products struct {
		Name      string `json:"name"`
		UnitPrice string `json:"unitPrice"`
		Qty       string `json:"qty"`
		SubTotal  int    `json:"subTotal"`
	}

	type InvoiceData struct {
		ShopName     string     `json:"ShopName"`
		OwnerName    string     `json:"OwnerName"`
		ProductsInfo []products `json:"ProductsInfo"`
	}
	// Parse JSON
	var invoiceData InvoiceData
	err = json.NewDecoder(r.Body).Decode(&invoiceData)
	if err != nil {
		fmt.Println(err)
		http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
		return
	}
	fmt.Println(invoiceData.ShopName, " ", invoiceData.OwnerName)
	customerInfo, err := getCustomerInfo(invoiceData.ShopName, invoiceData.OwnerName)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error on executing template", http.StatusBadRequest)
	}

	response := struct {
		CustomerInfo *models.Customer
		ProductsList []products
		NetTotal     int64
	}{CustomerInfo: customerInfo, ProductsList: invoiceData.ProductsInfo, NetTotal: 100}

	//Execute the template without passing any data.
	err = tem.Execute(w, response)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error on executing template", http.StatusBadRequest)
	}

}

func getCustomerInfo(shopName string, ownerName string) (*models.Customer, error) {
	for _, info := range data.CustomerList {
		if info.ShopName == shopName && info.OwnerName == ownerName {
			return &info, nil
		}
	}
	return nil, fmt.Errorf("no Customer Data Found")
}
