package models

import (
	"time"
)

type (
	// Product
	Products struct {
		ID                uint            `json:"id" gorm:"primary_key"`
		ProductCategoryID uint            `json:"product_category_ID"`
		Name              string          `json:"name"`
		Description       string          `json:"description"`
		Price             string          `json:"price"`
		CreatedAt         time.Time       `json:"created_at"`
		UpdatedAt         time.Time       `json:"updated_at"`
		OrderLines        []Orderline     `gorm:"foreignKey:ProductID" json:"-"`
		ProductCategory   ProductCategory `json:"-"`
	}
)
