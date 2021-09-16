package routers

import "github.com/gin-gonic/gin"

type Routers struct {
	*gin.Engine
}

type Router interface {
	Register(e *gin.Engine)
}

func NewRouters(e *gin.Engine) Routers {
	return Routers{
		e,
	}
}

func (rs Routers) Register(r Router) {
	r.Register(rs.Engine)
}
