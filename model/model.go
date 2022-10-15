package model

import (
	"log"
	"time"

	// "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Orders struct {
	Order_id      uint      `json:"orderId" gorm:"primaryKey"`
	Customer_name string    `json:"customerName"`
	Ordered_at    time.Time `json:"orderedAt"`
	MemberNumber  string
	Items         []Items `gorm:"foreignKey:Order_id"`
}

type Items struct {
	Item_id     uint   `json:"itemId" gorm:"primaryKey;autoIncrement"`
	Item_code   string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	Order_id    uint   `json:"orderId"`
	// perlu foreign key ga ya?
}

// func (o *Orders) AfterCreate(tx *gorm.DB) (err error) {
// 	log.Println("AFTER CREATE CALLED")
// 	tx.Model(o.Items).Update("order_id", o.Order_id)
// 	return
// }

func (o *Orders) BeforeDelete(tx *gorm.DB) (err error) {
	log.Println("BEFORE DELETE CALLED")
	tx.Model(o.Items).Where("order_id = ?", o.Order_id).Delete(o.Items)
	return 
}
