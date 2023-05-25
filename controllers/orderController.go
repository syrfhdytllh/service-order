package controllers

import (
	"fmt"
	"net/http"
	"service-order/database"
	"service-order/models"

	"github.com/gin-gonic/gin"
)

func GetAllOrder(c *gin.Context) {
	var db = database.GetDB()

	var orders []models.Order
	err := db.Find(&orders).Error

	if err != nil {
		fmt.Println("Error getting order datas :", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

func GetOneOrders(c *gin.Context) {
	var db = database.GetDB()

	var orderOne models.Order

	err := db.First(&orderOne, "Order_id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	errs := db.First(&orderOne.Item, "U_order_id = ?", c.Param("id")).Error
	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data One": orderOne})
}

func CreateOrders(c *gin.Context) {
	var db = database.GetDB()
	// Validate input
	var input models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderinput := models.Order{Customer_name: input.Customer_name, Ordered_at: input.Ordered_at}
	db.Create(&orderinput)

	iteminput := models.Item{Item_code: input.Item[0].Item_code, Description: input.Item[0].Description, Quantity: input.Item[0].Quantity, U_order_id: orderinput.Order_id}
	db.Create(&iteminput)

	c.JSON(http.StatusOK, gin.H{
		"data":    orderinput,
		"message": "Create data success",
		"success": true})
}

func UpdateOrders(c *gin.Context) {
	var db = database.GetDB()

	var order models.Order

	err := db.First(&order, "Order_id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	errr := db.First(&order.Item, "U_order_id = ?", c.Param("id")).Error
	if errr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var item models.Item
	errs := db.First(&item, "U_order_id = ?", c.Param("id")).Error
	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&order).Updates(input)

	db.Model(&item).Updates(models.Item{Item_code: input.Item[0].Item_code, Description: input.Item[0].Description, Quantity: input.Item[0].Quantity})

	c.JSON(http.StatusOK, gin.H{
		"data":    order,
		"message": "Update data success",
		"success": true})
}

func DeleteOrders(c *gin.Context) {
	var db = database.GetDB()
	// Get model if exist
	var orderDelete models.Order
	err := db.First(&orderDelete, "Order_id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var itemDelete models.Item
	errs := db.First(&itemDelete, "U_order_id = ?", c.Param("id")).Error
	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&itemDelete)
	db.Delete(&orderDelete)

	c.JSON(http.StatusOK, gin.H{
		"messages": "Delete data Success",
		"data":     true})
}
