package api

import "github.com/gin-gonic/gin"

type ApiAdminRouter struct {
}

func (r ApiAdminRouter) Register(e *gin.Engine) {
	v1 := e.Group("v1")
	api := v1.Group("api")
	admin := api.Group("admin")
	admin.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
