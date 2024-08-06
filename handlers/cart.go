package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/theinvincible/ecommerce-backend/models"
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
