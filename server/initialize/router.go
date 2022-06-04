package initialize

import (
	"GinWell-Server/middleware"
	"GinWell-Server/router"
	"github.com/gin-gonic/gin"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()
	apiRouter := router.ApiGroupApp.ApiGroup
	PublicGroup := Router.Group("")
	PublicGroup.Use(middleware.AntiBrush())
	{
		apiRouter.InitSearchRouter(PublicGroup)
	}
	return Router
}
