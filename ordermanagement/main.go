package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"ordermanagement/pkg/db"
	"ordermanagement/pkg/handlers"
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Welcome to the Order Management API",
	})
}

func main() {
    DB := db.Init()
	h := handlers.New(DB)

	router := mux.NewRouter()

	router.HandleFunc("/", welcomeHandler).Methods("GET")

	// Products
	router.HandleFunc("/products", h.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", h.GetProduct).Methods("GET")
	router.HandleFunc("/products", h.AddProduct).Methods("POST")
	router.HandleFunc("/products/{id}", h.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", h.DeleteProduct).Methods("DELETE")

	// Orders
	router.HandleFunc("/orders", h.GetOrders).Methods("GET")
	router.HandleFunc("/orders/{id}", h.GetOrder).Methods("GET")
	router.HandleFunc("/orders", h.AddOrder).Methods("POST")
	router.HandleFunc("/orders/{id}", h.UpdateOrder).Methods("PUT")
	router.HandleFunc("/orders/{id}", h.DeleteOrder).Methods("DELETE")

	// Payment Webhook
	router.HandleFunc("/webhook/payment", h.PaymentWebhookHandler).Methods("POST")

	// Start server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
