package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"ordermanagement/models"
)

func (h handler) GetOrders(w http.ResponseWriter, r *http.Request) {
	var orders []models.Order
	q := h.DB.Model(&models.Order{})
	// Pagination parameters
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	page := 1
	limit := 10
	if pageStr != "" && limitStr != "" {
		p, err1 := strconv.Atoi(pageStr)
		l, err2 := strconv.Atoi(limitStr)
		page = p
		limit = l
		if err1 == nil && err2 == nil && page > 0 && limit > 0 {
			offset := (page - 1) * limit
			q = q.Offset(offset).Limit(limit)
		}
	}
	if err := q.Find(&orders).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"orders": orders,
		"message":  "List of all orders",
		"meta": map[string]interface{}{
			"page":  page,
			"limit": limit,
		},
	})
}

// get order
func (h handler) GetOrder(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var order models.Order
	if err := h.DB.First(&order, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

// add order
func (h handler) AddOrder(w http.ResponseWriter, r *http.Request) {
	var req models.OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create order
	order := models.Order{
		Quantity:      0,
		TotalPrice:    0,
		Status:        "pending",
		PaymentStatus: "unpaid",
	}

	if err := h.DB.Create(&order).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Save order products
	for _, p := range req.Products {
		order.TotalPrice += p.Price * float64(p.Quantity)
		order.Quantity += p.Quantity

		op := models.OrderProduct{
			OrderID:   order.ID,
			ProductID: p.ID,
			Quantity:  p.Quantity,
			Price:     p.Price,
		}
		if err := h.DB.Create(&op).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Update order with total price
	if err := h.DB.Save(&order).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}


// update order
func (h handler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := uint(idInt) // Convert int to uint

	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order.ID = id
	if err := h.DB.Save(&order).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}


// delete order
func (h handler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.DB.Delete(&models.Order{}, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

// Payment Webhook Handler
func (h handler) PaymentWebhookHandler(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		OrderID       uint   `json:"order_id"`
		PaymentStatus string `json:"payment_status"`
	}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "Invalid payload", http.StatusBadRequest)
			return
		}

		var order models.Order
		if err := h.DB.First(&order, payload.OrderID).Error; err != nil {
			http.Error(w, "Order not found", http.StatusNotFound)
			return
		}

		if order.PaymentStatus != "unpaid" {
			http.Error(w, "Order payment status cannot be updated", http.StatusConflict)
			return
		}

		if payload.PaymentStatus != "paid" && payload.PaymentStatus != "failed" {
			http.Error(w, "Invalid payment status", http.StatusBadRequest)
			return
		}

		order.PaymentStatus = payload.PaymentStatus
		if err := h.DB.Save(&order).Error; err != nil {
			http.Error(w, "Failed to update payment status", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Payment status updated",
		})
	}