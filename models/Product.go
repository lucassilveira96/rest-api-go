package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id                  uint      `json:"id"`
	Description         string    `json:"description"`
	Price               float32   `json:"price"`
	CreatedAt           time.Time `json:"createdat"`
	UpdatedAt           time.Time `json:"updatedat"`
	ProductCategoriesId int
	ProductCategory     ProductCategory `gorm:"foreignKey:ProductCategoriesId" json:"ProductCategory"`
}

var Products []Product
