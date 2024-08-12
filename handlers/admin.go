package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-redis/redis"
	"github.com/theinvincible/ecommerce-backend/models"
	"github.com/theinvincible/ecommerce-backend/utils"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Welcome, Admin!"})
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get user from context or session (assuming JWT token-based auth)
		user, ok := r.Context().Value("user").(*models.User)
		if !ok || user.Role != "admin" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func AdminDashboardHandler(w http.ResponseWriter, r *http.Request) {
	var dashboardData map[string]interface{}
	cacheKey := "admin_dashboard"

	// Attempt to get the cached dashboard data from Redis
	cachedData, err := utils.InitRedisClient().Get(cacheKey).Result()
	if err == redis.Nil {
		// Cache miss: Fetch from the database
		var userCount, productCount, orderCount int64

		if err := db.Model(&models.User{}).Count(&userCount).Error; err != nil {
			http.Error(w, "Failed to retrieve user count", http.StatusInternalServerError)
			return
		}

		if err := db.Model(&models.Product{}).Count(&productCount).Error; err != nil {
			http.Error(w, "Failed to retrieve product count", http.StatusInternalServerError)
			return
		}

		if err := db.Model(&models.Order{}).Count(&orderCount).Error; err != nil {
			http.Error(w, "Failed to retrieve order count", http.StatusInternalServerError)
			return
		}

		dashboardData = map[string]interface{}{
			"userCount":    userCount,
			"productCount": productCount,
			"orderCount":   orderCount,
		}

		// Cache the data in Redis with an expiration time
		jsonData, _ := json.Marshal(dashboardData)
		utils.InitRedisClient().Set(cacheKey, jsonData, time.Minute)
	} else if err != nil {
		// Redis error
		http.Error(w, "Failed to retrieve data", http.StatusInternalServerError)
		return
	} else {
		// Cache hit: Unmarshal the cached data
		json.Unmarshal([]byte(cachedData), &dashboardData)
	}

	// Return the data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dashboardData)
}
