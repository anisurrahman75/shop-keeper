package models

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
