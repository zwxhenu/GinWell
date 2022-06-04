package api

import (
	"GinWell-Server/global"
	"GinWell-Server/model/common/response"
	"GinWell-Server/model/search"
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
)

type SearchApi struct{}

func (e *SearchApi) Search(c *gin.Context) {
	if err, info := searchService.GetPhoneInfo(1310001); err != nil {
		global.GW_LOG.Error("查询失败！", zap.Error(err))
		response.FailWithMessage("查询失败！", c)
	} else {
		response.OkWithDetailed(info, "获取成功", c)
	}
}

func (e *SearchApi) BatchCreate(c *gin.Context) {
	f, err := os.Open(global.GW_CONFIG.Excel.Dir + "phonetmp.csv")
	code := 0
	msg := "成功"
	if err != nil {
		global.GW_LOG.Error("打开csv文件失败")
		code = 1
		msg = "文件打开失败"
	}
	reader := csv.NewReader(f)
	prdData, err := reader.ReadAll()
	if err != nil {
		global.GW_LOG.Error("读取csv文件失败")
		code = 1
		msg = "读取csv失败"
	}
	for _, v := range prdData {
		phone := &search.SearchPhone{Phone: v[0], Province: v[1], City: v[2], Operator: v[3]}
		//phoneInfo := global.GW_DB.Where("phone=?", v[0]).First(&search.SearchPhone{}).Error
		err = global.GW_DB.Create(phone).Error
		fmt.Println(err)
	}
	response.Result(
		code,
		"1",
		msg,
		c)
}

func (e *SearchApi) InsertData(c *gin.Context) {
	f, err := os.Open(global.GW_CONFIG.Excel.Dir + "phone.csv")
	code := 0
	msg := "成功"
	if err != nil {
		global.GW_LOG.Error("打开csv文件失败")
		code = 1
		msg = "文件打开失败"
	}
	reader := csv.NewReader(f)
	prdData, err := reader.ReadAll()
	if err != nil {
		global.GW_LOG.Error("读取csv文件失败")
		code = 1
		msg = "读取csv失败"
	}
	for _, v := range prdData {
		phone := &search.SearchPhone{Phone: v[0], Province: v[1], City: v[2], Operator: v[3]}
		//phoneInfo := global.GW_DB.Where("phone=?", v[0]).First(&search.SearchPhone{}).Error
		err = global.GW_DB.Create(phone).Error
		fmt.Println(err)
	}
	response.Result(
		code,
		"1",
		msg,
		c)
}
