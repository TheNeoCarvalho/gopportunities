package router

import (
	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine){
		v1:= router.Group("/api/v1")
		{
			v1.GET("/opening", func(cxt *gin.Context) {
				cxt.JSON(200, gin.H{
					"message": "hello GET",
				})
			})

			v1.POST("/opening", func(cxt *gin.Context) {
				cxt.JSON(200, gin.H{
					"message": "hello POST",
				})
			})

			v1.DELETE("/opening", func(cxt *gin.Context) {
				cxt.JSON(200, gin.H{
					"message": "hello DELETE",
				})
			})

			v1.PUT("/opening", func(cxt *gin.Context) {
				cxt.JSON(200, gin.H{
					"message": "hello PUT",
				})
			})

			v1.GET("/openings", func(cxt *gin.Context) {
				cxt.JSON(200, gin.H{
					"message": "hello GET openings",
				})
			})
		}
}