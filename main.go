package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/theinvincible/ecommerce-backend/config"
	"github.com/theinvincible/ecommerce-backend/handlers"
)

func main() {
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

	config.ConnectDatabase()
	fmt.Println("Server is running on port 8000")
	http.ListenAndServe(":8080", router)

}
