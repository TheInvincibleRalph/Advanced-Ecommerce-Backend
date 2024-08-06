package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID      int       `json:"id" gorm:"primary_key,auto_increment"`
	Name    string    `json:"name" gorm:"not null,index,unique"`
	Product []Product `json:"product" gorm:"foreignKey:CategoryID"`
}
