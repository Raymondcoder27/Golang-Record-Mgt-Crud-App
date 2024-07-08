package main

import (
	// "fmt"

	"example/myGoProject/controllers"
	"example/myGoProject/initializers"
	// "log"

	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){
	r := gin.Default()

	r.POST("/users", controllers.UserCreate)
	r.GET("/users", controllers.UsersIndex)
	r.GET("/users/:id", controllers.ShowUser)
	r.PUT("/users/:id", controllers.UserUpdate)
	r.DELETE("/users/:id", controllers.UserDelete)
	r.Run()
}
