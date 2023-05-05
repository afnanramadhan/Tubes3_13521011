package main

import (
	"backend/controller"
	"backend/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	r.Use(cors.Default())

	r.GET("/api/listPertanyaan", controller.Index)
	r.GET("/api/product/*pertanyaan", controller.Show)
	r.POST("/api/product", controller.Create)
	r.PUT("/api/product/:id", controller.Update)
	r.DELETE("/api/product", controller.Delete)

	r.Run(":8080")
}