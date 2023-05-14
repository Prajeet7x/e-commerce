package controllers

import (
	"ecommerce/initializers"
	model "ecommerce/models"

	"github.com/gin-gonic/gin"
)

var body struct {
	Street string
	City   string
}

func AddressCreate(c *gin.Context) {

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

func GetAllAddress(c *gin.Context) {
	var addresses []model.Address

	initializers.DB.Find(&addresses)

	c.JSON(200, gin.H{
		"All addresses": addresses,
	})
}

func GetOneAddress(c *gin.Context) {
	id := c.Param("id")

	var address []model.Address

	initializers.DB.First(&address, id)

	c.JSON(200, gin.H{
		"The address": address,
	})
}

func DeleteAddress(c *gin.Context) {
	id := c.Param("id")

	var address []model.Address

	initializers.DB.Delete(&address, id)

	c.JSON(200, gin.H{
		"Deleted address": address,
	})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")

	c.Bind(&body)

	var address []model.Address
	initializers.DB.First(&address, id)

	initializers.DB.Model(&address).Updates(model.Address{
		Street: body.Street,
		City:   body.City,
	})

	c.JSON(200, gin.H{
		"Updated address": address,
	})
}
