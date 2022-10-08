package models

import "time"

type Order struct {
	ID           int       `gorm:"primaryKey" json:"orderId"`
	CustomerName string    `gorm:"not null;type:varchar(50)" json:"customerName"`
	OrdererdAt   time.Time `json:"orderedAt"`
	Items        []Item    `json:"items"`
}
