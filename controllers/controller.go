package controllers

import (
	"net/http"
	"strconv"

	"github.com/EcomwiseX/ecom-product/database"
	"github.com/EcomwiseX/ecom-product/dto"
	"github.com/EcomwiseX/ecom-product/models"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// FindProduct godoc
// @Summary Get a list of products
// @Description Get a list of products with pagination
// @Tags products
// @Accept  json
// @Produce  json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of products per page" default(10)
// @Success 200 {object} PaginatedResponse "List of products with pagination"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /products [get]
func Find(c *gin.Context) {
	// Paginacao
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		limit = 10
	}

	offset := (page - 1) * limit

	var products []models.Product
	if err := database.DB.Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "Unable to retrieve products",
		})
		return
	}

	var productResponses []dto.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, dto.ProductResponse{
			ID:            product.ID,
			Name:          product.Name,
			Description:   product.Description,
			Price:         product.Price,
			DiscountPrice: product.DiscountPrice,
			Stock:         product.Stock,
			Category:      product.Category,
			ImageURL:      product.ImageURL,
			IsActive:      product.IsActive,
			CreatedAt:     product.CreatedAt,
			UpdatedAt:     product.UpdatedAt,
			DeletedAt:     product.DeletedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"page":     page,
		"limit":    limit,
		"products": productResponses,
	})
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product by passing product information
// @Tags products
// @Accept  json
// @Produce  json
// @Param product body models.ProductRequest true "Product data"
// @Success 201 {object} models.ProductResponse
// @Failure 400 {object} ErrorResponse "Bad request"
// @Router /products [post]
func Create(c *gin.Context) {
	var productRequest dto.ProductRequest

	if err := c.ShouldBindJSON(&productRequest); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	if err := dto.ValidateProduct(&productRequest); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	product := models.Product{
		Name:          productRequest.Name,
		Description:   productRequest.Description,
		Price:         productRequest.Price,
		DiscountPrice: productRequest.DiscountPrice,
		Stock:         productRequest.Stock,
		Category:      productRequest.Category,
		ImageURL:      productRequest.ImageURL,
		IsActive:      productRequest.IsActive,
	}

	database.DB.Create(&product)

	resp := dto.ProductResponse{
		ID:            product.ID,
		Name:          product.Name,
		Description:   product.Description,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		Stock:         product.Stock,
		Category:      product.Category,
		ImageURL:      product.ImageURL,
		IsActive:      product.IsActive,
		CreatedAt:     product.CreatedAt,
		UpdatedAt:     product.UpdatedAt,
	}

	c.JSON(http.StatusCreated, resp)
}
