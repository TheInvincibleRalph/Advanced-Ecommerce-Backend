package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/theinvincible/ecommerce-backend/models"
	"github.com/theinvincible/ecommerce-backend/utils"
	"gorm.io/gorm"
)

func CheckoutHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.CheckoutRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//Fetch cart items for the user
		var cartItems []models.CartItem
		db.Where("user_id = ?", req.UserID).Find(&cartItems)

		if len(cartItems) == 0 {
			http.Error(w, "Cart is empty", http.StatusBadRequest)
			return
		}

		//Create order

		order := models.Order{
			UserID:             int(req.UserID),
			TotalAmount:        0,
			OrderPaymentStatus: "Pending",
			OrderTime:          time.Now(),
			Quantity:           0,
			PaymentMethod:      req.PaymentMethod,
		}

		//Build order items and calculate total
		var orderItems []models.OrderItem //represents the collection of ordered items belonging to a customer.
		for _, item := range cartItems {
			orderItem := models.OrderItem{ //Creates an OrderItem struct for each cart item.
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Price:     item.Price,
				Total:     item.Total,
			}
			orderItems = append(orderItems, orderItem)
			order.TotalAmount += orderItem.Price * float64(orderItem.Quantity)
		}

		order.OrderItems = orderItems //populates the OrderItems field of the order struct (which is a placeholder for models.Order) with the orderItems slice.

		if err := db.Create(&order).Error; err != nil { //Saves the order and its associated items to the database.
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//Clear cart
		db.Where("user_id = ?", req.UserID).Delete(&models.CartItem{}) //Deletes all cart items for the user from the database, effectively clearing the user's cart.

		// Prepare order details for email
		orderDetails := fmt.Sprintf("Order ID: %d\nTotal: $%.2f", order.ID, order.TotalAmount)

		// Send confirmation email
		user, err := getUserByID(db, req.UserID)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		if err := utils.SendOrderConfirmationEmail(user.Email, orderDetails); err != nil {
			http.Error(w, "Order created but failed to send confirmation email", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"order_id": order.ID,
		})
	}
}

func OrderConfirmationHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		orderID := vars["orderID"]

		var order models.Order
		if err := db.First(&order, orderID).Error; err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		order.OrderStatus = "Completed"
		if err := db.Save(&order).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(order)
	}
}

// getUserByID retrieves a user from the database based on the provided user ID.
func getUserByID(db *gorm.DB, userID uint) (*models.User, error) {
	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
