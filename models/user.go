package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           int            `json:"id" gorm:"primary_key,auto_increment"`
	Name         string         `json:"name" gorm:"not null,index"`
	Username     string         `json:"username" gorm:"unique,index,not null"`
	Password     string         `json:"-" gorm:"not null"`
	Email        string         `json:"email" gorm:"unique"`
	UserID       int            `json:"user_id" gorm:"not null"`
	Profile      Profile        `json:"profile" gorm:"foreignKey:UserID"`
	Role         string         `json:"role" gorm:"default:customer"` // can be customer or admin or vendor
	Notification []Notification `json:"notification" gorm:"foreignKey:UserID"`
	DeviceToken  string         `json:"device_token"`

	// Vendor-specific fields (optional)
	CompanyName     string `json:"company_name,omitempty" gorm:"check:company_name <> ''"`
	BusinessLicense string `json:"business_license,omitempty" gorm:"check:business_license <> ''"` //omitempty is added just in case the customer and vendor is provided with the same form (which is not recommended). However, check constraints are added to ensure that the fields are not empty if the user is a vendor.
	Phone           string `json:"phone,omitempty"`
	Address         string `json:"address,omitempty"`
}
