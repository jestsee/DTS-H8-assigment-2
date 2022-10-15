package controller

import (
	"assignment-2/model"
	"github.com/gin-gonic/gin"
	// "log"
	"net/http"
)

// create order
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

// get orders
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

// update order
// delete order
func (idb *InDB) DeleteOrder(c *gin.Context) {
	var (
		orders model.Orders
		result gin.H
	)
	id := c.Param("orderId")
	err := idb.DB.First(&orders, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err = idb.DB.Delete(&orders).Error
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": "Data deleted succesfully",
		}
	}
	c.JSON(http.StatusOK, result)
}
// TODO violates FK