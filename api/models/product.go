package models

type Product struct {
	Brand       Brand  `json:"brand"`
	Name        string `json:"name"`
	Grade       string `json:"grade"`
	Unit        Unit   `json:"unit"`
	Description string `json:"description"`
	Qty         int    `json:"qty"`
}
type Unit struct {
	Name string `json:"name"`
}
