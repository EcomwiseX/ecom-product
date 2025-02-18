package routes

import (
	"github.com/EcomwiseX/ecom-product/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/products", controllers.FindProduct)
	r.Run(":8081")
}
