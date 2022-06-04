package router

import (
	"GinWell-Server/router/api"
)

type ApiGroup struct {
	//SystemApiGroup sy
	ApiGroup api.RouterGroup
}

var ApiGroupApp = new(ApiGroup)
