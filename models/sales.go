package models

import "gorm.io/gorm"

type Sales struct {
	gorm.Model
	VendorID  uint    `json:"vendor_id"`
	ProductID uint    `json:"product_id"`
	Quantity  uint    `json:"quantity"`
	Amount    float64 `json:"amount"`
	Date      string  `json:"date"`
}
