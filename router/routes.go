package router

import (
	docs "github.com/TheNeoCarvalho/gopportunities/docs"
	"github.com/TheNeoCarvalho/gopportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initRoutes(router *gin.Engine){
	handler.InitHandler()
	basePath := "/api/v1"

	docs.SwaggerInfo.BasePath = basePath
	v1:= router.Group("/api/v1")
		{
			v1.GET("/opening/:id", handler.ShowOpeningHandler)
			v1.POST("/opening", handler.CreateOpeningHandler)
			v1.DELETE("/opening", handler.DeleteOpeningHandler)
			v1.PUT("/opening/:id", handler.UpdateOpeningHandler)
			v1.GET("/openings", handler.ListOpeningsHandler)
		}
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}