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

	temp, err := template.ParseFiles("./templates/views/dashboard.html")
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

	fmt.Println("-----------------Debugging----------------")
	fmt.Println(invoiceData.CustomerInfo)
	fmt.Println(invoiceData.NetTotal)
	fmt.Println(invoiceData.DiscountInPercent)
	fmt.Println(invoiceData.SaveInDiscount)
	fmt.Println(invoiceData.GrandTotal)
	fmt.Println("-----------------End Debugging-------------")

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

func getCustomerInfo(shopName string, ownerName string) (*models.Customer, error) {
	for _, info := range data.CustomerList {
		if info.ShopName == shopName && info.OwnerName == ownerName {
			return &info, nil
		}
	}
	return nil, fmt.Errorf("no Customer Data Found")
}
