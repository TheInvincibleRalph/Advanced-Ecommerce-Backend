package models

import "time"

type Inventory struct {
	ID           int       `json:"id" gorm:"primary_key,auto_increment"`
	ProductID    int       `json:"product_id" gorm:"not null"`
	Quantity     int       `json:"quantity" gorm:"not null"`
	Product      Product   `json:"product" gorm:"foreignKey:ProductID"`
	StockLevel   int       `json:"stock_level" gorm:"not null"`
	ReorderLevel int       `json:"reorder_level"`
	LastRestock  time.Time `json:"last_restock"`
}
