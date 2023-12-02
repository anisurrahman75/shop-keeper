package models

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/fatih/structs"
	"log"
	"strconv"
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
	InvoiceNo         int               `json:"InvoiceNo"`
	ProductsInfo      []SaleProductInfo `json:"ProductsInfo"`
	NetTotal          float64           `json:"NetTotal"`
	DiscountInPercent int               `json:"DiscountInPercent"`
	SaveInDiscount    float64           `json:"SaveInDiscount"`
	GrandTotal        float64           `json:"GrandTotal"`
}

func (i InvoiceData) Add(ctx context.Context, client *firestore.Client) error {
	dataMap := structs.Map(i)
	log.Printf("Adding Invoice data: %+v", dataMap)
	_, err := client.Collection("InvoiceList").Doc(strconv.Itoa(i.InvoiceNo)).Set(ctx, dataMap)
	return err
}

func (i InvoiceData) Update(ctx context.Context, client *firestore.Client) error {
	dataMap := structs.Map(i)
	log.Printf("Updating Invoice data: %+v", dataMap)
	_, err := client.Collection("InvoiceList").Doc(strconv.Itoa(i.InvoiceNo)).Set(ctx, dataMap, firestore.MergeAll)
	return err
}

func (i InvoiceData) Delete(ctx context.Context, client *firestore.Client) error {
	dataMap := structs.Map(i)
	log.Printf("Deleting Invoice : %+v", dataMap)
	_, err := client.Collection("InvoiceList").Doc(strconv.Itoa(i.InvoiceNo)).Delete(ctx)
	return err
}
