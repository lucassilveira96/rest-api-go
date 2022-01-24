package models

import (
	"time"
)

type ProductCategory struct {
	Id          uint      `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdat"`
	UpdatedAt   time.Time `json:"updatedat"`
}

var ProductsCategory []ProductCategory
