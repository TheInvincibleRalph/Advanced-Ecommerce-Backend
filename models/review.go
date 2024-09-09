package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	ProductID int     `json:"product_id" gorm:"not null"`
	UserID    int     `json:"user_id" gorm:"not null"`
	Rating    int     `json:"rating" gorm:"not null"`
	Comment   string  `json:"comment" gorm:"not null"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	User      User    `json:"user" gorm:"foreignKey:UserID;references:ID"`
}
