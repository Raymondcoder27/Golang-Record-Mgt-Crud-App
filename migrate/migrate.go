package main

import (
	"example/myGoProject/initializers"
	"example/myGoProject/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){
	initializers.DB.AutoMigrate(&models.User{})
}