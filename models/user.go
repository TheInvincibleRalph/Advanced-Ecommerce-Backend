package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `json:"id" gorm:"primary_key,auto_increment"`
	Name     string `json:"name" gorm:"not null,index"`
	Username string `json:"username" gorm:"unique,index,not null"`
	Password string `json:"-" gorm:"not null"`
	Email    string `json:"email" gorm:"unique"`
	Profile  Profile
	Role     string `json:"role" gorm:"default:customer"`
}
