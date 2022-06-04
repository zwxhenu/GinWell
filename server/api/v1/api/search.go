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

	var phone = c.DefaultQuery("phone", "")
	var dat = c.Query("dat")
	if len(phone) < 0 {
		response.FailWithMessage("查询的手机号不能为空", c)
	}
	if dat == "" {
		topSevenPhone := phone[0:7]
		if err, info := searchService.GetPhoneInfo(topSevenPhone); err != nil {
			global.GW_LOG.Error("查询失败！", zap.Error(err))
			response.FailWithMessage("查询失败！", c)
		} else {
			info.Phone = phone
			response.OkWithDetailed(info, "获取成功"+dat, c)
		}
	} else {
		if err, info := searchService.GetPhoneInfoByDat(phone); err != nil {
			response.FailWithMessage("查询失败！", c)
		} else {
			response.OkWithDetailed(info, "获取成功"+dat, c)
		}
	}

}

func (e *SearchApi) BatchCreate(c *gin.Context) {
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
		err, info := searchService.GetPhoneInfo(v[0])
		fmt.Println(info)
		fmt.Println(err)
		if err != nil {
			if len(info.Phone) > 0 {
				break
			} else {
				phoneObj := &search.SearchPhone{Phone: v[0], Province: v[1], City: v[2], Operator: v[3]}
				err = global.GW_DB.Create(phoneObj).Error
			}
		}
		fmt.Println(err)
	}
	response.OkWithDetailed(code, msg, c)
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
		phoneObj := &search.SearchPhone{Phone: v[0], Province: v[1], City: v[2], Operator: v[3]}
		//phoneInfo := global.GW_DB.Where("phone=?", v[0]).First(&search.SearchPhone{}).Error
		err = global.GW_DB.Create(phoneObj).Error
		fmt.Println(err)
	}
	response.Result(
		code,
		"1",
		msg,
		c)
}
