package models

type Customer struct {
	ShopName    string `json:"shop_name"`
	OwnerName   string `json:"owner_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	TotalDue    int    `json:"total_due"`
}
