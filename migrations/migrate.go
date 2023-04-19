package main

import (
	"ecommerce/initializers"
	model "ecommerce/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.InitializeDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&model.Address{}, &model.Order{}, &model.OrderStatus{}, &model.Product{}, &model.ProductCategory{}, &model.ShippingMethod{}, &model.User{}, &model.UserPaymentMethod{}, &model.UserAddress{}, &model.ShoppingCart{}, &model.ShoppingCartItem{},)
}
