package model

type Orders struct {
	Order_id      uint    `json:"orderId" gorm:"primaryKey"`
	Customer_name string  `json:"customerName"`
	Ordered_at    string  `json:"orderedAt"`
	Items         []Items `gorm:"foreignKey:Item_id"`
}

type Items struct {
	Item_id     uint   `json:"itemId" gorm:"primaryKey"`
	Item_code   string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	Order_id    uint   `json:"orderId"`
	// Orders      Orders `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:Order_id"`
}

// satu orders ada many items
// satu items belongs to orders
