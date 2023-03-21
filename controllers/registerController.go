package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willian95/fastfoodapp.git/db"
	"github.com/willian95/fastfoodapp.git/models"
)

func Register(context *gin.Context) {

	var user models.User
	var userFind models.User
	json.NewDecoder(context.Request.Body).Decode(&user)

	result := db.DB.Where("email = ?", user.Email).First(&userFind)

	if result.RowsAffected > 0 {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"success": false,
			"message": "EMAIL_ALREADY_EXISTS",
		})
		return
	}

	createdUser := db.DB.Create(&user)
	err := createdUser.Error

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "SOMETHING_WENT_WRONG",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "USER_CREATED",
	})
	return

}
