package config

import (
	"assignment-2/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectDB(config *Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("? Connected Successfully to the Database")

	DB.Migrator().CreateConstraint(&model.Orders{}, "Items")
	DB.Migrator().CreateConstraint(&model.Orders{}, "fk_orders_items")
	// DB.Migrator().CreateConstraint(&model.Items{}, "Orders")
	// DB.Migrator().CreateConstraint(&model.Items{}, "fk_items_orders")
	DB.AutoMigrate(&model.Orders{}, &model.Items{})
	if err != nil {
		log.Println(err)
	}
	return DB
}
