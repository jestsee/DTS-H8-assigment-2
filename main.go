package main

import (
	"log"
	"assignment-2/config"
	"assignment-2/controller"
	"github.com/gin-gonic/gin"
)

func main () {
	conf, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	db := config.ConnectDB(&conf)
	inDB := &controller.InDB{DB: db}
	router := gin.Default()

	router.POST("/orders", inDB.CreateOrder)
	router.Run(":3000")
}