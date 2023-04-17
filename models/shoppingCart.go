package model

import "gorm.io/gorm"

type ShoppingCart struct {
	gorm.Model
	UserID User `gorm:"foreignKey:UserID"`
}