package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name            string   `json:"name" gorm:"not null,index"`
	Description     string   `json:"description" gorm:"not null"`
	Price           float64  `json:"price" gorm:"not null,index"`
	Quantity        int      `json:"quantity" gorm:"not null"`
	Image           string   `json:"image" gorm:"not null"`
	CategoryID      int      `json:"category_id" gorm:"not null"`
	Category        Category `json:"category" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Discount        float64  `json:"discount,omitempty" gorm:"type:decimal(10,2)"`
	SKU             string   `json:"sku,omitempty" gorm:"unique;not null"`
	Brand           string   `json:"brand,omitempty" gorm:"index"`
	Weight          float64  `json:"weight,omitempty" gorm:"type:decimal(10,2)"`
	Dimensions      string   `json:"dimensions,omitempty"`
	AverageRating   float64  `json:"average_rating,omitempty" gorm:"type:decimal(3,2)"`
	NumberOfRatings int      `json:"number_of_ratings,omitempty"`
}
