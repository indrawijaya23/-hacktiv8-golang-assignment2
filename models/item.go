package models

type Item struct {
	ID          int    `gorm:"primaryKey" json:"lineItemId"`
	Code        string `gorm:"not null;type:varchar(10)" json:"itemCode"`
	Description string `gorm:"not null;type:varchar(50)" json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     int    `json:"orderId"`
}
