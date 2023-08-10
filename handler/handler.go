package handler

import (
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(cxt *gin.Context){
	cxt.JSON(200, gin.H{
		"message": "hello GET",
	})
}

func ShowOpeningHandler(cxt *gin.Context){
	cxt.JSON(200, gin.H{
		"message": "hello GET",
	})
}

func UpdateOpeningHandler(cxt *gin.Context){
	cxt.JSON(200, gin.H{
		"message": "hello GET",
	})
}

func DeleteOpeningHandler(cxt *gin.Context){
	cxt.JSON(200, gin.H{
		"message": "hello GET",
	})
}

func ListOpeningsHandler(cxt *gin.Context){
	cxt.JSON(200, gin.H{
		"message": "hello GET",
	})
}