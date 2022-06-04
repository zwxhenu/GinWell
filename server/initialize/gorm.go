/**
 @author: zangl
 @date: 2022/5/29
 @note:
**/
package initialize

import (
	"GinWell-Server/global"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.GW_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}
