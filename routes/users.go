package routes

import (
	"events-api/models"
	"events-api/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse request", "error": err})
		return
	}
	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to register user", "error": err})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "user created", "location": fmt.Sprintf("/users/%d", user.Id)})
}
func Login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse request", "error": err})
		return
	}

	err = user.Login()
	fmt.Printf("show user2 %v", user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized", "error": err})
		return
	}
	fmt.Printf("**uid %v **", user.Id)
	token, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to generate token", "error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "user login", "token": token, "user-id": user.Id})
}
