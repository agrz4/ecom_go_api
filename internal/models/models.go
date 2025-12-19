package models

import (
	"time"
)

type Product struct {
	ID             int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name           string    `gorm:"not null" json:"name"`
	PriceInCenters int32     `gorm:"not null" json:"price_in_centers"`
	Quantity       int32     `gorm:"not null" json:"quantity"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type Order struct {
	ID         int64       `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerID int64       `gorm:"not null" json:"customer_id"`
	CreatedAt  time.Time   `gorm:"autoCreateTime" json:"created_at"`
	Items      []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
}

type OrderItem struct {
	ID         int64 `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID    int64 `gorm:"not null" json:"order_id"`
	ProductID  int64 `gorm:"not null" json:"product_id"`
	Quantity   int32 `gorm:"not null" json:"quantity"`
	PriceCents int32 `gorm:"not null" json:"price_cents"`
}
