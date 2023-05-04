package main

import (
	"backend/controller"
	"backend/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", controller.Index)
	r.GET("/api/product/:id", controller.Show)
	r.POST("/api/product", controller.Create)
	r.PUT("/api/product/:id", controller.Update)
	r.DELETE("/api/product", controller.Delete)

	r.GET("/form", func(c *gin.Context) {
		c.File("../frontend/tes.html")
	   })

	r.Run(":8000")
}
