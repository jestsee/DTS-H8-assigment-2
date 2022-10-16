package controller

import (
	"assignment-2/model"
	"github.com/gin-gonic/gin"
	// "log"
	"net/http"
)

// CreateOrder godoc
// @Summary Create new order
// @Description Create new order
// @Tag orders
// @Produce json
// @Param order body model.Orders true "Create order"
// @Success 201 {object} model.Orders
// @Router /orders [post]
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

// GetOrders godoc
// @Summary Get details of all orders
// @Description Get details of all orders
// @Tag orders
// @Produce json
// @Success 200 {array} model.Orders
// @Router /orders [get]
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

// UpdateOrder godoc
// @Summary Update existing order
// @Description Update existing order
// @Tag orders
// @Produce json
// @Param id path int true "order id" 
// @Param order body model.Orders true "order body"
// @Success 200 {object} model.Orders
// @Router /orders/{orderId} [put]
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

// DeleteOrder godoc
// @Summary Delete existing order
// @Description Delete existing order
// @Tag orders
// @Produce json
// @Param id path string true "todo ID"
// @Success 200 {object} string
// @Router /orders/{orderId} [delete]
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
