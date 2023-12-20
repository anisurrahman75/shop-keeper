package db

import (
	"context"
	"github.com/anisurrahman75/go-stock-management/api/models"
	"github.com/anisurrahman75/go-stock-management/pkg/routes"
)

func Product(server *routes.Server) error {
	p := models.Product{
		Brand: "Mobil",
		Name:  "Super™ 4T",
		Grade: "20W-50",
		Unit:  "1 Lit",
		Qty:   100,
	}
	if err := p.Add(context.Background(), server.Client); err != nil {
		return err
	}

	//if err := p.Delete(context.Background(), server.Client); err != nil {
	//	return err
	//}

	return nil
}

func Brand(server *routes.Server) error {
	b := models.Brand{
		Name: "SuperV",
	}
	if err := b.Add(context.Background(), server.Client); err != nil {
		return err
	}

	//if err := p.Delete(context.Background(), server.Client); err != nil {
	//	return err
	//}

	return nil
}

func Unit(server *routes.Server) error {
	u := models.Unit{
		Name: "1 Lit",
	}
	if err := u.Add(context.Background(), server.Client); err != nil {
		return err
	}

	//if err := p.Delete(context.Background(), server.Client); err != nil {
	//	return err
	//}

	return nil
}

func Customer(server *routes.Server) error {
	c := models.Customer{
		Shop:        "Ma Motors",
		Owner:       "Kamal Hossein",
		PhoneNumber: "+8801629397575",
		Address:     "Notun Bazar, Konabari, Gazipur",
		TotalDue:    12000,
	}
	if err := c.Add(context.Background(), server.Client); err != nil {
		return err
	}

	//if err := p.Delete(context.Background(), server.Client); err != nil {
	//	return err
	//}

	return nil
}

//func Invoice(server *routes.Server) error {
//	i := models.InvoiceData{
//		CustomerInfo: data.CustomerList[0],
//		Date:         "02-13-2023",
//		InvoiceNo:    12,
//		ProductsInfo: []models.SaleProductInfo{
//			{
//				Index:     1,
//				Name:      data.ProductList[0].Name,
//				UnitPrice: "12.00",
//				Qty:       "12",
//				SubTotal:  "244.00",
//			},
//			{
//				Index:     2,
//				Name:      data.ProductList[0].Name,
//				UnitPrice: "12.00",
//				Qty:       "12",
//				SubTotal:  "244.00",
//			},
//		},
//		NetTotal:          1200.00,
//		DiscountInPercent: 12,
//		SaveInDiscount:    123.00,
//		GrandTotal:        1177.00,
//	}
//
//	if err := i.Add(context.Background(), server.Client); err != nil {
//		return err
//	}
//	return nil
//
//}

func InitDatabase(server *routes.Server) error {
	//if err := Product(server); err != nil {
	//	return err
	//}

	//if err := Brand(server); err != nil {
	//	return err
	//}

	//if err := Unit(server); err != nil {
	//	return err
	//}

	//if err := Customer(server); err != nil {
	//	return err
	//}

	//if err := Invoice(server); err != nil {
	//	return err
	//}

	return nil
}
