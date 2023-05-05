package controller

import (
	"backend/lib"
	"fmt"
	"net/http"

	"backend/models"
	// "gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

var IsVal bool

func Index(c *gin.Context) {

	var products []models.Data

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"listPertanyaan": products})

}

func IndexHistory(c *gin.Context) {

	var products []models.History

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"histories": products})

}

func Show(c *gin.Context) {
	// var product models.Data
	id := c.Param("pertanyaan")

	fmt.Println(id)
	product := lib.Utama(id, IsVal)

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func ShowHistory(c *gin.Context) {
	id := c.Param("page_history")

	fmt.Println(id)
	product := lib.Utama(id, IsVal)

	c.JSON(http.StatusOK, gin.H{"product": product})
}


func ShowRadioButton(c *gin.Context) {
	type RadioValue struct {
		Value bool `json:"value"`
	}
	var radioValue RadioValue
	if err := c.ShouldBind(&radioValue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	IsVal = radioValue.Value
	c.JSON(http.StatusOK, gin.H{"Value": radioValue.Value})

}
