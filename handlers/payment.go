package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"github.com/theinvincible/ecommerce-backend/models"
)

func Payment(w http.ResponseWriter, r *http.Request) {
	var paymentRequest models.Payment
	var shipping models.Shipping

	// Parse the JSON request body
	err := json.NewDecoder(r.Body).Decode(&paymentRequest)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Initialize Stripe with secret key
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	// Convert amount to cents
	amountInCents := int64(paymentRequest.Amount * 100)

	// Create charge parameters
	chargeParams := &stripe.ChargeParams{
		Amount:      stripe.Int64(amountInCents),
		Currency:    stripe.String("usd"),
		Description: stripe.String("Charge for order " + strconv.Itoa(paymentRequest.OrderID)),
	}

	// Add email for receipt
	chargeParams.ReceiptEmail = stripe.String(paymentRequest.Email)

	// Add metadata to charge parameters
	chargeParams.AddMetadata("order_id", strconv.Itoa(paymentRequest.OrderID))
	chargeParams.AddMetadata("transaction_id", paymentRequest.TransactionID)
	chargeParams.AddMetadata("payment_method", paymentRequest.PaymentMethod)

	// Add shipping details to metadata (for physical goods)
	chargeParams.AddMetadata("shipping_carrier", shipping.Carrier)
	chargeParams.AddMetadata("tracking_number", shipping.TrackingNumber)
	chargeParams.AddMetadata("shipping_method", shipping.ShippingMethod)
	chargeParams.AddMetadata("shipping_cost", strconv.FormatFloat(shipping.ShippingCost, 'f', 2, 64))
	chargeParams.AddMetadata("estimated_delivery", shipping.EstimatedDelivery.String())
	chargeParams.AddMetadata("shipping_date", shipping.ShippingDate)
	chargeParams.AddMetadata("shipping_type", shipping.ShippingType)
	chargeParams.AddMetadata("shipping_address", shipping.ShippingAddress)
	chargeParams.AddMetadata("shipping_city", shipping.ShippingCity)
	chargeParams.AddMetadata("shipping_state", shipping.ShippingState)
	chargeParams.AddMetadata("shipping_zip_code", shipping.ShippingZipCode)
	chargeParams.AddMetadata("shipping_country", shipping.ShippingCountry)

	// Create the charge
	ch, err := charge.New(chargeParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Prevents or handles duplicate charges gracefully
	chargeParams.IdempotencyKey = stripe.String(paymentRequest.TransactionID)

	// Update payment status and transaction ID
	paymentRequest.Status = ch.Status
	paymentRequest.TransactionID = ch.ID

	// Respond with the charge details
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ch)
}
