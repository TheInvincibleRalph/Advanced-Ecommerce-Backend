package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	ID            int     `json:"id" gorm:"primary_key,auto_increment"`
	OrderID       int     `json:"order_id" gorm:"not null"`
	TransactionID string  `json:"transaction_id" gorm:"not null"`
	Amount        float64 `json:"amount" gorm:"not null"`
	Status        string  `json:"status" gorm:"not null"`
	Order         Order   `json:"order" gorm:"foreignKey:OrderID"`
	PaymentMethod string  `json:"payment_method" gorm:"not null"`
}
