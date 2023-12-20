package models

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"github.com/fatih/structs"
	"log"
)

type Customer struct {
	ShopNameOwner string `json:"ShopNameOwner"`
	Shop          string `json:"Shop"`
	Owner         string `json:"Owner"`
	PhoneNumber   string `json:"PhoneNumber"`
	Address       string `json:"Address"`
	TotalDue      int    `json:"TotalDue"`
}

type CustomerRecord struct {
	CustomerID int           `json:"CustomerID"`
	Purchased  []InvoiceData `json:"Purchased"`
	Payment    []Payment     `json:"Payment"`
}
type Payment struct {
}

func (c Customer) Add(ctx context.Context, client *firestore.Client) error {
	customerList, err := c.List(ctx, client)
	if err != nil {
		return err
	}
	customerList = append(customerList, c)

	mpData := make(map[string][]Customer)
	mpData["List"] = customerList

	bytes, err := json.Marshal(&mpData)
	if err != nil {
		return err
	}

	var data map[string][]Customer
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}
	log.Printf("Adding Customer data: %+v", data)
	_, err = client.Collection("Customer").Doc("List").Set(ctx, data)
	return err
}

func (c Customer) List(ctx context.Context, client *firestore.Client) ([]Customer, error) {
	var list []Customer
	snap, _ := client.Collection("Customer").Doc("List").Get(ctx)
	dbByte, err := json.Marshal(snap.Data())
	if err != nil {
		return nil, err
	}
	var customerList map[string][]Customer
	if err := json.Unmarshal(dbByte, &customerList); err != nil {
		return nil, err
	}
	if val, ok := customerList["List"]; ok {
		for _, v := range val {
			list = append(list, v)
		}
	}
	return list, nil
}

func (c Customer) Update(ctx context.Context, client *firestore.Client) error {
	dataMap := structs.Map(c)
	log.Printf("Updating Customer data: %+v", dataMap)
	_, err := client.Collection("CustomerList").Doc(c.GetShopNameOwner()).Set(ctx, dataMap, firestore.MergeAll)
	return err
}

func (c Customer) Delete(ctx context.Context, client *firestore.Client) error {
	dataMap := structs.Map(c)
	log.Printf("Deleting Customer : %+v", dataMap)
	_, err := client.Collection("CustomerList").Doc(c.GetShopNameOwner()).Delete(ctx)
	return err
}

func (c Customer) GetShopNameOwner() string {
	return c.Shop + " (" + c.Owner + ")"
}
