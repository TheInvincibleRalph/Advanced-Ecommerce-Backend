package models

type Sales struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	VendorID  uint    `json:"vendor_id"`
	ProductID uint    `json:"product_id"`
	Quantity  uint    `json:"quantity"`
	Amount    float64 `json:"amount"`
}
