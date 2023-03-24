package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willian95/fastfoodapp.git/db"
	"github.com/willian95/fastfoodapp.git/models"
	"golang.org/x/crypto/bcrypt"
)

type UserLogin struct {
	Email    string
	Password string
}

func Login(context *gin.Context) {

	var userLogin UserLogin
	var user models.User

	json.NewDecoder(context.Request.Body).Decode(&userLogin)

	result := db.DB.Where("email = ?", userLogin.Email).First(&user)

	if result.RowsAffected <= 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "USER_DOES_NOT_EXIST",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password))

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"success":  false,
			"messsage": "UNAUTHORIZED",
		})
		return
	}

	jwtString, _ := generateJWT(&user)

	context.JSON(http.StatusOK, gin.H{
		"access_token": jwtString,
	})

}
