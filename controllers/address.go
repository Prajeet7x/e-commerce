package controllers

import (
	"ecommerce/initializers"
	model "ecommerce/models"

	"github.com/gin-gonic/gin"
)

func AddressCreate(c *gin.Context) {
	var body struct {
		Street string
		City   string
	}
	c.Bind(&body)

	address := model.Address{Street: body.Street, City: body.City}

	result := initializers.DB.Create(&address)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"address": address,
	})

}