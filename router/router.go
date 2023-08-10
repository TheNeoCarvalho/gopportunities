package router

import "github.com/gin-gonic/gin"

func Init() {

	router:= gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":3333") // listen and serve on 0.0.0.0:8080

}