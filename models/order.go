package models

type Order struct {
	Order_id      int64  `gorm:"primaryKey" json:"order_id"`
	Customer_name string `gorm:"varchar(100)" json:"customer_name"`
	Ordered_at    string `gorm:"varchar(100)" json:"ordered_at"`
	Item          []Item `gorm:"foreignKey:U_order_id" json:"item"`
}

type Item struct {
	Item_id     int64  `gorm:"primaryKey" json:"item_id"`
	Item_code   string `gorm:"varchar(100)" json:"item_code"`
	Description string `gorm:"varchar(100)" json:"description"`
	Quantity    int64  `gorm:"integer(5)" json:"quantity"`
	U_order_id  int64  `gorm:"foreignKey:Order_id" json:"u_order_id"`
}
