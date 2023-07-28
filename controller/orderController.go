package controllers

import (
	"net/http"
	"time"

	"api-finalproject/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type orderInput struct {
	Address    string `json:"address"`
	CustomerID uint   `json:"customerID"`
}

// GetAllOrder godoc
// @Summary Get all Order.
// @Description Get a list of Order.
// @Tags Order
// @Produce json
// @Success 200 {object} []models.Order
// @Router /order [get]
func GetAllOrder(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var order []models.Order
	db.Find(&order)

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// CreateOrder godoc
// @Summary Create New Order.
// @Description Creating a new Order.
// @Tags Order
// @Param Body body orderInput true "the body to create a new Order"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Order
// @Router /order [post]
func CreateOrder(c *gin.Context) {
	// Validate input
	var input orderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Category
	order := models.Order{Address: input.Address, CustomerID: input.CustomerID, OrderDate: time.Now()}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&order)

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// GetOrderById godoc
// @Summary Get Order.
// @Description Get an Order by id.
// @Tags Order
// @Produce json
// @Param id path string true "Order id"
// @Success 200 {object} models.Order
// @Router /order/{id} [get]
func GetOrderById(c *gin.Context) { // Get model if exist
	var order models.Order

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// GetOrderlineByOrderId godoc
// @Summary Get Orderline.
// @Description Get all Orderline by OrderId.
// @Tags Order
// @Produce json
// @Param id path string true "Order id"
// @Success 200 {object} []models.Orderline
// @Router /order/{id}/orderline [get]
func GetOrderlineByOderId(c *gin.Context) { // Get model if exist
	var orderline []models.Orderline

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("order_id = ?", c.Param("id")).Find(&orderline).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orderline})
}

// UpdateOrder godoc
// @Summary Update Order.
// @Description Update Order by id.
// @Tags Order
// @Produce json
// @Param id path string true "Order id"
// @Param Body body orderInput true "the body to update order"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Order
// @Router /order/{id} [patch]
func UpdateOrder(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var order models.Order
	if err := db.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input orderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Order
	updatedInput.OrderDate = time.Now()
	updatedInput.Address = input.Address
	updatedInput.CustomerID = input.CustomerID
	updatedInput.UpdatedAt = time.Now()

	db.Model(&order).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// DeleteOrder godoc
// @Summary Delete one Order.
// @Description Delete a Order by id.
// @Tags Order
// @Produce json
// @Param id path string true "Order id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /order/{id} [delete]
func DeleteOrder(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var order models.Order
	if err := db.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&order)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
