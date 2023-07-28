package models

import (
	"time"
)

type (
	// Product
	ProductCategory struct {
		ID        uint       `gorm:"primary_key" json:"id"`
		Category  string     `json:"category"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		Products  []Products `json:"-"`
	}
)
