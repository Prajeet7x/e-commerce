package main

import (
	"ecommerce/initializers"
	model "ecommerce/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDataBase()
}

func main() {
	initializers.DB.AutoMigrate(
		&model.User{},
		&model.Address{},
		&model.OrderStatus{},
		&model.ProductCategory{},
		&model.Product{},
		&model.ShippingMethod{},
		&model.UserPaymentMethod{},
		&model.UserAddress{},
		&model.Order{},
		&model.ShoppingCart{},
		&model.ShoppingCartItem{},
	)
}
