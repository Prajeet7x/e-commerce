package model

import "gorm.io/gorm"

type OrderStatus struct {
	gorm.Model
	Status string
}
