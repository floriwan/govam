package controllers

import (
	"log"
	"net/http"

	"github.com/floriwan/govam/models"
	"github.com/floriwan/govam/pkg/database"
	"github.com/gin-gonic/gin"
)

func RegisterUserForm(context *gin.Context) {
	//formContext := context.PostForm("registrationform")

	user := models.User{
		Email:    context.PostForm("email"),
		Name:     context.PostForm("firstname"),
		Username: context.PostForm("lastname"),
		Password: context.PostForm("password"),
	}

	log.Printf("email: %v\n", context.PostForm("email"))
	log.Printf("firstname: %v\n", context.PostForm("firstname"))

	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := database.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})

}

func RegisterUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := database.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}
