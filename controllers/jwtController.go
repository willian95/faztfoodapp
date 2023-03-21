package controllers

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/willian95/fastfoodapp.git/models"
)

func generateJWT(user *models.User) (string, error) {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"name":  user.Name,
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
