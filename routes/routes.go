package routes

import (
	"github.com/EcomwiseX/ecom-product/controllers"
	_ "github.com/EcomwiseX/ecom-product/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/products", controllers.Find)
	r.POST("/products", controllers.Create)
	r.Run(":8081")
}
