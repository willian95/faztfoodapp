package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willian95/fastfoodapp.git/db"
	"github.com/willian95/fastfoodapp.git/models"
)

func GetCategories(context *gin.Context) {

	var category []models.Category

	db.DB.Find(&category)
	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    category,
	})

}
