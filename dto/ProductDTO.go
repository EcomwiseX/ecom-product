package dto

import (
	"time"

	"gopkg.in/validator.v2"
)

type ProductRequest struct {
	Name          string  `json:"name" validate:"nonzero" example:"Sabão de côco"`
	Description   string  `json:"description" validate:"nonzero" example:"Sabão de côco"`
	Price         float64 `json:"price" validate:"nonzero" example:"10"`
	DiscountPrice float64 `json:"discountPrice" example:"10"`
	Stock         int     `json:"stock" validate:"nonzero" example:"10"`
	Category      string  `json:"category" validate:"nonzero" example:"Limpeza"`
	ImageURL      string  `json:"imageURL" validate:"nonzero"`
	IsActive      bool    `json:"isActive" example:"true"`
}

type ProductResponse struct {
	ID            uint      `json:"id" example:"1"`
	Name          string    `json:"name" example:"Sabão de côco"`
	Description   string    `json:"description"`
	Price         float64   `json:"price"`
	DiscountPrice float64   `json:"discountPrice"`
	Stock         int       `json:"stock"`
	Category      string    `json:"category"`
	ImageURL      string    `json:"imageURL"`
	IsActive      bool      `json:"isActive"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	DeletedAt     time.Time `json:"deletedAt,omitempty"`
}

func ValidateProduct(products *ProductRequest) error {
	if err := validator.Validate(products); err != nil {
		return err
	}
	return nil
}
