package handler

import (
	"net/http"

	"github.com/TheNeoCarvalho/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context){
	openings := []schemas.Opening{}
	
	if err := db.Find(&openings).Error; err != nil{
			sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-opening", openings)

}