package web

import "github.com/gin-gonic/gin"

type AdminRouter struct {
}

func (r AdminRouter) Register(e *gin.Engine) {
	v1 := e.Group("v1")
	admin := v1.Group("admin")
	admin.GET("/welcome", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
