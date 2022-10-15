package controller

import (
	"assignment-2/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// create order
func (idb *InDB) CreateOrder(c *gin.Context) {
	var orders model.Orders
	c.Bind(&orders)
	log.Println(orders)
	idb.DB.Create(&orders)
	c.JSON(http.StatusCreated, gin.H{
		"result": orders,
	})
	// c.Status(http.StatusCreated).JSON(orders)
}

// get orders
// update order
// delete order