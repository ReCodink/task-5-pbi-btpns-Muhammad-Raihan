package controllers

import (
	"net/http"
	"strconv"

	"github.com/ReCodink/task-5-pbi-btpns-Muhammad-Raihan/app"
	db "github.com/ReCodink/task-5-pbi-btpns-Muhammad-Raihan/database"
	"github.com/ReCodink/task-5-pbi-btpns-Muhammad-Raihan/helpers"
	"github.com/ReCodink/task-5-pbi-btpns-Muhammad-Raihan/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IsEmailExist(email string) bool {
	var user models.User

	db := db.Init()
	result := db.Where("email = ?", email).First(&user)

	if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
		return false
	}

	return true
}

func IsUsernameExist(username string) bool {
	var user models.User
	db := db.Init()
	result := db.Where("username = ?", username).First(&user)

	if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func LoginUser(ctx *gin.Context) {
	var input app.LoginUser
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := db.Init()
	var user models.User
	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid Credentials",
			"success": false,
		})
		return
	}

	if err := user.ComparePassword(input.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid Credentials",
			"success": false,
		})
	}

	userIDString := strconv.Itoa(int(user.ID))
	token, err := helpers.GenerateToken(userIDString)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token "})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token":   token,
		"success": true,
		"message": "User Login Successfully",
	})
}

func RegisterUser(ctx *gin.Context) {
	var input app.RegisterUser

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(input.Password) < 6 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Password must be at least 6 characters",
		})

		return
	}

	if IsEmailExist(input.Email) {
		ctx.JSON(http.StatusConflict, gin.H{
			"success": false,
			"message": "Username already exist, must be unique",
		})
		return
	}

	register := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	db := db.Init()
	if err := db.Create(&register).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "User Registered Successfully",
	})
}
