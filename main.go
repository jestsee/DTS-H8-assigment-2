package main

import (
	"assignment-2/config"
	"assignment-2/controller"
	"github.com/gin-gonic/gin"
	"log"

	_ "assignment-2/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Orders API
// @version 1.0
// @description DTS H8 Assigment 2

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @termsOfService http://swagger.io/terms
// @host localhost:3000
// @BasePath /
func main() {
	conf, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	db := config.ConnectDB(&conf)
	inDB := &controller.InDB{DB: db}
	router := gin.Default()

	// docs route
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/orders", inDB.GetOrders)
	router.POST("/orders", inDB.CreateOrder)
	router.PUT("/orders/:orderId", inDB.UpdateOrder)
	router.DELETE("/orders/:orderId", inDB.DeleteOrder)
	router.Run(":3000")
}
