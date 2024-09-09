package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID   int     `json:"order_id" gorm:"not null"`
	ProductID int     `json:"product_id" gorm:"not null"`
	Quantity  int     `json:"quantity" gorm:"not null"`
	Price     float64 `json:"price" gorm:"not null"`
	Total     float64 `json:"total" gorm:"not null"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
}
