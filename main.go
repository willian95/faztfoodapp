package main

import (
	"github.com/gin-gonic/gin"
	"github.com/willian95/fastfoodapp.git/controllers"
	"github.com/willian95/fastfoodapp.git/db"
	"github.com/willian95/fastfoodapp.git/middlewares"
	"github.com/willian95/fastfoodapp.git/models"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Product{})

	router := gin.Default()

	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.POST("/forgot-password", controllers.ForgotPassword)
	router.POST("/verify-password-code", controllers.VerifyForgotPasswordCode)
	router.POST("/change-password", controllers.ChangePassword)

	protected := router.Group("/")
	protected.Use(middlewares.JwtAuthMiddleware())
	router.GET("/products/:page", controllers.GetProducts)

	router.Run()

}
