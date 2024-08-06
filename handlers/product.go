package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/theinvincible/ecommerce-backend/models"
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

	// Query the database with pagination and sorting
	var products []models.Product
	query := db.Model(&models.Product{})
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

	// Return the results
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// GetProduct returns a product by ID
func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var product models.Product
	if err := db.Preload("Category").First(&product, id).Error; err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(product)
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
