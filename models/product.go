package models

import (
	"time"
)

type Product struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Name          string    `json:"name" validate:"nonzero" example:"Sabão de côco"`
	Description   string    `json:"description" validate:"nonzero" example:"Sabão de côco"`
	Price         float64   `json:"price" validate:"nonzero" example:"10"`
	DiscountPrice float64   `json:"discountPrice" example:"10"`
	Stock         int       `json:"stock" validate:"nonzero" example:"10"`
	Category      string    `json:"category" validate:"nonzero" example:"Limpeza"`
	ImageURL      string    `json:"imageURL" validate:"nonzero"`
	IsActive      bool      `json:"isActive" example:"true"`
	CreatedAt     time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"-" gorm:"autoUpdateTime"`
	DeletedAt     time.Time `json:"-" gorm:"index"`
}
