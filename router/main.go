package router

import (
	"github.com/ReCodink/task-5-pbi-btpns-Muhammad-Raihan/controllers"
	"github.com/ReCodink/task-5-pbi-btpns-Muhammad-Raihan/middlewares"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	appAuth := router.Group("/api/auth")
	{
		appAuth.POST("/register", controllers.RegisterUser)
		appAuth.POST("/login", controllers.LoginUser)
	}

	appMain := router.Group("/api")
	{
		appMain.Use(middlewares.AuthMiddleware())
		appMain.PUT("/users/:id", controllers.UpdateUser)
		appMain.DELETE("/users/:id", controllers.DeleteUser)

		// Perbaikan pada endpoint /photos
		appMain.GET("/photos", controllers.GetPhotoProfile) // Tambahkan penanganan endpoint GET /photos
		appMain.POST("/photos", controllers.UploadPhotoProfile)
		appMain.PUT("/photos/:id", controllers.UpdatePhotoProfile)
		appMain.DELETE("/photos/:id", controllers.DeletePhotoProfile)
	}
	return router
}
