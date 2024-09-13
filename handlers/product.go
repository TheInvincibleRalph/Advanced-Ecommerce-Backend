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

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Retrieve pagination and filtering parameters
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	sortBy := r.URL.Query().Get("sort_by")
	order := r.URL.Query().Get("order")
	category := r.URL.Query().Get("category")
	minPriceStr := r.URL.Query().Get("min_price")
	maxPriceStr := r.URL.Query().Get("max_price")
	search := r.URL.Query().Get("search")

	// Set default pagination values
	page := 1
	limit := 10
	if pageStr != "" {
		if parsedPage, err := strconv.Atoi(pageStr); err == nil {
			page = parsedPage
		} else {
			http.Error(w, "Invalid page number", http.StatusBadRequest)
			return
		}
	}
	if limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil {
			limit = parsedLimit
		} else {
			http.Error(w, "Invalid limit number", http.StatusBadRequest)
			return
		}
	}

	// Set default sort and order values
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
	redisClient := utils.GetRedisClient()
	cacheKey := fmt.Sprintf("products:%d:%d:%s:%s:%s:%s:%s:%s", page, limit, sortBy, order, category, minPriceStr, maxPriceStr, search)

	// Create context for Redis operations
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Step 1: Check Redis cache
	cachedProducts, err := redisClient.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		// Step 2: Cache miss, fetch from database
		log.Println("Cache miss for products, fetching from database")

		var products []models.Product
		query := config.DB.Model(&models.Product{})

		// Apply filters
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

		// Apply pagination and sorting
		query = query.Offset((page - 1) * limit).Limit(limit)
		if strings.ToLower(order) == "desc" {
			query = query.Order(sortBy + " desc")
		} else {
			query = query.Order(sortBy + " asc")
		}

		// Execute query and fetch products
		if err := query.Find(&products).Error; err != nil {
			log.Printf("Database error: %v", err)
			http.Error(w, "Error fetching products", http.StatusInternalServerError)
			return
		}

		// Step 3: Cache the result in Redis
		productsJSON, err := json.Marshal(products)
		if err == nil {
			redisClient.Set(ctx, cacheKey, productsJSON, 10*time.Minute)
		} else {
			log.Printf("Error marshalling products: %v", err)
		}

		// Step 4: Return the result from the database
		w.Write(productsJSON)

	} else if err != nil {
		// Redis error, fallback to database
		log.Printf("Redis error: %v, falling back to database", err)
		// Continue to database fetch if Redis fails
	} else {
		// Step 5: Cache hit, return cached products
		log.Println("Cache hit, returning cached products")
		w.Write([]byte(cachedProducts))
		return
	}

	// Fallback to database if Redis fails
	var products []models.Product
	query := config.DB.Model(&models.Product{})

	// Apply filters again
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

	// Apply pagination and sorting
	query = query.Offset((page - 1) * limit).Limit(limit)
	if strings.ToLower(order) == "desc" {
		query = query.Order(sortBy + " desc")
	} else {
		query = query.Order(sortBy + " asc")
	}

	// Execute the database query
	if err := query.Find(&products).Error; err != nil {
		log.Printf("Database error: %v", err)
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}

	// Return products from database
	productsJSON, err := json.Marshal(products)
	if err != nil {
		log.Printf("Error marshalling products: %v", err)
		http.Error(w, "Error preparing products data", http.StatusInternalServerError)
		return
	}

	w.Write(productsJSON)
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

		// Marshal the product into JSON format
		productJSON, err := json.Marshal(product)
		if err != nil {
			log.Printf("Error marshalling product ID %s: %v", id, err)
			http.Error(w, "Error marshalling product data", http.StatusInternalServerError)
			return
		}

		// Cache the result in Redis for future requests (non-critical)
		if err := redisClient.Set(ctx, cacheKey, productJSON, 10*time.Minute).Err(); err != nil {
			// Log the error but do not return an error to the client
			log.Printf("Error caching product ID %s: %v", id, err)
		}

		// Return the product data from the database
		w.Write(productJSON)
		return
	} else if err != nil {
		// Redis is down or some error happened, just log and proceed to the database
		log.Printf("Redis is unavailable or other issue, querying database for product ID: %s", id)
	} else {
		// Cache hit, return cached product
		log.Printf("Returning cached product data for ID: %s", id)
		w.Write([]byte(cachedProduct))
		return
	}

	// If Redis fails, query the database and proceed
	var product models.Product
	if err := config.DB.Where("id = ?", id).First(&product).Error; err != nil {
		log.Printf("Database error while fetching product with ID %s: %v", id, err)
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// Marshal the product into JSON
	productJSON, err := json.Marshal(product)
	if err != nil {
		log.Printf("Error marshalling product ID %s: %v", id, err)
		http.Error(w, "Error marshalling product data", http.StatusInternalServerError)
		return
	}

	// Return the product data from the database
	w.Write(productJSON)
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
