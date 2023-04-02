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

type ForgotPasswordStruct struct {
	Email string
}

type VerifyPasswordCodeStruct struct {
	Email string
	Code  string
}

type changePasswordStruct struct {
	Password string
	Email    string
	Code     string
}

func ForgotPassword(context *gin.Context) {
	var forgotPassword ForgotPasswordStruct
	var user models.User

	json.NewDecoder(context.Request.Body).Decode(&forgotPassword)
	result := db.DB.Where("email = ?", forgotPassword.Email).First(&user)

	if result.RowsAffected <= 0 {
		context.JSON(http.StatusNotFound, gin.H{
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

func VerifyForgotPasswordCode(context *gin.Context) {
	var verifyPasswordCode VerifyPasswordCodeStruct
	var user models.User

	json.NewDecoder(context.Request.Body).Decode(&verifyPasswordCode)
	result := db.DB.Where("email = ?", verifyPasswordCode.Email).First(&user)

	if result.RowsAffected <= 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "userDoesNotExist",
		})
		return
	}

	if verifyPasswordCode.Code == user.ForgotPasswordCode {

		context.JSON(http.StatusOK, gin.H{
			"success": true,
		})
		return
	}

	context.JSON(http.StatusNotFound, gin.H{
		"success": false,
		"message": "tokenNotMatch",
	})

}

func ChangePassword(context *gin.Context) {
	var changePassword changePasswordStruct
	var user models.User

	json.NewDecoder(context.Request.Body).Decode(&changePassword)
	result := db.DB.Where("email = ?", changePassword.Email).Where("forgot_password_code = ?", changePassword.Code).First(&user)

	if result.RowsAffected <= 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "userDoesNotExist",
		})
		return
	}

	hash, err := MakePassword(changePassword.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "somethingWentWrong",
		})
		return
	}

	user.ForgotPasswordCode = ""
	user.Password = hash
	db.DB.Save(&user)

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "passwordChanged",
	})

}
