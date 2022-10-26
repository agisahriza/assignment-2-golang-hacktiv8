package model

type Item struct {
	Item_ID     uint   `gorm:"primaryKey"`
	ItemCode    string `gorm:"not null;type:VARCHAR(20)"`
	Description string `gorm:"type:TEXT"`
	Quantity    uint   `gorm:"not null"`
	OrderID     uint
}