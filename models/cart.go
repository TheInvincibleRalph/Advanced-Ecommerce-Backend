package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID    int        `json:"user_id" gorm:"not null"`
	Items     []CartItem `json:"items" gorm:"foreignKey:CartID"`
	Total     float64    `json:"total" gorm:"not null"`
	TotalItem int        `json:"total_item" gorm:"not null"`
	User      User       `json:"user" gorm:"foreignKey:UserID"`
}
