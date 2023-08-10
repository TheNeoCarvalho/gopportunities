package handler

import (
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(cxt *gin.Context){
	cxt.JSON(200, gin.H{
		"message": "hello GET",
	})
}