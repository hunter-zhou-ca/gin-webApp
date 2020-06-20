package main

import(
	//Package gin implements a HTTP web framework called gin
	"github.com/gin-gonic/gin"
	"api/models"
	"api/controllers"
	"api/middleware"
)

func main(){
	r := gin.Default()
	models.ConnectDataBase()

	//create a jwt middlware
	jwt := middleware.JwtMiddleware()

	//use localhost:8080/login to get token
	//username: admin
	//password: password
	r.POST("/login", jwt.LoginHandler)

	//check token
	auth := r.Group("/", jwt.MiddlewareFunc())

	auth.GET("/users", controllers.GetPeople)
	auth.GET("/user/:id", controllers.GetPerson)

	r.Run(":8080")
}