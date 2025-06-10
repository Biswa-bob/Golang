package routes

import (
	"fmt"
	"net/http"

	"example.com/event-booking/models"
	"example.com/event-booking/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldnot create user , Try again later",
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "User Created Successfully",
	})
}
func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse requested data",
		})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		fmt.Println(">>>>err", err)
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": err,
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldnot authenticate user",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "User Logged in Succesfully",
		"token":   token,
	})
}
