package models

import "gorm.io/gorm"

type Affliate struct {
	gorm.Model
	Name           string  `json:"name" gorm:"not null,index"`
	ReferralCode   string  `json:"referral_code" gorm:"unique;not null"`
	CommissionRate float64 `json:"commission_rate" gorm:"not null"`
	Active         bool    `json:"active" gorm:"default:true"`
	ProfileID      uint    `json:"profile_id" gorm:"not null"`
	Profile        Profile `gorm:"foreignkey:ProfileID"`
}
