package models

import "gorm.io/gorm"

type Affliate struct {
	gorm.Model
	ID             int     `json:"id" gorm:"primary_key,auto_increment"`
	Name           string  `json:"name" gorm:"not null,index"`
	ReferralCode   string  `json:"referral_code" gorm:"unique;not null"`
	CommissionRate float64 `json:"commission_rate" gorm:"not null"`
	Active         bool    `json:"active" gorm:"default:true"`
	Profile        Profile
}
