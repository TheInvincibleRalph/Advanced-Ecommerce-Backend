package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CartID    int     `json:"cart_id" gorm:"not null"`
	ProductID int     `json:"product_id" gorm:"not null"`
	Quantity  int     `json:"quantity" gorm:"not null"`
	Price     float64 `json:"price" gorm:"not null"`
	Total     float64 `json:"total" gorm:"not null"`
	Cart      Cart    `json:"cart" gorm:"foreignKey:CartID"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
}
