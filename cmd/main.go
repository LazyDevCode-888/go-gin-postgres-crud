package main

import (
	"go-gin-postgres-crud/config"
	"go-gin-postgres-crud/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDatabase()

	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":8080")

}
