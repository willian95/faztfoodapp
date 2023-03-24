package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willian95/fastfoodapp.git/db"
	"github.com/willian95/fastfoodapp.git/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(context *gin.Context) {

	var user models.User
	var userFind models.User
	json.NewDecoder(context.Request.Body).Decode(&user)

	result := db.DB.Where("email = ?", user.Email).First(&userFind)

	if result.RowsAffected > 0 {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"success": false,
			"message": "emailAlreadyExist",
		})
		return
	}

	createdUser := db.DB.Create(&user)
	err := createdUser.Error

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "somethingWentWrong",
		})
		return
	}

	hash, err := MakePassword(user.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "somethingWentWrong",
		})
		return
	}

	user.Password = hash
	db.DB.Save(&user)

	jwtString, _ := generateJWT(&user)
	context.JSON(http.StatusOK, gin.H{
		"success":     true,
		"message":     "userCreated",
		"accessToken": jwtString,
	})

}

func MakePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
