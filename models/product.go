package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductCategoryID ProductCategory `gorm:"foreignKey:ProductCategoryID"`
	ProductCategory ProductCategory
	Name              string
	Price             float32
	Description       string
	Image             string
	Stock             int
}
