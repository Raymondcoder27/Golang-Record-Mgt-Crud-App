package controllers

import (
	"example/myGoProject/initializers"
	"example/myGoProject/models"

	"github.com/gin-gonic/gin"
)

func UserCreate(c *gin.Context) {
	//get the data off the req body

	//bind it to the model struct
	var body struct {
		Name string
		Age int
	}
	c.Bind(&body)

	// var user models.User
	user := models.User{Name: body.Name, Age: body.Age,}
	initializers.DB.Create(&user)

	//respond with the user
	c.JSON(200, gin.H{
		"user": user,
	})
}

//get all users
func UsersIndex(c *gin.Context){
	//find all the users
	var users []models.User
	initializers.DB.Find(&users)

	//respond with them
	c.JSON(200, gin.H{
		"users": users,
	})
}

//find a user by id
func ShowUser(c *gin.Context){
	//get the id param
	id := c.Param("id")

	//retrieve the user
	var user models.User
	initializers.DB.First(&user, id)

	//return with the user
	c.JSON(200, gin.H{
		"user": user,
	})
}

//update a user
func UserUpdate(c *gin.Context){
	//get the id param
	id := c.Param("id")

	//get the new user off req body
	var body struct{
		Name string
		Age int
	}
	c.Bind(&body)

	//find the user
	var user models.User
	initializers.DB.First(&user, id)

	//update the user
	initializers.DB.Model(&user).Updates(models.User{
		Name: body.Name, 
		Age: body.Age,
	})

	//respond with the user
	c.JSON(200, gin.H{
		"user": user,
	})
}


//deleting a user
func UserDelete(c *gin.Context){
	//get the id off param
	id := c.Param("id")

	//delete the user
	// var user models.User
	// initializers.DB.Delete(&user, id)
	initializers.DB.Unscoped().Delete(&models.User{}, id)

	//respond
	c.Status(200)
}