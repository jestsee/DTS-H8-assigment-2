package controller

import (
	"assignment-2/model"
	"github.com/gin-gonic/gin"
	// "log"
	"net/http"
)

// Create order
func (idb *InDB) CreateOrder(c *gin.Context) {
	var orders model.Orders
	c.Bind(&orders)
	err := idb.DB.Create(&orders).Error
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"result": orders,
	})
}

// Get orders
func (idb *InDB) GetOrders(c *gin.Context) {
	var (
		orders []model.Orders
		result gin.H
	)

	err := idb.DB.Find(&orders).Error
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	err = idb.DB.Preload("Items").Find(&orders).Error
	// log.Println(orders)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	if len(orders) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": orders,
			"count":  len(orders),
		}
	}

	c.JSON(http.StatusOK, result)
}

// Update order
func (idb *InDB) UpdateOrder(c *gin.Context) {
	var orders model.Orders

	id := c.Param("orderId")

	err := idb.DB.First(&orders, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	
	c.Bind(&orders)

	err = idb.DB.Save(&orders).Error
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": orders,
	})
}

// Delete order
func (idb *InDB) DeleteOrder(c *gin.Context) {
	var orders model.Orders
	id := c.Param("orderId")

	err := idb.DB.First(&orders, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = idb.DB.Delete(&orders).Error
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Data deleted succesfully",
	})
}
