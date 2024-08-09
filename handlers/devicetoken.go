package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/theinvincible/ecommerce-backend/models"
)

func UpdateDeviceToken(userID int, token string) error {
	err := db.Model(&models.User{}).Where("id = ?", userID).Update("device_token", token).Error
	return err
}

func GetDeviceToken(userID int) (string, error) {
	var token string
	err := db.Model(&models.User{}).Where("id = ?", userID).Select("device_token").Scan(&token).Error
	return token, err
}

// Handler function for storing the device token
func StoreTokenHandler(w http.ResponseWriter, r *http.Request) {
	var req models.User

	// Decode the incoming JSON request body
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Store the device token
	err = UpdateDeviceToken(req.UserID, req.DeviceToken)
	if err != nil {
		http.Error(w, "Failed to store token", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Token stored successfully"))
}
