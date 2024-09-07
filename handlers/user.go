package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/theinvincible/ecommerce-backend/models"
	"github.com/theinvincible/ecommerce-backend/partition"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// This calls the validation function before saving the user to the database
	if err := partition.ValidateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save user to the database
	if err := db.Create(&user).Error; err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUser returns a user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"] // Extract the user ID from the request URL
	var user models.User
	if err := db.First(&user, id).Error; err != nil { // Find user by ID
		http.Error(w, "User not found", http.StatusNotFound) // Handle error if user not found
		return
	}

	json.NewEncoder(w).Encode(user) // Send the user details as JSON
}

// GetUsers returns all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

// UpdateUser updates a user by ID
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"] // Extract the user ID from the request URL
	var user models.User
	if err := db.First(&user, id).Error; err != nil { // Find user by ID
		http.Error(w, "User not found", http.StatusNotFound) // Handle error if user not found
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil { // Decode the updated user data
		http.Error(w, err.Error(), http.StatusBadRequest) // Handle error if decoding fails
		return
	}

	if err := db.Save(&user).Error; err != nil { // Save the updated user information to the database
		http.Error(w, err.Error(), http.StatusInternalServerError) // Handle error if saving fails
		return
	}

	json.NewEncoder(w).Encode(user) // Send the updated user details as JSON
}

// DeleteUser deletes a user by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]                                     // Extract the user ID from the request URL
	if err := db.Delete(&models.User{}, id).Error; err != nil { // Delete the user by ID
		http.Error(w, err.Error(), http.StatusInternalServerError) // Handle error if deleting fails
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted"}) // Send success message
}
