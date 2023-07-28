package controllers

import (
	"net/http"
	"time"

	"api-finalproject/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type orderlineInput struct {
	OrderID   uint `json:"orderID"`
	ProductID uint `json:"productID"`
	Quantity  int  `json:"quantity"`
}

// GetAllOrderline godoc
// @Summary Get all Orderline.
// @Description Get a list of Orderline.
// @Tags Orderline
// @Produce json
// @Success 200 {object} []models.Orderline
// @Router /orderline [get]
func GetAllOrderline(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var orderline []models.Orderline
	db.Find(&orderline)

	c.JSON(http.StatusOK, gin.H{"data": orderline})
}

// CreateOrderline godoc
// @Summary Create New Orderline.
// @Description Creating a new Orderline.
// @Tags Orderline
// @Param Body body orderlineInput true "the body to create a new Orderline"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Orderline
// @Router /orderline [post]
func CreateOrderline(c *gin.Context) {
	// Validate input
	var input orderlineInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create orderline
	orderline := models.Orderline{OrderID: input.OrderID, ProductID: input.ProductID, Quantity: input.Quantity}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&orderline)

	c.JSON(http.StatusOK, gin.H{"data": orderline})
}

// GetOrderlineById godoc
// @Summary Get Orderline.
// @Description Get an Orderline by id.
// @Tags Orderline
// @Produce json
// @Param id path string true "Orderline id"
// @Success 200 {object} models.Orderline
// @Router /orderline/{id} [get]
func GetOrderlineById(c *gin.Context) { // Get model if exist
	var orderline models.Orderline

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&orderline).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orderline})
}

// UpdateOrderline godoc
// @Summary Update Orderline.
// @Description Update Orderline by id.
// @Tags Orderline
// @Produce json
// @Param id path string true "Orderline id"
// @Param Body body orderlineInput true "the body to update age rating category"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Orderline
// @Router /orderline/{id} [patch]
func UpdateOrderline(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var orderline models.Orderline
	if err := db.Where("id = ?", c.Param("id")).First(&orderline).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input orderlineInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Orderline
	updatedInput.OrderID = input.OrderID
	updatedInput.ProductID = input.ProductID
	updatedInput.Quantity = input.Quantity
	updatedInput.UpdatedAt = time.Now()

	db.Model(&orderline).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": orderline})
}

// DeleteOrderline godoc
// @Summary Delete one Orderline.
// @Description Delete a Orderline by id.
// @Tags Orderline
// @Produce json
// @Param id path string true "Orderline id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /orderline/{id} [delete]
func DeleteOrderline(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var orderline models.Orderline
	if err := db.Where("id = ?", c.Param("id")).First(&orderline).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&orderline)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
