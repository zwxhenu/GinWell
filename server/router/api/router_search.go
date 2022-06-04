package api

import (
	v1 "GinWell-Server/api/v1"
	"GinWell-Server/middleware"
	"github.com/gin-gonic/gin"
)

type SearchRouter struct{}

func (e *SearchRouter) InitSearchRouter(Router *gin.RouterGroup) {
	searchRouter := Router.Group("api").Use(middleware.AntiBrush())
	//userRouterWithoutRecord := Router.Group("user")
	searchApi := v1.ApiGroupApp.ApiApiGroup
	{
		searchRouter.GET("search", searchApi.Search)     // 查询Api
		searchRouter.GET("create", searchApi.InsertData) // 创建Api
		searchRouter.GET("batch", searchApi.BatchCreate) //
	}
}
