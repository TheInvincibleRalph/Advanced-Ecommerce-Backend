package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/theinvincible/ecommerce-backend/models"
	"gorm.io/gorm"
)

func checkoutHandler(db *gorm.DB) http.HandlerFunc {
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

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"order_id": order.ID,
		})
	}
}
