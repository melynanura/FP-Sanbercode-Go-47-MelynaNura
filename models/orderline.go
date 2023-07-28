package models

import (
	"time"
)

type (
	// Orderline
	Orderline struct {
		ID        uint      `json:"id" gorm:"primary_key"`
		OrderID   uint      `json:"orderID"`
		ProductID uint      `json:"productID"`
		Quantity  int       `json:"quantity"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Order     Order     `json:"-"`
		Product   Products  `json:"-"`
	}
)
