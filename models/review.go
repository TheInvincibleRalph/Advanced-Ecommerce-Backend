package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	ID        int     `json:"id" gorm:"primary_key,auto_increment"`
	ProductID int     `json:"product_id" gorm:"not null"`
	UserID    int     `json:"user_id" gorm:"not null"`
	Rating    int     `json:"rating" gorm:"not null"`
	Comment   string  `json:"comment" gorm:"not null"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	User      User    `json:"user" gorm:"foreignKey:UserID;references:ID"`
}
