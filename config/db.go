package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/theinvincible/ecommerce-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

var db *gorm.DB

func ConnectDatabase() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, user, password, dbname, port)

	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return
	}
	fmt.Println("Database connection successful")

	err = db.AutoMigrate(

		&models.CartItem{},
		&models.User{},
		&models.Cart{},
		&models.BlogPost{},
		&models.Affliate{},
		&models.Category{},
		&models.Product{},
		&models.OrderItem{},
		&models.Order{},
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
