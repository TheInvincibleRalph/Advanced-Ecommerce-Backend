package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/theinvincible/ecommerce-backend/models"
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
