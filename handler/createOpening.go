package handler

import (
	"net/http"

	"github.com/TheNeoCarvalho/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context){
	
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role: 		request.Role,
		Company: 	request.Company,
		Location: request.Location,
		Remote: 	*request.Remote,
		Salary: 	request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil{
		logger.Errorf("Error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}
	sendSuccess(ctx, "create-opening", opening)

}