package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SearchApi struct{}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func (e *SearchApi) Search(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		1,
		"1",
		"成功",
	})
}
