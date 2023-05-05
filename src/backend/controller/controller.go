package controller

import (
	"backend/lib"
	"encoding/json"
	"fmt"
	"net/http"

	"backend/models"
	// "gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var products []models.Data

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"listPertanyaan": products})

}

func Show(c *gin.Context) {
	// var product models.Data
	id := c.Param("pertanyaan")

	fmt.Println(id)
	product := lib.Utama(id)

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Create(c *gin.Context) {

	var product models.Data

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Update(c *gin.Context) {
	var product models.Data
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})

}

func Delete(c *gin.Context) {

	var product models.Data

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}