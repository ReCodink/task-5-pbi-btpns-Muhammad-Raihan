package controllers

import (
	"net/http"

	"github.com/ReCodink/task-5-pbi-btpns-Muhammad-Raihan/app"
	db "github.com/ReCodink/task-5-pbi-btpns-Muhammad-Raihan/database"
	"github.com/ReCodink/task-5-pbi-btpns-Muhammad-Raihan/models"
	"github.com/gin-gonic/gin"
)

func UpdateUser(ctx *gin.Context) {
	userID := ctx.Param("id")

	db := db.Init()

	var foundUser models.User
	if err := db.First(&foundUser, userID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "User Not Found",
		})
		return
	}

	var updateUser app.UpdateUser
	if err := ctx.ShouldBindJSON(&updateUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Bad Request",
		})
		return
	}

	if len(updateUser.Password) < 6 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Password must be at least 6 characters",
		})
		return
	}

	foundUser.Username = updateUser.Username
	foundUser.Email = updateUser.Email
	foundUser.Password = updateUser.Password

	if err := db.Save(&foundUser).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to update user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User Update Successfully",
	})
}

func DeleteUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	db := db.Init()

	var foundUser models.User
	if err := db.First(&foundUser, userID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "User Not Found",
		})
		return
	}

	if err := db.Where("id = ?", userID).Unscoped().Delete(&foundUser).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to delete user",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User Deleted Successfully",
	})
}
