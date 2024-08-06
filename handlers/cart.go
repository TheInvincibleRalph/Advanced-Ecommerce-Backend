package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/theinvincible/ecommerce-backend/models"
	"gorm.io/gorm"
)

func CreateCart(w http.ResponseWriter, r *http.Request) {
	var cart models.Cart

	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Create(&cart).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

func GetCart(w http.ResponseWriter, r *http.Request) {
	// Extract variables from the URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid cart ID", http.StatusNotFound)
		return
	}

	// Retrieve the cart from the database
	var cart models.Cart
	if err := db.Preload("Items").Find(&cart, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Cart not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Respond with the cart details
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}
