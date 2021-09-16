package main

import (
	"github.com/gin-gonic/gin"
	"shop/routers"
	"shop/routers/api"
	"shop/routers/web"
)

func main() {
	//gin.DisableConsoleColor()
	r := gin.Default()
	rs := routers.NewRouters(r)
	rs.Register(web.AdminRouter{})
	rs.Register(api.ApiMiniAppRouter{})
	rs.Register(api.ApiAdminRouter{})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
