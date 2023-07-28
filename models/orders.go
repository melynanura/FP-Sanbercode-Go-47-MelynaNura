package models

import (
	"time"
)

type (
	// Order
	Order struct {
		ID         uint        `json:"id" gorm:"primary_key"`
		OrderDate  time.Time   `json:"order_date"`
		Address    string      `json:"address"`
		CustomerID uint        `json:"customerID"`
		CreatedAt  time.Time   `json:"created_at"`
		UpdatedAt  time.Time   `json:"updated_at"`
		Customer   Customer    `json:"-"`
		OrderLines []Orderline `json:"-"`
	}
)
