package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	//fk user_id, payment_method_id, shipping_address, order_status
	UserID int
	User User `gorm:"foreignKey:UserID"`
	OrderDate time.Time
	OrderTotal float32
}