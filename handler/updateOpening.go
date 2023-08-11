package handler

import (
	"net/http"

	"github.com/TheNeoCarvalho/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

//BasePath /api/v1

//@Sumary Update opening
//@Description Update a Job opening
//@Tags Openings
//@Accept json
//@Produce json
//@Param request body CreateOpeningRequest true "Request body"
//@Param id path string true "Opening indentification"
//@Success 200 {object} UpdateOpeningResponse
//@Failure 400 {object} ErrorResponse
//@Failure 404 {object} ErrorResponse
//@Router /opening/{id} [put]
func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Param("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "paramenter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil{
		logger.Errorf("error updateing opening: %v", err.Error())
		return
	}

	sendSuccess(ctx, "update-opening", opening)


}
