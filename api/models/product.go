package models

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"github.com/fatih/structs"
	"log"
)

type Unit struct {
	Name string `json:"name"`
}

type Brand struct {
	Name string `json:"name"`
}

type Product struct {
	Name        string `json:"Name"`
	Brand       string `json:"Brand"`
	Description string `json:"Description"`
	Grade       string `json:"Grade"`
	Qty         int64  `json:"Qty"`
	Unit        string `json:"Unit"`
}

func (p Product) Add(ctx context.Context, client *firestore.Client) error {
	productList, err := p.List(ctx, client)
	if err != nil {
		return err
	}
	productList = append(productList, p)

	mpData := make(map[string][]Product)
	mpData["List"] = productList

	bytes, err := json.Marshal(&mpData)
	if err != nil {
		return err
	}

	var data map[string][]Product
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}
	log.Printf("Adding Product data: %+v", data)
	_, err = client.Collection("Product").Doc("List").Set(ctx, data)
	return err
}

func (p Product) Update(ctx context.Context, client *firestore.Client) error {
	dataMap := structs.Map(p)
	log.Printf("Updating Product data: %+v", dataMap)
	_, err := client.Collection("ProductList").Doc(p.Name).Set(ctx, dataMap, firestore.MergeAll)
	return err
}

func (p Product) Delete(ctx context.Context, client *firestore.Client) error {
	dataMap := structs.Map(p)
	log.Printf("Deleting Product : %+v", dataMap)
	_, err := client.Collection("ProductList").Doc(p.Name).Delete(ctx)
	return err
}

func (p Product) List(ctx context.Context, client *firestore.Client) ([]Product, error) {
	var list []Product
	snap, _ := client.Collection("Product").Doc("List").Get(ctx)
	dbByte, err := json.Marshal(snap.Data())
	if err != nil {
		return nil, err
	}
	var productList map[string][]Product
	if err := json.Unmarshal(dbByte, &productList); err != nil {
		return nil, err
	}
	if val, ok := productList["List"]; ok {
		for _, v := range val {
			list = append(list, v)
		}
	}
	return list, nil
}

func (u Unit) Add(ctx context.Context, client *firestore.Client) error {
	unitList, err := u.List(ctx, client)
	if err != nil {
		return err
	}
	unitList = append(unitList, u)

	var stringSlice []string
	for _, i := range unitList {
		stringSlice = append(stringSlice, i.Name)
	}
	mpData := make(map[string][]string)
	mpData["List"] = stringSlice

	bytes, err := json.Marshal(&mpData)
	if err != nil {
		return err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}
	log.Printf("Adding Unit data: %+v", u)
	_, err = client.Collection("Unit").Doc("List").Set(ctx, data)
	return err
}

func (u Unit) Delete(ctx context.Context, client *firestore.Client) error {
	unitList, err := u.List(ctx, client)
	if err != nil {
		return err
	}
	for index, i := range unitList {
		if i.Name == u.Name {
			unitList = append(unitList[0:index], unitList[index+1:]...)
		}
	}

	var stringSlice []string
	for _, i := range unitList {
		stringSlice = append(stringSlice, i.Name)
	}

	mpData := make(map[string][]string)
	mpData["List"] = stringSlice

	bytes, err := json.Marshal(&mpData)
	if err != nil {
		return err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}
	log.Printf("Adding Unit data: %+v", u)
	_, err = client.Collection("Unit").Doc("List").Set(ctx, data)
	return err
}

func (u Unit) List(ctx context.Context, client *firestore.Client) ([]Unit, error) {
	var list []Unit
	snap, _ := client.Collection("Unit").Doc("List").Get(ctx)
	dbByte, err := json.Marshal(snap.Data())
	if err != nil {
		return nil, err
	}
	var unitList map[string]interface{}
	if err := json.Unmarshal(dbByte, &unitList); err != nil {
		return nil, err
	}
	if val, ok := unitList["List"].([]interface{}); ok {
		for _, v := range val {
			if str, ok := v.(string); ok {
				list = append(list, Unit{Name: str})
			}
		}
	}
	return list, nil
}

func (b Brand) Add(ctx context.Context, client *firestore.Client) error {
	brandList, err := b.List(ctx, client)
	if err != nil {
		return err
	}
	brandList = append(brandList, b)

	var stringSlice []string
	for _, i := range brandList {
		stringSlice = append(stringSlice, i.Name)
	}
	mpData := make(map[string][]string)
	mpData["List"] = stringSlice

	bytes, err := json.Marshal(&mpData)
	if err != nil {
		return err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}
	log.Printf("Adding Brand data: %+v", b)
	_, err = client.Collection("Brand").Doc("List").Set(ctx, data)
	return err
}

func (b Brand) Update(ctx context.Context, client *firestore.Client) error {
	dataMap := structs.Map(b)
	log.Printf("Updating Brand data: %+v", dataMap)
	_, err := client.Collection("BrandList").Doc(b.Name).Set(ctx, dataMap, firestore.MergeAll)
	return err
}

func (b Brand) Delete(ctx context.Context, client *firestore.Client) error {
	dataMap := structs.Map(b)
	log.Printf("Deleting Brand : %+v", dataMap)
	_, err := client.Collection("BrandList").Doc(b.Name).Delete(ctx)
	return err
}

func (b Brand) List(ctx context.Context, client *firestore.Client) ([]Brand, error) {
	var list []Brand
	snap, _ := client.Collection("Brand").Doc("List").Get(ctx)
	dbByte, err := json.Marshal(snap.Data())
	if err != nil {
		return nil, err
	}
	var brandList map[string]interface{}
	if err := json.Unmarshal(dbByte, &brandList); err != nil {
		return nil, err
	}
	if val, ok := brandList["List"].([]interface{}); ok {
		for _, v := range val {
			if str, ok := v.(string); ok {
				list = append(list, Brand{Name: str})
			}
		}
	}
	return list, nil
}
