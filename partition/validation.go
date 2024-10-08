package partition

import (
	"errors"
	"fmt"
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
			return fmt.Errorf("vendor must provide a company name")
		}
		if strings.TrimSpace(user.BusinessLicense) == "" {
			return fmt.Errorf("vendor must provide a business license")
		}
	}

	// If validation passes
	return nil
}
