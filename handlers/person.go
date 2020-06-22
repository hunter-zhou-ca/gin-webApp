// handlers/person.go

package handlers

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"api/models"
)

// GET /users
// Get all users
func GetPeople(c *gin.Context) {
	var people []models.Person
	if err := models.DB.Find(&people).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, people)
	}
}

// GET /user/:id
// Get a particular user
func GetPerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person models.Person
	if err := models.DB.Where("id = ?", id).First(&person).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, person)
	}
}