package models

import (
	"time"
)

type Order struct {
	ID            uint 			 `gorm:"primaryKey"`
	OrderProducts []OrderProduct  `gorm:"foreignKey:OrderID"`
	Quantity      int            `json:"quantity"`
	TotalPrice    float64        `json:"total_price"`
	Status        string         `json:"status"`
	PaymentStatus string         `json:"payment_status"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

type OrderProduct struct {
    ID        uint    `gorm:"primaryKey"`
    OrderID   uint    `gorm:"not null"`
    ProductID uint    `gorm:"not null"`
    Quantity  int     `gorm:"not null"`
    Price     float64 `gorm:"not null"`
}

// request payload (decodes JSON)
type OrderRequest struct {
    Products      []OrderItemInput `json:"products"`
    TotalPrice    float64          `json:"total_price"`
    Status        string           `json:"status"`
    PaymentStatus string           `json:"payment_status"`
    Quantity      int              `json:"quantity"`
}

type OrderItemInput struct {
    ID       uint    `json:"id"`
    Quantity int     `json:"quantity"`
    Price    float64 `json:"price"`
}