package handler

import (
	"net/http"

	"github.com/TheNeoCarvalho/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

//BasePath /api/v1

//@Sumary List opening
//@Description List a Job opening
//@Tags Openings
//@Accept json
//@Produce json
//@Success 200 {object} ShowOpeningResponse
//@Failure 400 {object} ErrorResponse
//@Failure 404 {object} ErrorResponse
//@Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context){
	openings := []schemas.Opening{}
	
	if err := db.Find(&openings).Error; err != nil{
			sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-opening", openings)

}