package config

import (
	"log"

	"github.com/theinvincible/ecommerce-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=56526681 dbname=ecommerce port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.AutoMigrate(
		&models.User{},
		&models.BlogPost{},
		&models.Affliate{},
		&models.Product{},
		&models.Category{},
		&models.Order{},
		&models.OrderItem{},
		&models.Cart{},
		&models.CartItem{},
		&models.Payment{},
		&models.Tag{},
		&models.Inventory{},
		&models.Review{},
		&models.Profile{},
		&models.Shipping{},
		&models.Notification{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}
}
