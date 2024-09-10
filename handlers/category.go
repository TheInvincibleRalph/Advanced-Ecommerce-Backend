package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/theinvincible/ecommerce-backend/config"
	"github.com/theinvincible/ecommerce-backend/models"
)

// CreateCategory creates a new category
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := config.DB.Create(&category).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(category)
}

// GetCategories returns all categories
func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var categories []models.Category
	if err := config.DB.Preload("Products").Find(&categories).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(categories)
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var category models.Category
	if err := config.DB.Preload("Products").First(&category, id).Error; err != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(category)
}

// UpdateCategory updates a category by ID
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := config.DB.Save(&category).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(category)
}

// DeleteCategory deletes a category by ID
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	if err := config.DB.Delete(&models.Category{}, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Category deleted"})
}
