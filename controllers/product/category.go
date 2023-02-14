package product

import (
	"net/http"
	"shopping/models"
	"shopping/mysql"

	"github.com/gin-gonic/gin"
)

/* Find all categories. */
func CategoryFindAll(c *gin.Context) {
	var categories []models.Category

	if err := mysql.DB.Find(&categories).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, categories)
}

/* Get category by primary key. */
func CategoryFindById(c *gin.Context) {
	var category models.Category
	id := c.Params.ByName("id")

	if err := mysql.DB.First(&category, id).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, category)
}

/* Save category item. */
func CategorySave(c *gin.Context) {
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	record := mysql.DB.Create(&category)
	if record.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, category)
}

/* Update category item. */
func CategoryUpdate(c *gin.Context) {
	var category models.Category
	id := c.Params.ByName("id")

	if err := mysql.DB.First(&category, id).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if err := c.ShouldBindJSON(&category); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	record := mysql.DB.Save(&category)
	if record.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, category)
}

/* Remove category item by primary key. */
func CategoryRemove(c *gin.Context) {
	var category models.Category
	id := c.Params.ByName("id")

	record := mysql.DB.Where("id = ?", id).Delete(&category)
	if record.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, category)
}
