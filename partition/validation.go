package partition

import (
	"errors"
	"strings"

	"github.com/theinvincible/ecommerce-backend/models"
)

// ValidateUser checks that all required fields are present based on the user's role.
// The user role will be extracted from the browser through the frontend.
func ValidateUser(user *models.User) error {
	// Common validation for all users
	if strings.TrimSpace(user.Username) == "" {
		return errors.New("username is required")
	}
	if strings.TrimSpace(user.Password) == "" {
		return errors.New("password is required")
	}
	if strings.TrimSpace(user.Email) == "" {
		return errors.New("email is required")
	}

	// Additional validation for vendors
	if user.Role == "vendor" {
		if strings.TrimSpace(user.CompanyName) == "" {
			return errors.New("company name is required for vendors")
		}
		if strings.TrimSpace(user.BusinessLicense) == "" {
			return errors.New("business license is required for vendors")
		}
	}

	// If validation passes
	return nil
}
