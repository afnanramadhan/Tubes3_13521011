package main

import (
	"backend/controller"
	"backend/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	r.Use(cors.Default())

	r.GET("/api/listPertanyaan", controller.Index)
	r.GET("/api/product/*pertanyaan", controller.Show)
	r.POST("/api/radio-button", controller.ShowRadioButton)


	r.Run(":8080")
}
