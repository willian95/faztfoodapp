package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/willian95/fastfoodapp.git/db"
	"github.com/willian95/fastfoodapp.git/models"
)

type ForgorPassword struct {
	Email string
}

func ForgotPassword(context *gin.Context) {
	var forgotPassword ForgorPassword
	var user models.User

	json.NewDecoder(context.Request.Body).Decode(&forgotPassword)
	result := db.DB.Where("email = ?", forgotPassword.Email).First(&user)

	if result.RowsAffected <= 0 {
		context.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "userDoesNotExist",
		})
		return
	}

	forgotPasswordCode := rand.Intn(900000) + 100000

	user.ForgotPasswordCode = strconv.Itoa(forgotPasswordCode)
	db.DB.Save(&user)

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"Message": "tokenCreated",
	})
}
