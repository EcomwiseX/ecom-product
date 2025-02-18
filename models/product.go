package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name          string  `json:"name" validate:"nonzero"`
	Description   string  `json:"description" validate:"nonzero"`
	Price         float64 `json:"price" validate:"nonzero"`
	DiscountPrice float64 `json:"discountPrice"`
	Stock         int     `json:"stock" validate:"nonzero"`
	Category      string  `json:"category"`
	ImageURL      string  `json:"imageURL"`
	IsActive      bool    `json:"isActive"`
}

func validateProduct(products *Product) error {
	if err := validator.Validate(products); err != nil {
		return err
	}
	return nil
}
