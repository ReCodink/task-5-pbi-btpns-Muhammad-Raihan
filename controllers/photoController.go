package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ReCodink/task-5-pbi-btpns-Muhammad-Raihan/app"
	db "github.com/ReCodink/task-5-pbi-btpns-Muhammad-Raihan/database"
	"github.com/ReCodink/task-5-pbi-btpns-Muhammad-Raihan/models"
	"github.com/gin-gonic/gin"
)

func UploadPhotoProfile(ctx *gin.Context) {
	var req app.UploadPhoto
	userID, _ := ctx.Get("userID") // Ganti dari ctx.Get("id") menjadi ctx.Get("userID")

	file, err := ctx.FormFile("PhotoUrl")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing photo file"})
		return
	}

	timestamp := time.Now().UnixNano()
	filePhoto := fmt.Sprintf("%d-%s", timestamp, strings.ReplaceAll(file.Filename, " ", ""))
	filename := "http://localhost:8080/api/photos/" + filePhoto // Perbaiki URL sesuai dengan format yang benar

	req.Title = ctx.PostForm("title")
	req.Caption = ctx.PostForm("caption")
	req.PhotoUrl = filename

	photo := models.Photo{
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoUrl: req.PhotoUrl,
		UserID:   userID.(uint),
	}

	db := db.Init()

	err = ctx.SaveUploadedFile(file, "./uploads/"+filePhoto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	if err := db.Create(&photo).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Photo Uploaded Successfully",
		"success": true,
	})
}

// Sisanya tidak perlu diubah karena sudah cocok dengan model Photo yang diperbarui.

func GetPhotoProfile(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	db := db.Init()
	var photos []models.Photo

	result := db.Where("user_id = ?", userID).Preload("User").Find(&photos)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to fetch photos",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Successfully Retrieve Data",
		"data":    photos,
	})

}

func UpdatePhotoProfile(ctx *gin.Context) {
	photoID := ctx.Param("id")
	userID, _ := ctx.Get("userID")

	file, err := ctx.FormFile("photo_file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing photo file"})
		return
	}

	timestamp := time.Now().UnixNano()
	filePhoto := fmt.Sprintf("%d-%s", timestamp, strings.ReplaceAll(file.Filename, " ", ""))
	filename := "http://localhost:8080/api/photos/" + filePhoto

	updateData := app.UploadPhoto{
		Title:    ctx.PostForm("title"),
		Caption:  ctx.PostForm("caption"),
		PhotoUrl: filename,
	}

	photo := models.Photo{
		Title:    updateData.Title,
		Caption:  updateData.Caption,
		PhotoUrl: updateData.PhotoUrl,
	}

	db := db.Init()
	result := db.Where("id = ?", photoID).Where("user_id", userID).Preload("User").First(&photo)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Photo Not Found",
			"success": false,
		})
		return
	}

	if err := db.Save(&photo).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Photo Update Sucessfully",
		"data":    photo,
	})
}

func DeletePhotoProfile(ctx *gin.Context) {
	photoID := ctx.Param("id")
	userID, _ := ctx.Get("userID")

	var photo models.Photo

	db := db.Init()

	if err := db.Where("id = ?", photoID).Where("user_id = ?", userID).First(&photo).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Photo Not Found",
		})
		return
	}

	if err := db.Where("id = ?", photoID).Where("user_id = ?", userID).Unscoped().Delete(&photo).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to delete photo",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Photo Deleted Successfully",
	})
}
