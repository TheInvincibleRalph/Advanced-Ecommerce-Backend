package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/sub"
	"github.com/theinvincible/ecommerce-backend/models"
)

// This handles one-time payments using Stripe
func PaymentHandler(w http.ResponseWriter, r *http.Request) {
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

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(1)

	// Channel to receive the charge result
	resultChan := make(chan *stripe.Charge)
	errorChan := make(chan error)

	go func() {
		defer wg.Done()
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

		// Prevents or handles duplicate charges gracefully
		chargeParams.IdempotencyKey = stripe.String(paymentRequest.TransactionID)

		// Retry logic in case of failure
		maxRetries := 3
		var ch *stripe.Charge
		for i := 0; i < maxRetries; i++ {
			ch, err = charge.New(chargeParams)
			if err != nil {
				if stripeErr, ok := err.(*stripe.Error); ok {
					// Retry on certain transient errors
					if stripeErr.Type == stripe.ErrorTypeAPIConnection || stripeErr.Code == "lock_timeout" {
						log.Printf("Transient error: %v. Retrying...", err)
						time.Sleep(2 * time.Second)
						continue
					}
				}
				// Send the error to the error channel if not retryable
				errorChan <- err
				return
			}
			// If charge is successful, send it to result channel
			resultChan <- ch
			return
		}
		// If all retries are exhausted, send the last error to the error channel
		errorChan <- err
	}()
	go func() {
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	select {
	case ch := <-resultChan:
		// Update payment status and transaction ID
		paymentRequest.Status = ch.Status
		paymentRequest.TransactionID = ch.ID

		// Respond with the charge details
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ch)

	case err := <-errorChan:
		log.Printf("Stripe charge creation failed: %v", err)
		http.Error(w, "Payment processing failed", http.StatusInternalServerError)

	case <-ctx.Done():
		log.Printf("Request timed out")
		http.Error(w, "Request timed out", http.StatusRequestTimeout)
	}

}

// This handles recurring payments using Stripe
func CreateSubscription(customerID, planID string) (*stripe.Subscription, error) {

	// Initialize Stripe with secret key
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	subParams := &stripe.SubscriptionParams{
		Customer: stripe.String(customerID),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Plan: stripe.String(planID),
			},
		},
	}

	return sub.New(subParams)
}

// This creates a new customer in Stripe

func CreateCustomer(email, name string) (*stripe.Customer, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	customerParams := &stripe.CustomerParams{
		Email: stripe.String(email),
		Name:  stripe.String(name),
	}

	return customer.New(customerParams)
}
