package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/theinvincible/ecommerce-backend/config"
	"github.com/theinvincible/ecommerce-backend/handlers"
)

func main() {
	log.Println("Starting server...")

	router := mux.NewRouter()

	//Login routes
	router.HandleFunc("/api/v1/signup", handlers.SignUp).Methods("POST")
	router.HandleFunc("/api/v1/login", handlers.Login).Methods("POST")

	// User routes
	router.HandleFunc("/api/v1/users", handlers.GetUsers).Methods("POST")
	router.HandleFunc("/api/v1/users/{id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/api/v1/users/{id}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/v1/users/{id}", handlers.DeleteUser).Methods("DELETE")

	// Product routes
	router.HandleFunc("/api/v1/products", handlers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/v1/products", handlers.GetProducts).Methods("GET")
	router.HandleFunc("/api/v1/products/{id}", handlers.GetProduct).Methods("GET")
	router.HandleFunc("/api/v1/products/{id}", handlers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/v1/products/{id}", handlers.DeleteProduct).Methods("DELETE")

	// Order routes
	router.HandleFunc("/api/v1/orders", handlers.CreateOrder).Methods("POST")
	router.HandleFunc("/api/v1/orders", handlers.GetOrders).Methods("GET")
	router.HandleFunc("/api/v1/orders/{id}", handlers.GetOrder).Methods("GET")
	router.HandleFunc("/api/v1/orders/{id}", handlers.UpdateOrder).Methods("PUT")
	router.HandleFunc("/api/v1/orders/{id}", handlers.DeleteOrder).Methods("DELETE")

	// Category routes
	router.HandleFunc("/api/v1/categories", handlers.CreateCategory).Methods("POST")
	router.HandleFunc("/api/v1/categories", handlers.GetCategories).Methods("GET")
	router.HandleFunc("/api/v1/categories/{id}", handlers.GetCategory).Methods("GET")
	router.HandleFunc("/api/v1/categories/{id}", handlers.UpdateCategory).Methods("PUT")
	router.HandleFunc("/api/v1/categories/{id}", handlers.DeleteCategory).Methods("DELETE")

	// Cart routes
	router.HandleFunc("/api/v1/cart", handlers.CreateCart).Methods("POST")
	router.HandleFunc("/api/v1/cart/{id}", handlers.GetCart).Methods("GET")
	router.HandleFunc("/api/v1/cart/{id}", handlers.UpdateCart).Methods("PUT")
	router.HandleFunc("/api/v1/cart/{id}", handlers.DeleteCart).Methods("DELETE")

	// Payment routes
	router.HandleFunc("/api/v1/payment", handlers.PaymentHandler).Methods("POST")
	router.HandleFunc("/api/v1/webhook", handlers.WebhookHandler).Methods("POST")

	// Checkout routes
	router.HandleFunc("/api/v1/checkout", handlers.CheckoutHandler(config.DB)).Methods("POST")
	router.HandleFunc("/api/v1/order/confirm/{orderID}", handlers.OrderConfirmationHandler(config.DB)).Methods("POST")

	router.HandleFunc("/api/v1/store-device-token", handlers.StoreTokenHandler).Methods("POST")

	//<=====================================================MIddleware routes=====================================================>

	router.HandleFunc("/api/v1/admin", handlers.AdminHandler).Methods("GET").Subrouter().Use(handlers.RoleMiddleware("admin"))
	router.HandleFunc("/api/v1/admin/dashboard", handlers.AdminDashboardHandler).Methods("GET").Subrouter().Use(handlers.RoleMiddleware("admin"))
	router.HandleFunc("/api/v1/admin/users", handlers.GetUsersHandler).Methods("GET").Subrouter().Use(handlers.RoleMiddleware("admin"))
	router.HandleFunc("/api/v1/admin/users/{id}", handlers.UpdateUserHandler).Methods("POST").Subrouter().Use(handlers.RoleMiddleware("admin"))
	router.HandleFunc("/api/v1/admin/users/{id}", handlers.DeleteUserHandler).Methods("DELETE").Subrouter().Use(handlers.RoleMiddleware("admin"))
	router.HandleFunc("/api/v1/admin/products", handlers.AddProductHandler).Methods("POST").Subrouter().Use(handlers.RoleMiddleware("admin"))
	router.HandleFunc("/api/v1/admin/products/{id}", handlers.UpdateProductHandler).Methods("POST").Subrouter().Use(handlers.RoleMiddleware("admin"))
	router.HandleFunc("/api/v1/admin/products/{id}", handlers.DeleteProductHandler).Methods("DELETE").Subrouter().Use(handlers.RoleMiddleware("admin"))
	router.HandleFunc("/api/v1/admin/orders", handlers.GetOrdersHandler).Methods("GET").Subrouter().Use(handlers.RoleMiddleware("admin"))
	router.HandleFunc("/api/v1/admin/orders/{id}", handlers.UpdateOrderStatusHandler).Methods("POST").Subrouter().Use(handlers.RoleMiddleware("admin"))

	router.HandleFunc("/vendor", VendorHandler).Methods("GET").Subrouter().Use(handlers.RoleMiddleware("vendor"))
	router.HandleFunc("/customer", CustomerHandler).Methods("GET").Subrouter().Use(handlers.RoleMiddleware("customer"))
	router.HandleFunc("/dashboard", DashboardHandler).Methods("GET").Subrouter().Use(handlers.RoleMiddleware("admin", "vendor"))

	log.Println("Connecting to database...")

	config.ConnectDatabase()

	// rdbErr := utils.InitRedisClient()
	// if rdbErr != nil {
	// 	log.Fatal("Error connecting to Redis:", rdbErr)
	// } else {
	// 	log.Println("Connected to Redis successfully!")
	// }

	fmt.Println("Server is running on port 3001")

	err := http.ListenAndServe(":3001", router)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
