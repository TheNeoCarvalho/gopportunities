package router

import (
	"github.com/TheNeoCarvalho/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine){
	handler.InitHandler()
		v1:= router.Group("/api/v1")
		{
			v1.GET("/opening/:id", handler.ShowOpeningHandler)
			v1.POST("/opening", handler.CreateOpeningHandler)
			v1.DELETE("/opening", handler.DeleteOpeningHandler)
			v1.PUT("/opening/:id", handler.UpdateOpeningHandler)
			v1.GET("/openings", handler.ListOpeningsHandler)
		}
}