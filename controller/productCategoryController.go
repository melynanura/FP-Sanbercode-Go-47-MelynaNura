package controllers

import (
	"net/http"
	"time"

	"api-finalproject/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productCategoryInput struct {
	Category string `json:"category"`
}

// GetAllCategory godoc
// @Summary Get all ProductCategory.
// @Description Get a list of ProductCategory.
// @Tags ProductCategory
// @Produce json
// @Success 200 {object} []models.ProductCategory
// @Router /product-categories [get]
func GetAllCategory(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var productCategory []models.ProductCategory
	db.Find(&productCategory)

	c.JSON(http.StatusOK, gin.H{"data": productCategory})
}

// CreateCategory godoc
// @Summary Create New ProductCategory.
// @Description Creating a new ProductCategory.
// @Tags ProductCategory
// @Param Body body productCategoryInput true "the body to create a new ProductCategory"
// @Produce json
// @Success 200 {object} models.ProductCategory
// @Router /product-categories [post]
func CreateCategory(c *gin.Context) {
	// Validate input
	var input productCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Category
	productCategory := models.ProductCategory{Category: input.Category}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&productCategory)

	c.JSON(http.StatusOK, gin.H{"data": productCategory})
}

// GetCategoryById godoc
// @Summary Get ProductCategory.
// @Description Get an ProductCategory by id.
// @Tags ProductCategory
// @Produce json
// @Param id path string true "ProductCategory id"
// @Success 200 {object} models.ProductCategory
// @Router /product-categories/{id} [get]
func GetCategoryById(c *gin.Context) { // Get model if exist
	var Category models.ProductCategory

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&Category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Category})
}

// GetProductByCategoryId godoc
// @Summary Get Product.
// @Description Get all Product by ProductCategoryId.
// @Tags ProductCategory
// @Produce json
// @Param id path string true "ProductCategory id"
// @Success 200 {object} []models.Products
// @Router /product-categories/{id}/product [get]
func GetProductByCategoryId(c *gin.Context) { // Get model if exist
	var product []models.Products

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("product_category_id = ?", c.Param("id")).Find(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// UpdateCategory godoc
// @Summary Update ProductCategory.
// @Description Update ProductCategory by id.
// @Tags ProductCategory
// @Produce json
// @Param id path string true "ProductCategory id"
// @Param Body body productCategoryInput true "the body to update product category"
// @Success 200 {object} models.ProductCategory
// @Router /product-categories/{id} [patch]
func UpdateCategory(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var category models.ProductCategory
	if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input productCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.ProductCategory
	updatedInput.Category = input.Category
	updatedInput.UpdatedAt = time.Now()

	db.Model(&category).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// DeleteCategory godoc
// @Summary Delete one ProductCategory.
// @Description Delete a ProductCategory by id.
// @Tags ProductCategory
// @Produce json
// @Param id path string true "ProductCategory id"
// @Success 200 {object} map[string]boolean
// @Router /product-categories/{id} [delete]
func DeleteCategory(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var category models.ProductCategory
	if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&category)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
