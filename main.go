package main

import (
	"github.com/EcomwiseX/ecom-product/database"
	"github.com/EcomwiseX/ecom-product/routes"
)

func main() {
	database.ConnectDb()
	routes.HandleRequests()
}
