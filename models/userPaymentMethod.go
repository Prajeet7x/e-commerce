package model

import "gorm.io/gorm"

type UserPaymentMethod struct {
	gorm.Model
	User User
	UserID User `gorm:"foreignKey:UserID"`
	Provider      string
	AccountNumber string
}