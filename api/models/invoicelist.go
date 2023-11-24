package models

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
	NetTotal          float64           `json:"NetTotal"`
	DiscountInPercent int               `json:"DiscountInPercent"`
	SaveInDiscount    float64           `json:"SaveInDiscount"`
	GrandTotal        float64           `json:"GrandTotal"`
}
