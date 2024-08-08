package models

type CheckoutRequest struct {
	UserID          uint   `json:"user_id" gorm:"not null"`
	ShippingAddress string `json:"shipping_address" gorm:"not null"`
	PaymentMethod   string `json:"payment_method" gorm:"not null"`
	DeliveryNotes   string `json:"delivery_notes"`
}
