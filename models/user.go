package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           int            `json:"id" gorm:"primary_key,auto_increment"`
	Name         string         `json:"name" gorm:"not null,index"`
	Username     string         `json:"username" gorm:"unique,index,not null"`
	Password     string         `json:"password" gorm:"not null"`
	Email        string         `json:"email" gorm:"unique"`
	Profile      Profile        `json:"profile" gorm:"foreignKey:UserID"`
	Role         string         `json:"role" gorm:"default:customer"` // can be customer or admin or vendor
	Notification []Notification `json:"notification" gorm:"foreignKey:UserID"`
	DeviceToken  string         `json:"device_token"`

	// Vendor-specific fields. The check constraint only applies when the role is "vendor". Otherwise, these fields can be null ('')
	CompanyName     string `json:"company_name,omitempty" gorm:"check:(role != 'vendor' OR company_name <> ''),not null"`
	BusinessLicense string `json:"business_license,omitempty" gorm:"check:(role != 'vendor' OR business_license <> ''),not null"`
	Phone           string `json:"phone,omitempty"`
	Address         string `json:"address,omitempty"`
}
