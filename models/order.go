package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID             int         `json:"user_id" gorm:"not null"`
	ProductID          int         `json:"product_id" gorm:"not null"`
	Quantity           int         `json:"quantity" gorm:"not null"`
	TotalAmount        float64     `json:"total_amount" gorm:"not null"`
	OrderStatus        string      `json:"order_status" gorm:"not null"`
	OrderItems         []OrderItem `json:"order_items" gorm:"foreignKey:OrderItemID"` //represents the collection of ordered items belonging to a customer.
	PaymentMethod      string      `json:"payment_method"`
	OrderDate          string      `json:"order_date"`
	OrderTime          time.Time   `json:"order_time"`
	OrderTotal         string      `json:"order_total"`
	OrderDiscount      string      `json:"order_discount"`
	OrderPaymentStatus string      `json:"order_payment_status"`
	OrderID            int         `json:"order_id" gorm:"not null"`
	OrderItemID        int         `json:"order_item_id" gorm:"not null"`
}
