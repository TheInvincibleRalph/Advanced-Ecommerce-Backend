package models

import (
	"time"

	"gorm.io/gorm"
)

type Shipping struct {
	gorm.Model
	OrderID           int       `json:"order_id" gorm:"not null"`
	Carrier           string    `json:"carrier" gorm:"not null"`
	TrackingNumber    string    `json:"tracking_number" gorm:"unique"`
	ShippingMethod    string    `json:"shipping_method" gorm:"not null"`
	ShippingCost      float64   `json:"shipping_cost"`
	EstimatedDelivery time.Time `json:"estimated_delivery"`
	Order             Order     `json:"order" gorm:"foreignKey:OrderID"`
	ShippingDate      string    `json:"shipping_date"`
	ShippingType      string    `json:"shipping_type"`
	ShippingAddress   string    `json:"shipping_address"`
	ShippingCity      string    `json:"shipping_city"`
	ShippingState     string    `json:"shipping_state"`
	ShippingZipCode   string    `json:"shipping_zip_code"`
	ShippingCountry   string    `json:"shipping_country"`
}
