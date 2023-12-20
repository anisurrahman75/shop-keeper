package models

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type SaleProductInfo struct {
	Index     int
	Name      string `json:"name"`
	UnitPrice string `json:"unitPrice"`
	Qty       string `json:"qty"`
	SubTotal  string `json:"subTotal"`
}

type InvoiceData struct {
	CustomerInfo      Customer          `json:"CustomerInfo"`
	Date              string            `json:"Date"`
	InvoiceNo         string            `json:"InvoiceNo"`
	ProductsInfo      []SaleProductInfo `json:"ProductsInfo"`
	NetTotal          string            `json:"NetTotal"`
	DiscountInPercent string            `json:"DiscountInPercent"`
	SaveInDiscount    string            `json:"SaveInDiscount"`
	GrandTotal        string            `json:"GrandTotal"`
}

func (i InvoiceData) Add(ctx context.Context, client *firestore.Client) error {
	invoiceList, err := i.List(ctx, client, "")
	if err != nil {
		return err
	}
	invoiceList = append(invoiceList, i)

	mpData := make(map[string][]InvoiceData)
	mpData["List"] = invoiceList

	bytes, err := json.Marshal(&mpData)
	if err != nil {
		return err
	}

	var data map[string][]InvoiceData
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}
	monthYear := i.GetMonthYear()
	_, err = client.Collection("CustomerRecord").Doc(i.CustomerInfo.ShopNameOwner).Collection(monthYear).Doc("Invoice").Set(ctx, data)
	return err
}

func (i InvoiceData) List(ctx context.Context, client *firestore.Client, monthYear string) ([]InvoiceData, error) {
	if monthYear == "" {
		monthYear = i.GetMonthYear()
	}
	var list []InvoiceData

	fmt.Println(i.CustomerInfo.ShopNameOwner, monthYear)
	snap, _ := client.Collection("CustomerRecord").Doc(i.CustomerInfo.ShopNameOwner).Collection(monthYear).Doc("Invoice").Get(ctx)
	dbByte, err := json.Marshal(snap.Data())
	if err != nil {
		return nil, err
	}
	var invoiceList map[string][]InvoiceData
	if err := json.Unmarshal(dbByte, &invoiceList); err != nil {
		return nil, err
	}
	if val, ok := invoiceList["List"]; ok {
		for _, v := range val {
			list = append(list, v)
		}
	}
	return list, nil
}

func (i InvoiceData) GetMonthYear() string {
	date, _ := time.Parse("02-01-2006", i.Date)

	month := date.Format("January")
	year := date.Format("2006")

	monthYear := fmt.Sprintf("%s-%s", month, year)
	fmt.Println(monthYear)
	return monthYear
}
