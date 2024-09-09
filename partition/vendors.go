package partition

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/theinvincible/ecommerce-backend/config"
	"github.com/theinvincible/ecommerce-backend/models"
	"github.com/theinvincible/ecommerce-backend/utils"
	"golang.org/x/crypto/bcrypt"
)

/*
Key Features for the Vendor Handlers:

- Vendor Registration: Endpoint to create a new vendor account.
- Vendor Authentication: Login and token generation for vendors.
- Vendor Profile Management: Ability for vendors to update their profile.
- Product Management: Vendors can add, update, and delete products.
- Order Management: Vendors can view and manage orders.
- Sales Analytics: Provide sales data and analytics to vendors.

*/

func VendorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get user from context or session (assuming JWT token-based auth)
		user, ok := r.Context().Value("user").(*models.User)
		if !ok || user.Role != "vendor" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func VendorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Welcome, Vendor!"})
}

// Vendor registration handler
func CreateVendor(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var vendor models.User
	err := json.NewDecoder(r.Body).Decode(&vendor)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Check if username or email is already taken
	var existingVendor models.User
	if err := config.DB.Where("username = ? OR email = ?", vendor.Username, vendor.Email).First(&existingVendor).Error; err == nil {
		http.Error(w, "Username or email already taken", http.StatusConflict)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(vendor.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	vendor.Password = string(hashedPassword)
	vendor.CreatedAt = time.Now()
	vendor.UpdatedAt = time.Now()

	// Save the vendor to the database
	if err := config.DB.Create(&vendor).Error; err != nil {
		http.Error(w, "Error creating vendor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Account created successfully!"})
}

// Vendor authentication handler
func LoginVendor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vendor models.User
	err := json.NewDecoder(r.Body).Decode(&vendor)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var existingVendor models.User
	if err := config.DB.Where("username = ?", vendor.Username).First(&existingVendor).Error; err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingVendor.Password), []byte(vendor.Password))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(existingVendor.UserID)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// <=============================================Product Management=============================================>

func AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Get vendor ID from the JWT token or session.
	// This associates the product with the vendor who added it.
	// Ensuring that each product can be traced back to the vendor who created it, which is crucial for managing inventory, order processing, and overall business logic.
	vendorID := r.Context().Value("vendorID").(uint)
	product.ID = vendorID

	if err := config.DB.Create(&product).Error; err != nil {
		http.Error(w, "Error adding product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Get vendor ID from the JWT token or session.
	// This ensures that only the vendor who created the product can update it.
	vendorID := r.Context().Value("vendorID").(uint)
	product.ID = vendorID

	if err := config.DB.Save(&product).Error; err != nil {
		http.Error(w, "Error updating product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	productID := r.Context().Value("productID").(uint)

	if err := config.DB.Delete(&models.Product{}, productID).Error; err != nil {
		http.Error(w, "Error deleting product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted successfully"})
}

// <=============================================Order Management=============================================>

func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vendorID := r.Context().Value("vendorID").(uint)

	var orders []models.Order
	if err := config.DB.Where("vendor_id = ?", vendorID).Find(&orders).Error; err != nil {
		http.Error(w, "Error fetching orders", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	orderID := r.Context().Value("orderID").(uint)

	var order models.Order
	if err := config.DB.First(&order, orderID).Error; err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	orderID := r.Context().Value("orderID").(uint)

	if err := config.DB.Delete(&models.Order{}, orderID).Error; err != nil {
		http.Error(w, "Error deleting order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Order deleted successfully"})
}

// <=============================================Sales Analytics=============================================>

func GetSalesData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vendorID := r.Context().Value("vendorID").(uint)

	var salesData []models.Sales
	if err := config.DB.Where("vendor_id = ?", vendorID).Find(&salesData).Error; err != nil {
		http.Error(w, "Error fetching sales data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(salesData)
}

func GetSalesDataByProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vendorID := r.Context().Value("vendorID").(uint)
	productID := r.Context().Value("productID").(uint)

	var salesData []models.Sales
	if err := config.DB.Where("vendor_id = ? AND product_id = ?", vendorID, productID).Find(&salesData).Error; err != nil {
		http.Error(w, "Error fetching sales data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(salesData)
}

func GetSalesDataByDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vendorID := r.Context().Value("vendorID").(uint)

	var salesData []models.Sales
	if err := config.DB.Where("vendor_id = ?", vendorID).Find(&salesData).Error; err != nil {
		http.Error(w, "Error fetching sales data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(salesData)
}
