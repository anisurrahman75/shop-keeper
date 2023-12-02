package models

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"errors"
	"github.com/fatih/structs"
	"google.golang.org/api/iterator"
	"log"
)

type Unit struct {
	Name string `json:"name"`
}

type Brand struct {
	Name string `json:"name"`
}

type Product struct {
	Brand       Brand  `json:"brand"`
	Name        string `json:"name"`
	Grade       string `json:"grade"`
	Unit        Unit   `json:"unit"`
	Description string `json:"description"`
	Qty         int    `json:"qty"`
}

func (p Product) Add(ctx context.Context, client *firestore.Client) error {
	dataMap := structs.Map(p)
	log.Printf("Adding Product data: %+v", dataMap)
	_, err := client.Collection("ProductList").Doc(p.Name).Set(ctx, dataMap)
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

func (u Unit) Add(ctx context.Context, client *firestore.Client) error {
	dataMap := structs.Map(u)
	log.Printf("Adding Unit data: %+v", dataMap)
	_, err := client.Collection("UnitList").Doc(u.Name).Set(ctx, dataMap)
	return err
}

func (u Unit) Update(ctx context.Context, client *firestore.Client) error {
	dataMap := structs.Map(u)
	log.Printf("Updating Unit data: %+v", dataMap)
	_, err := client.Collection("UnitList").Doc(u.Name).Set(ctx, dataMap, firestore.MergeAll)
	return err
}

func (u Unit) Delete(ctx context.Context, client *firestore.Client) error {
	dataMap := structs.Map(u)
	log.Printf("Deleting Unit : %+v", dataMap)
	_, err := client.Collection("UnitList").Doc(u.Name).Delete(ctx)
	return err
}

func (u Unit) List(ctx context.Context, client *firestore.Client) ([]Unit, error) {
	var list []Unit
	iter := client.Collection("UnitList").Documents(ctx)

	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, err
		}
		dbByte, err := json.Marshal(doc.Data())
		if err != nil {
			return nil, err
		}
		data := Unit{}
		if err := json.Unmarshal(dbByte, &data); err != nil {
			return nil, err
		}
		list = append(list, data)
	}
	return list, nil
}

func (b Brand) Add(ctx context.Context, client *firestore.Client) error {
	bytes, err := json.Marshal(&b)
	if err != nil {
		return err
	}
	var data map[string]interface{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}

	log.Printf("Adding Brand data: %+v", data)
	_, err = client.Collection("BrandList").Doc(b.Name).Set(ctx, data)
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
	iter := client.Collection("BrandList").Documents(ctx)

	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, err
		}
		dbByte, err := json.Marshal(doc.Data())
		if err != nil {
			return nil, err
		}
		data := Brand{}
		if err := json.Unmarshal(dbByte, &data); err != nil {
			return nil, err
		}
		list = append(list, data)
	}
	return list, nil
}
