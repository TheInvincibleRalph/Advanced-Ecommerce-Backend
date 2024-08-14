package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
	"github.com/theinvincible/ecommerce-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateOrder(t *testing.T) {
	// Step 1: Set up the mock database
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to open sqlmock database: %v", err)
	}
	defer dbMock.Close()

	dialector := postgres.New(postgres.Config{
		Conn: dbMock,
	})
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open gorm database: %v", err)
	}

	// Step 2: Prepare the input data
	order := models.Order{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		UserID:             1,
		ProductID:          1,
		Quantity:           2,
		TotalAmount:        100.0,
		OrderStatus:        "pending",
		PaymentMethod:      "cash",
		OrderDate:          "2021-01-01",
		OrderTime:          time.Time{},
		OrderTotal:         "100.0",
		OrderDiscount:      "0.0",
		OrderPaymentStatus: "pending",
		OrderID:            1,
		OrderItemID:        1,
		OrderItems: []models.OrderItem{
			{
				ProductID: 1,
				Quantity:  2,
				Price:     50.0,
				ID:        1,
				OrderID:   1,
				Product: models.Product{
					ID:          1,
					Name:        "Product 1",
					Description: "Product 1 Description",
					Price:       50.0,
					Quantity:    100,
					Image:       "product1.jpg",
					CategoryID:  1,
					Category: models.Category{
						ID:   1,
						Name: "Category 1",
					},
					Discount:        0.0,
					SKU:             "SKU001",
					Brand:           "Brand 1",
					Weight:          0.5,
					Dimensions:      "10x10x10",
					AverageRating:   4.5,
					NumberOfRatings: 100,
				},
				Total: 100.0,
			},
		},
	}

	orderJSON, _ := json.Marshal(order)

	// Step 3: Create a new HTTP request
	req, err := http.NewRequest("POST", "/api/v1/orders", bytes.NewBuffer(orderJSON))
	if err != nil {
		t.Fatalf("Failed to create a new request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Step 4: Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Step 5: Mock the database behavior
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "orders"`).
		WithArgs(
			sqlmock.AnyArg(), sqlmock.AnyArg(), nil,
			1, 1, 2, 100.0, "pending", "cash",
			"2021-01-01", sqlmock.AnyArg(),
			100.0, 0.0, "pending", 1, 1, 1,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Step 6: Create the mux router and define the route
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/orders", func(w http.ResponseWriter, r *http.Request) {
		CreateOrder(w, r, db)
	}).Methods("POST")

	// Step 7: Serve the HTTP request using the router
	router.ServeHTTP(rr, req)

	// Step 8: Assert the results
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body if needed
	expected := `{"ID":1,"UserID":1,"ProductID":1,"Quantity":2,"TotalAmount":100.0,"OrderStatus":"pending","PaymentMethod":"cash","OrderDate":"2021-01-01","OrderTime":"0001-01-01T00:00:00Z","OrderTotal":"100.0","OrderDiscount":"0.0","OrderPaymentStatus":"pending","OrderID":1,"OrderItems":[{"ProductID":1,"Quantity":2,"Price":50,"ID":1,"OrderID":1,"Total":100.0}]}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}
