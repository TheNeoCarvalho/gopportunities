package handler

import (
	"fmt"
	"net/http"

	"github.com/TheNeoCarvalho/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

//BasePath /api/v1

//@Sumary Show opening
//@Description Show a Job opening
//@Tags Openings
//@Accept json
//@Produce json
//@Param id path string true "Opening indentification"
//@Success 200 {object} ShowOpeningResponse
//@Failure 400 {object} ErrorResponse
//@Failure 404 {object} ErrorResponse
//@Router /opening/{id} [get]
func ShowOpeningHandler(ctx *gin.Context){
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}


	sendSuccess(ctx, "show-opening", opening)
}