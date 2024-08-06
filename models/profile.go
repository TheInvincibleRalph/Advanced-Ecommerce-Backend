package models

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID                  int       `json:"id" gorm:"primary_key,auto_increment"`
	Name                    string    `json:"name" gorm:"not null,index"`
	Username                string    `json:"username" gorm:"unique,index,not null"`
	Password                string    `json:"-" gorm:"not null"`
	Email                   string    `json:"email" gorm:"unique"`
	Phone                   string    `json:"phone" gorm:"unique"`
	Role                    string    `json:"role" gorm:"not null"` //admin, vendor, customer, or support
	Avatar                  string    `json:"avatar" gorm:"not null"`
	Active                  bool      `json:"active" gorm:"not null"`
	Address                 string    `json:"address"`
	City                    string    `json:"city"`
	State                   string    `json:"state"`
	ZipCode                 string    `json:"zip_code"`
	Country                 string    `json:"country"`
	LastLogin               string    `json:"last_login"`
	DateOfBirth             string    `json:"date_of_birth"`
	PreferredLanguage       string    `json:"preferred_language"`
	DateJoined              time.Time `json:"date_joined" gorm:"autoCreateTime"`
	TwoFactorEnabled        bool      `json:"two_factor_enabled" gorm:"default:false"`
	SubscriptionStatus      string    `json:"subscription_status"`
	FacebookLink            string    `json:"facebook_link"`
	TwitterLink             string    `json:"twitter_link"`
	LinkedInLink            string    `json:"linkedin_link"`
	NotificationPreferences string    `json:"notification_preferences"`
}
