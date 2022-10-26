package model

import "time"

type Order struct {
	Order_ID     uint   	`gorm:"primaryKey"`
	CustomerName string 	`gorm:"not null;type:varchar(250)"`
	OrderedAt 	 time.Time  `gorm:"not null;type:timestamp;autoCreateTime"`
	Items 		 []Item
}