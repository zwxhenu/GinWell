/**
 @author: zangl
 @date: 2022/5/31
 @note:
**/
package search

type SearchPhone struct {
	Phone    string `json:"Phone" gorm:"comment:手机号"`    //手机号
	Province string `json:"Province" gorm:"comment:省份"`  // 省份
	City     string `json:"City" gorm:"comment:城市"`      // 城市
	Operator string `json:"Operator" gorm:"comment:运营商"` // 运营商
}

func (SearchPhone) TableName() string {
	return "gw_phone"
}
