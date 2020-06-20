// controllers/person.go

package controllers

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"api/models"
)

// GET /users
// Get all users

func GetPeople(c *gin.Context) {
	var people []models.Person
	if err := models.DB.Find(&people).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, people)
	}
}

// GET /user/:id
// Get a particular user

func GetPerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person models.Person
	if err := models.DB.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, person)
	}
}