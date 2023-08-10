package handler

import (
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(cxt *gin.Context){
	cxt.JSON(200, gin.H{
		"message": "hello GET",
	})
}
