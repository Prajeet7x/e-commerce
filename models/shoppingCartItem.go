package model

import "gorm.io/gorm"

type ShoppingCartItem struct {
	gorm.Model
	ShoppingCart ShoppingCart 
	ShoppingCartID ShoppingCart `gorm:"foreignKey:ShoppingCartID"`
	Product Product
	ProductID Product `gorm:"foreignKey:ProductID"`
	Quantity int
}