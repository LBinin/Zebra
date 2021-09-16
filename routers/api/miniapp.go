package api

import "github.com/gin-gonic/gin"

type ApiMiniAppRouter struct {
}

func (r ApiMiniAppRouter) Register(e *gin.Engine) {
	v1 := e.Group("v1")
	api := v1.Group("api")
	miniapp := api.Group("miniapp")
	miniapp.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
