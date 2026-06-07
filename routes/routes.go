package routes

import (
	"github.com/gin-gonic/gin"
	"go-gin-postgres-crud/controllers"

)

func SetupRoutes(router *gin.Engine) {

	api := router.Group("/api")

	{
		api.POST("/products", controllers.CreateProduct)
	}
	{
		api.GET("/products", controllers.GetProducts)
	}
	{
		api.GET("/products/:id", controllers.GetProduct)
	}
	{
		api.PUT("/products/:id", controllers.UpdateProduct)
	}
	{
		api.DELETE("/products/:id", controllers.DeleteProduct)
	}

}
