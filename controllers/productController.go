package controllers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/willian95/fastfoodapp.git/db"
	"github.com/willian95/fastfoodapp.git/helpers"
	"github.com/willian95/fastfoodapp.git/models"
)

func GetProducts(context *gin.Context) {

	var product []models.Product
	var totalRows int64
	var totalPages int64

	const perPage = 10
	pageStr := context.Param("page")
	page, _ := strconv.Atoi(pageStr)
	offset := helpers.CalculateOffset(page, perPage)

	db.DB.Model(&models.Product{}).Count(&totalRows)
	db.DB.Limit(perPage).Offset(offset).Find(&product)
	totalPages = int64(math.Ceil(float64(totalRows) / perPage))

	context.JSON(http.StatusOK, gin.H{
		"success":    true,
		"data":       product,
		"page":       page,
		"perPage":    perPage,
		"totalPages": totalPages,
	})

}
