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

	router := gin.Default()

	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.POST("/forgot-password", controllers.ForgotPassword)

	protected := router.Group("/")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/home", controllers.Home)

	router.Run()

}
