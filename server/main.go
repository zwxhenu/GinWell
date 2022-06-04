package main

import (
	"GinWell-Server/core"
	"GinWell-Server/global"
	"GinWell-Server/initialize"
	"go.uber.org/zap"
)

func main() {
	global.GW_VP = core.Viper() // 初始化Viper
	global.GW_LOG = core.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(global.GW_LOG)
	global.GW_DB = initialize.Gorm() // gorm连接数据库
	core.RunServer()
}
