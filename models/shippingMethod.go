package model

import "gorm.io/gorm"

type ShippingMethod struct {
	gorm.Model
	MethodName string	
	Price float32
}
