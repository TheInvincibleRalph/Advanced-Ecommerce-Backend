package models

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	ID        int     `json:"id"`
	UserID    int     `json:"user_id"`
	ProductID int     `json:"product_id"`
	IsRead    bool    `json:"is_read"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	User      User    `json:"user" gorm:"foreignKey:UserID"`
	Title     string  `json:"title" gorm:"not null"`
	Message   string  `json:"message" gorm:"not null"`
}
