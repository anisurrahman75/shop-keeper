package data

import (
	"github.com/anisurrahman75/go-stock-management/api/models"
)

var UserList = []models.User{
	{
		FullName: "Anisur Rahman",
		Email:    "anisur123",
		Password: "123",
	},
	{
		FullName: "Nurul Islam",
		Email:    "nurul123",
		Password: "123",
	},
	{
		FullName: "Sayem Sheikh",
		Email:    "sayem123",
		Password: "123",
	},
}

var BrandList = []models.Brand{
	{
		Name: "SuperV",
	},
	{
		Name: "Mobil",
	},
	{
		Name: "Motul",
	},
	{
		Name: "Shell",
	},
}

var UnitList = []models.Unit{
	{
		Name: "1 Lit",
	},
	{
		Name: "5 Lit",
	},
	{
		Name: "10 Lit",
	},
	{
		Name: "15 Lit",
	},
	{
		Name: "20 Lit",
	},
	{
		Name: "205 Lit",
	},
}

var ProductList = []models.Product{
	{
		Brand: models.Brand{Name: "Mobil"},
		Name:  "Super™ 4T",
		Grade: "20W-50",
		Unit: models.Unit{
			Name: "1 Lit",
		},
		Qty: 10,
	},
	{
		Brand: models.Brand{Name: "BP"},
		Name:  "Super V",
		Grade: "20W-50",
		Unit: models.Unit{
			Name: "1 Lit",
		},
		Qty: 6,
	},
	{
		Brand: models.Brand{Name: "BP"},
		Name:  "VISCO 3000",
		Grade: " 20W-50",
		Unit: models.Unit{
			Name: "5 Lit",
		},
		Qty: 10,
	},
	{
		Brand: models.Brand{Name: "Motul"},
		Name:  "Motul 3000 4T Plus",
		Grade: "10W30",
		Unit: models.Unit{
			Name: "1 Lit",
		},
		Qty: 50,
	},
	{
		Brand: models.Brand{Name: "Mobil"},
		Name:  "Delvac™ 1300 Super",
		Grade: "15W-40",
		Unit: models.Unit{
			Name: "5 Lit",
		},
		Qty: 20,
	},
}

var CustomerList = []models.Customer{
	{
		ShopName:    "Ma Motors",
		OwnerName:   "Kamal Hossein",
		PhoneNumber: "+8801629397575",
		Address:     "Notun Bazar, Konabari, Gazipur",
		TotalDue:    12000,
	},
	{
		ShopName:    "Nolota Motors",
		OwnerName:   "Babul Sheikh",
		PhoneNumber: "+8801629397575",
		Address:     "Notun Bazar, Konabari, Gazipur",
		TotalDue:    1223456,
	},
	{
		ShopName:    "Rana Ratul Motors",
		OwnerName:   "Ripon Hossain",
		PhoneNumber: "+8801629397575",
		Address:     "Notun Bazar, Konabari, Gazipur",
		TotalDue:    8000,
	},
	{
		ShopName:    "MaShAllah Motors",
		OwnerName:   "Akram Hossain",
		PhoneNumber: "+8801629397575",
		Address:     "Notun Bazar, Konabari, Gazipur",
		TotalDue:    12000,
	},
}

var SalesList = []models.InvoiceData{
	{
		ShopName:  CustomerList[0].ShopName,
		OwnerName: CustomerList[0].OwnerName,
		Date:      "02-13-2023",
		InvoiceNo: "12",
		ProductsInfo: []models.SaleProductInfo{
			{},
		},
		NetTotal:          1200.00,
		DiscountInPercent: 12,
		SaveInDiscount:    123.00,
		GrandTotal:        1177.00,
	},
	{
		ShopName:  CustomerList[0].ShopName,
		OwnerName: CustomerList[0].OwnerName,
		Date:      "02-13-2023",
		InvoiceNo: "12",
		ProductsInfo: []models.SaleProductInfo{
			{},
		},
		NetTotal:          1200.00,
		DiscountInPercent: 13,
		SaveInDiscount:    123.00,
		GrandTotal:        1177.00,
	},
	{
		ShopName:  CustomerList[0].ShopName,
		OwnerName: CustomerList[0].OwnerName,
		Date:      "02-13-2023",
		InvoiceNo: "13",
		ProductsInfo: []models.SaleProductInfo{
			{},
		},
		NetTotal:          1200.00,
		DiscountInPercent: 12,
		SaveInDiscount:    123.00,
		GrandTotal:        1177.00,
	},
	{
		ShopName:  CustomerList[0].ShopName,
		OwnerName: CustomerList[0].OwnerName,
		Date:      "02-13-2023",
		InvoiceNo: "14",
		ProductsInfo: []models.SaleProductInfo{
			{},
		},
		NetTotal:          1200.00,
		DiscountInPercent: 12,
		SaveInDiscount:    123.00,
		GrandTotal:        1177.00,
	},
}
