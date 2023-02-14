package product

import (
	"net/http"
	"shopping/models"
	"shopping/mysql"

	"github.com/gin-gonic/gin"
)

/* Get all products. */
func FindAll(c *gin.Context) {
	var products []models.Product

	if err := mysql.DB.Find(&products).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, products)
}

/* Get product by primary key. */
func FindById(c *gin.Context) {
	var product models.Product
	id := c.Params.ByName("id")

	if err := mysql.DB.First(&product, id).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, product)
}

/* Save product item. */
func Save(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	record := mysql.DB.Create(&product)
	if record.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, product)
}

/* Update product item. */
func Update(c *gin.Context) {
	var product models.Product
	id := c.Params.ByName("id")

	if err := mysql.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	record := mysql.DB.Save(&product)
	if record.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, product)
}

/* Remove product item by primary key. */
func Remove(c *gin.Context) {
	var product models.Product
	id := c.Params.ByName("id")

	record := mysql.DB.Where("id = ?", id).Delete(&product)
	if record.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, product)
}
