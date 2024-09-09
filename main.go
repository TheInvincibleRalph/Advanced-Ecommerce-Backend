package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/theinvincible/ecommerce-backend/config"
	"github.com/theinvincible/ecommerce-backend/handlers"
	"github.com/theinvincible/ecommerce-backend/partition"
)

func main() {

	log.Println("Connecting to database...")
	config.ConnectDatabase()

	// config.ReinitializeDatabase()

	// Check if the database connection is established
	if config.DB == nil {
		log.Fatal("Database connection failed")
	}

	// Set up router
	router := mux.NewRouter()

	//Login routes
	router.HandleFunc("/api/v1/signup", handlers.SignUp).Methods("POST")
	router.HandleFunc("/api/v1/login", handlers.Login).Methods("POST")

	// User routes
	router.HandleFunc("/api/v1/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/api/v1/users/{id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/api/v1/users/{id}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/v1/users/{id}", handlers.DeleteUser).Methods("DELETE")

	// Product routes
	router.HandleFunc("/api/v1/products", handlers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/v1/products", handlers.GetProducts).Methods("GET")
	router.HandleFunc("/api/v1/products/{id}", handlers.GetProduct).Methods("GET")
	router.HandleFunc("/api/v1/products/{id}", handlers.UpdateProduct).Methods("PUT")
	// router.HandleFunc("/api/v1/products/{id}", handlers.DeleteProduct(db)).Methods("DELETE")

	// Order routes
	router.HandleFunc("/api/v1/orders", handlers.CreateOrderHandler(config.DB)).Methods("POST")
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

	router.HandleFunc("/api/v1/admin", partition.AdminHandler).Methods("GET").Subrouter().Use(handlers.RoleMiddleware("admin"))
	router.HandleFunc("/api/v1/admin/dashboard", partition.AdminDashboardHandler).Methods("GET").Subrouter().Use(handlers.RoleMiddleware("admin"))
	router.HandleFunc("/api/v1/admin/users", partition.GetUsersHandler).Methods("GET").Subrouter().Use(handlers.RoleMiddleware("admin"))
	router.HandleFunc("/api/v1/admin/users/{id}", partition.UpdateUserHandler).Methods("POST").Subrouter().Use(handlers.RoleMiddleware("admin"))
	router.HandleFunc("/api/v1/admin/users/{id}", partition.DeleteUserHandler).Methods("DELETE").Subrouter().Use(handlers.RoleMiddleware("admin"))
	router.HandleFunc("/api/v1/admin/products", partition.AddProductHandler).Methods("POST").Subrouter().Use(handlers.RoleMiddleware("admin"))
	router.HandleFunc("/api/v1/admin/products/{id}", partition.UpdateProductHandler).Methods("POST").Subrouter().Use(handlers.RoleMiddleware("admin"))
	router.HandleFunc("/api/v1/admin/products/{id}", partition.DeleteProductHandler).Methods("DELETE").Subrouter().Use(handlers.RoleMiddleware("admin"))
	router.HandleFunc("/api/v1/admin/orders", partition.GetOrdersHandler).Methods("GET").Subrouter().Use(handlers.RoleMiddleware("admin"))
	router.HandleFunc("/api/v1/admin/orders/{id}", partition.UpdateOrderStatusHandler).Methods("POST").Subrouter().Use(handlers.RoleMiddleware("admin"))
	router.HandleFunc("/api/v1/admin/categories", partition.AssignRoleHandler).Methods("POST").Subrouter().Use(handlers.RoleMiddleware("admin"))

	router.HandleFunc("/api/v1/vendor", partition.VendorHandler).Methods("GET").Subrouter().Use(handlers.RoleMiddleware("vendor"))
	router.HandleFunc("/api/v1/login/{id}", partition.LoginVendor).Methods("POST").Subrouter().Use(handlers.RoleMiddleware("vendor"))
	router.HandleFunc("/api/v1/vendor/products", partition.AddProduct).Methods("POST").Subrouter().Use(handlers.RoleMiddleware("vendor"))
	router.HandleFunc("/api/v1/vendor/products/{id}", partition.UpdateProduct).Methods("PUT").Subrouter().Use(handlers.RoleMiddleware("vendor"))
	router.HandleFunc("/api/v1vendor/products/{id}", partition.DeleteProduct).Methods("DELETE").Subrouter().Use(handlers.RoleMiddleware("vendor"))
	router.HandleFunc("/api/v1/vendor/orders", partition.GetOrders).Methods("GET").Subrouter().Use(handlers.RoleMiddleware("vendor"))
	router.HandleFunc("/api/v1/vendor/orders/{id}", partition.DeleteOrder).Methods("DELETE").Subrouter().Use(handlers.RoleMiddleware("vendor"))
	router.HandleFunc("/api/v1/vendor/{id}", partition.GetSalesData).Methods("GET").Subrouter().Use(handlers.RoleMiddleware("vendor"))

	// router.HandleFunc("/customer", CustomerHandler).Methods("GET").Subrouter().Use(handlers.RoleMiddleware("customer"))
	// router.HandleFunc("/dashboard", DashboardHandler).Methods("GET").Subrouter().Use(handlers.RoleMiddleware("admin", "vendor"))

	// rdbErr := utils.InitRedisClient()
	// if rdbErr != nil {
	// 	log.Fatal("Error connecting to Redis:", rdbErr)
	// } else {
	// 	log.Println("Connected to Redis successfully!")
	// }

	log.Println("Starting server...")

	err := http.ListenAndServe(":3001", router)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	fmt.Println("Server is running on port 3001")

}
