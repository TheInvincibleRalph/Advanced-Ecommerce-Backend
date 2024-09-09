package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string         `json:"name" gorm:"not null,index"`
	Username     string         `json:"username" gorm:"unique,index,not null"`
	Password     string         `json:"password" gorm:"not null"`
	UserID       int            `json:"user_id" gorm:"not null"`
	Email        string         `json:"email" gorm:"unique"`
	Profile      Profile        `json:"profile" gorm:"foreignKey:UserID"`
	Role         string         `json:"role" gorm:"default:customer"` // can be customer or admin or vendor
	Notification []Notification `json:"notification" gorm:"foreignKey:UserID"`
	DeviceToken  string         `json:"device_token"`

	// Vendor-specific fields. The check constraint only applies when the role is "vendor". Otherwise, these fields can be null ('')
	CompanyName     string `json:"company_name,omitempty"`
	BusinessLicense string `json:"business_license,omitempty"`
	Phone           string `json:"phone,omitempty"`
	Address         string `json:"address,omitempty"`
}
