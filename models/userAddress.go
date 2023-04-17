package model

import "gorm.io/gorm"

type UserAddress struct {
	gorm.Model
	UserID    int
	User      User      `gorm:"foreignKey:UserID"`
	AddressID int
	Address   Address   `gorm:"foreignKey:AddressID"`
}
