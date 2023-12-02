package models

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/fatih/structs"
	"log"
	"strconv"
)

type Customer struct {
	ID          int    `json:"ID"`
	ShopName    string `json:"ShopName"`
	OwnerName   string `json:"OwnerName"`
	PhoneNumber string `json:"PhoneNumber"`
	Address     string `json:"Address"`
	TotalDue    int    `json:"TotalDue"`
}

type CustomerRecord struct {
	CustomerID int           `json:"CustomerID"`
	Purchased  []InvoiceData `json:"Purchased"`
	Payment    []Payment     `json:"Payment"`
}
type Payment struct {
}

func (c Customer) Add(ctx context.Context, client *firestore.Client) error {
	dataMap := structs.Map(c)
	log.Printf("Adding Customer data: %+v", dataMap)
	_, err := client.Collection("CustomerList").Doc(strconv.Itoa(c.ID)).Set(ctx, dataMap)
	return err
}

func (c Customer) Update(ctx context.Context, client *firestore.Client) error {
	dataMap := structs.Map(c)
	log.Printf("Updating Customer data: %+v", dataMap)
	_, err := client.Collection("CustomerList").Doc(strconv.Itoa(c.ID)).Set(ctx, dataMap, firestore.MergeAll)
	return err
}

func (c Customer) Delete(ctx context.Context, client *firestore.Client) error {
	dataMap := structs.Map(c)
	log.Printf("Deleting Customer : %+v", dataMap)
	_, err := client.Collection("CustomerList").Doc(strconv.Itoa(c.ID)).Delete(ctx)
	return err
}
