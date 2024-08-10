package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/theinvincible/ecommerce-backend/models"
	"github.com/theinvincible/ecommerce-backend/utils"
)

// CreateProduct creates a new product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Create(&product).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	// Retrieve pagination parameters
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	// Retrieve sort parameters
	sortBy := r.URL.Query().Get("sort_by")
	order := r.URL.Query().Get("order")

	// Retrieve filtering parameters
	category := r.URL.Query().Get("category")
	minPriceStr := r.URL.Query().Get("min_price")
	maxPriceStr := r.URL.Query().Get("max_price")

	// Retrive search parameters
	search := r.URL.Query().Get("search")

	// Define allowed sort columns
	allowedSortFields := map[string]bool{
		"name":     true,
		"price":    true,
		"category": true,
	}

	// Set default values for pagination
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

	// Validate and set default sorting
	if _, valid := allowedSortFields[sortBy]; !valid {
		sortBy = "name" // Default to a safe column if invalid
	}
	if sortBy == "" {
		sortBy = "name"
	}
	if order == "" {
		order = "asc" // Default to ascending order
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

	// Before querying the database, check if the results are already cached in Redis
	// If cached, return the results directly.
	// If not, proceed to fetching from the database, cache the results, and then return them.

	// First create a unique cache key based on the query parameters
	cacheKey := "products:" + pageStr + ":" + limitStr + ":" + sortBy + ":" + order + ":" + category + ":" + minPriceStr + ":" + maxPriceStr + ":" + search

	cachedProducts, err := utils.InitRedisClient().Get(cacheKey).Result() //utils.InitRedisClient() returns an instance of a Redis db
	if err == redis.Nil {
		// Cache miss, fetch from the database
		// Query the database with pagination, sorting, filtering, and search
		var products []models.Product
		query := db.Model(&models.Product{})
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

		// Using Redis to cache the results of our database queries
		productsJSON, err := json.Marshal(products)
		if err == nil {
			utils.InitRedisClient().Set(cacheKey, productsJSON, 10*time.Minute)
		}

		// Return the result
		w.Header().Set("Content-Type", "application/json")
		w.Write(productsJSON)
	} else if err != nil {
		// Redis error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	} else {
		// Cache hit, return the cached products
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(cachedProducts))
	}
}

// GetProduct returns a product by ID
func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	// Try to get the product from cache
	cacheKey := "product:" + id

	cachedProduct, err := utils.InitRedisClient().Get(cacheKey).Result()
	if err == redis.Nil {
		// Cache miss, fetch from the database
		var product models.Product
		if err := db.Preload("Category").First(&product, id).Error; err != nil {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		// Cache the product data
		productJSON, err := json.Marshal(product)
		if err == nil {
			utils.InitRedisClient().Set(cacheKey, productJSON, 10*time.Minute)
		}

		// Return the result
		w.Header().Set("Content-Type", "application/json")
		w.Write(productJSON)
	} else if err != nil {
		// Redis error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	} else {
		// Cache hit, return the cached product
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(cachedProduct))
	}
}

// UpdateProduct updates a product by ID
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var product models.Product
	if err := db.First(&product, id).Error; err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Save(&product).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

// DeleteProduct deletes a product by ID
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	if err := db.Delete(&models.Product{}, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted"})
}
