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
	// r.POST("/api/product", controller.Create)
	r.POST("/api/radio-button", controller.ShowRadioButton)
	// r.PUT("/api/product/:id", controller.Update)
	// r.DELETE("/api/product", controller.Delete)

	// r.POST("api/radio-button", func(c *gin.Context) {

	// 	type RadioValue struct {
	// 		Value bool `json:"value"`
	// 	}
	// 	var radioValue RadioValue
	// 	if err := c.ShouldBind(&radioValue); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	fmt.Println(radioValue)
	// 	value := radioValue.Value
	// 	fmt.Println("Radio value: ")
	// 	fmt.Println(value)
	// 	fmt.Println("============")
	// 	c.JSON(http.StatusOK, gin.H{"Value": value})
	// })

	r.Run(":8080")
}
