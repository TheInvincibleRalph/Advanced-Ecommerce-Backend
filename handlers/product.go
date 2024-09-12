package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/theinvincible/ecommerce-backend/config"
	"github.com/theinvincible/ecommerce-backend/models"
	"github.com/theinvincible/ecommerce-backend/utils"
	"gorm.io/gorm"
)

// CreateProduct creates a new product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := config.DB.Create(&product).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create the response map
	response := map[string]interface{}{
		"message": "Product created successfully",
		"product": product,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
	// json.NewEncoder(w).Encode(product)
}

// ProductService defines the interface for product-related operations.
type ProductService interface {
	CreateProduct(product *models.Product) error
	GetProducts() ([]models.Product, error)
	GetProductByID(id string) (*models.Product, error)
	UpdateProduct(product *models.Product) error
}

// CreateProductHandler handles HTTP requests to create a product.
func CreateProductHandler(ps ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var product models.Product
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := ps.CreateProduct(&product); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"message": "Product created successfully",
			"product": product,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

// ProductServiceImpl is a concrete implementation of ProductService.
type ProductServiceImpl struct {
	DB *gorm.DB
}

// CreateProduct creates a new product in the database.
func (ps *ProductServiceImpl) CreateProduct(product *models.Product) error {
	return ps.DB.Create(product).Error
}

// func GetProducts(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	// Retrieve pagination parameters
// 	pageStr := r.URL.Query().Get("page")
// 	limitStr := r.URL.Query().Get("limit")

// 	// Retrieve sort parameters
// 	sortBy := r.URL.Query().Get("sort_by")
// 	order := r.URL.Query().Get("order")

// 	// Retrieve filtering parameters
// 	category := r.URL.Query().Get("category")
// 	minPriceStr := r.URL.Query().Get("min_price")
// 	maxPriceStr := r.URL.Query().Get("max_price")

// 	// Retrive search parameters
// 	search := r.URL.Query().Get("search")

// 	// Define allowed sort columns
// 	allowedSortFields := map[string]bool{
// 		"name":     true,
// 		"price":    true,
// 		"category": true,
// 	}

// 	// Set default values for pagination
// 	page := 1
// 	limit := 10
// 	if pageStr != "" {
// 		var err error
// 		page, err = strconv.Atoi(pageStr)
// 		if err != nil {
// 			http.Error(w, "Invalid page number", http.StatusBadRequest)
// 			return
// 		}
// 	}
// 	if limitStr != "" {
// 		var err error
// 		limit, err = strconv.Atoi(limitStr)
// 		if err != nil {
// 			http.Error(w, "Invalid limit number", http.StatusBadRequest)
// 			return
// 		}
// 	}

// 	// Validate and set default sorting
// 	if _, valid := allowedSortFields[sortBy]; !valid {
// 		sortBy = "name" // Default to a safe column if invalid
// 	}
// 	if sortBy == "" {
// 		sortBy = "name"
// 	}
// 	if order == "" {
// 		order = "asc" // Default to ascending order
// 	}

// 	// Parse price filters
// 	var minPrice, maxPrice float64
// 	var err error
// 	if minPriceStr != "" {
// 		minPrice, err = strconv.ParseFloat(minPriceStr, 64)
// 		if err != nil {
// 			http.Error(w, "Invalid minimum price", http.StatusBadRequest)
// 			return
// 		}
// 	}
// 	if maxPriceStr != "" {
// 		maxPrice, err = strconv.ParseFloat(maxPriceStr, 64)
// 		if err != nil {
// 			http.Error(w, "Invalid maximum price", http.StatusBadRequest)
// 			return
// 		}
// 	}

// 	// Before querying the database, check if the results are already cached in Redis
// 	// If cached, return the results directly.
// 	// If not, proceed to fetching from the database, cache the results, and then return them.

// 	// First create a unique cache key based on the query parameters
// 	cacheKey := "products:" + pageStr + ":" + limitStr + ":" + sortBy + ":" + order + ":" + category + ":" + minPriceStr + ":" + maxPriceStr + ":" + search

// 	// Create a context with a timeout for Redis operations
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	cachedProducts, err := utils.InitRedisClient().Get(ctx, cacheKey).Result() //utils.InitRedisClient() returns an instance of a Redis db
// 	if err == redis.Nil {
// 		// Cache miss, fetch from the database
// 		// Query the database with pagination, sorting, filtering, and search
// 		var products []models.Product
// 		query := config.DB.Model(&models.Product{})
// 		if category != "" {
// 			query = query.Where("category = ?", category)
// 		}
// 		if minPriceStr != "" {
// 			query = query.Where("price >= ?", minPrice)
// 		}
// 		if maxPriceStr != "" {
// 			query = query.Where("price <= ?", maxPrice)
// 		}
// 		if search != "" {
// 			query = query.Where("name ILIKE ?", "%"+search+"%")
// 		}
// 		query = query.Offset((page - 1) * limit).Limit(limit)
// 		if strings.ToLower(order) == "desc" {
// 			query = query.Order(sortBy + " desc")
// 		} else {
// 			query = query.Order(sortBy + " asc")
// 		}
// 		if err := query.Find(&products).Error; err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		// Using Redis to cache the results of our database queries
// 		productsJSON, err := json.Marshal(products)
// 		if err == nil {
// 			utils.InitRedisClient().Set(ctx, cacheKey, productsJSON, 10*time.Minute)
// 		}

// 		if err != nil {
// 			// Redis error
// 			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		} else if cachedProducts != "" {
// 			// Cache hit, return the cached products
// 			w.Write([]byte(cachedProducts))
// 		} else {
// 			// Cache miss, return the products from the database
// 			w.Write(productsJSON)
// 		}
// 	}

// else if err != nil {
// 	// Redis error
// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// } else {
// 	// Cache hit, return the cached products
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write([]byte(cachedProducts))
// }
// }

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Retrieve pagination parameters
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	sortBy := r.URL.Query().Get("sort_by")
	order := r.URL.Query().Get("order")
	category := r.URL.Query().Get("category")
	minPriceStr := r.URL.Query().Get("min_price")
	maxPriceStr := r.URL.Query().Get("max_price")
	search := r.URL.Query().Get("search")

	// Default values and parameter parsing
	page := 1
	limit := 10
	if pageStr != "" {
		var err error
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			http.Error(w, "Invalid page number", http.StatusBadRequest)
			return
		}
	}
	if limitStr != "" {
		var err error
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, "Invalid limit number", http.StatusBadRequest)
			return
		}
	}

	if sortBy == "" {
		sortBy = "name"
	}
	if order == "" {
		order = "asc"
	}

	// Parse price filters
	var minPrice, maxPrice float64
	var err error
	if minPriceStr != "" {
		minPrice, err = strconv.ParseFloat(minPriceStr, 64)
		if err != nil {
			http.Error(w, "Invalid minimum price", http.StatusBadRequest)
			return
		}
	}
	if maxPriceStr != "" {
		maxPrice, err = strconv.ParseFloat(maxPriceStr, 64)
		if err != nil {
			http.Error(w, "Invalid maximum price", http.StatusBadRequest)
			return
		}
	}

	// Initialize Redis client
	redisClient := utils.InitRedisClient()
	cacheKey := fmt.Sprintf("products:%d:%d:%s:%s:%s:%s:%s:%s", page, limit, sortBy, order, category, minPriceStr, maxPriceStr, search)

	// Create a context with a timeout for Redis operations
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check Redis cache
	cachedProducts, err := redisClient.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		// Cache miss, fetch from the database
		var products []models.Product
		query := config.DB.Model(&models.Product{})
		if category != "" {
			query = query.Where("category = ?", category)
		}
		if minPriceStr != "" {
			query = query.Where("price >= ?", minPrice)
		}
		if maxPriceStr != "" {
			query = query.Where("price <= ?", maxPrice)
		}
		if search != "" {
			query = query.Where("name ILIKE ?", "%"+search+"%")
		}
		query = query.Offset((page - 1) * limit).Limit(limit)
		if strings.ToLower(order) == "desc" {
			query = query.Order(sortBy + " desc")
		} else {
			query = query.Order(sortBy + " asc")
		}
		if err := query.Find(&products).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Cache the result
		productsJSON, err := json.Marshal(products)
		if err == nil {
			redisClient.Set(ctx, cacheKey, productsJSON, 10*time.Minute)
		}
		w.Write(productsJSON)
	} else if err != nil {
		// Redis error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	} else {
		// Cache hit, return cached products
		w.Write([]byte(cachedProducts))
	}
}

// GetProductByID returns a product by ID
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	cacheKey := "product:" + id

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Use the initialized Redis client
	redisClient := utils.GetRedisClient()

	// Try to get the product from Redis cache
	cachedProduct, err := redisClient.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		// Cache miss, query the database
		log.Printf("Cache miss for product ID: %s", id)

		var product models.Product
		// Fetch the product from the database
		if err := config.DB.Where("id = ?", id).First(&product).Error; err != nil {
			log.Printf("Database error while fetching product with ID %s: %v", id, err)
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		// Cache the result in Redis
		productJSON, err := json.Marshal(product)
		if err != nil {
			log.Printf("Error marshalling product ID %s: %v", id, err)
			http.Error(w, "Error marshalling product data", http.StatusInternalServerError)
			return
		}

		// Store the product in Redis with a timeout
		if err := redisClient.Set(ctx, cacheKey, productJSON, 10*time.Minute).Err(); err != nil {
			log.Printf("Error caching product ID %s: %v", id, err)
		}

		// Return the result from the database
		log.Printf("Returning product data for ID: %s from database", id)
		w.Write(productJSON)
	} else if err != nil {
		// Log Redis connection or retrieval error
		log.Printf("Redis error while fetching product with ID %s: %v", id, err)
		http.Error(w, "Internal Server Error: Redis problem", http.StatusInternalServerError)
	} else {
		// Cache hit, return cached product
		log.Printf("Returning cached product data for ID: %s", id)
		w.Write([]byte(cachedProduct))
	}
}

// UpdateProduct updates a product by ID
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := config.DB.Save(&product).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

// DeleteProduct deletes a product by ID
func DeleteProduct(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	if err := config.DB.Delete(&models.Product{}, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted"})
}
