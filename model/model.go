package model

import (
	"log"
	"time"
	"gorm.io/gorm"
)

type Orders struct {
	Order_id      uint      `json:"orderId" gorm:"primaryKey"`
	Customer_name string    `json:"customerName"`
	Ordered_at    time.Time `json:"orderedAt"`
	Items         []Items   `gorm:"foreignKey:Item_id"`
}

type Items struct {
	Item_id     uint   `json:"itemId" gorm:"primaryKey;autoIncrement"`
	Item_code   string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	Order_id    uint   `json:"orderId" gorm:"foreignKey:Item_id"`
}

func (u *Orders) AfterCreate(tx *gorm.DB) (err error) {
	log.Println("AFTER CREATE CALLED")
	tx.Model(u.Items).Update("order_id", u.Order_id)
	return
}