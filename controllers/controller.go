package controllers

import (
	"strconv"

	"github.com/EcomwiseX/ecom-product/database"
	"github.com/EcomwiseX/ecom-product/models"
	"github.com/gin-gonic/gin"
)

func FindProduct(c *gin.Context) {
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
		c.JSON(500, gin.H{"error": "Unable to retrieve products"})
		return
	}

	c.JSON(200, gin.H{
		"page":     page,
		"limit":    limit,
		"products": products,
	})
}
