package model

import "gorm.io/gorm"

type OrderStatus struct {
	gorm.Model
	Status string `gorm:"type:enum(Canceled, Processing, Delivered);column:order_status"`
}
