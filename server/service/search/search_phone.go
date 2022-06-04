/**
 @author: zangl
 @date: 2022/5/31
 @note:
**/
package search

import (
	"GinWell-Server/global"
	"GinWell-Server/model/search"
	"errors"
	"gorm.io/gorm"
)

type SearchService struct {
}

func (searchService *SearchService) insert(p search.SearchPhone) (err error, phoneInter search.SearchPhone) {
	var phone search.SearchPhone
	if !errors.Is(global.GW_DB.Where("phone=?", p.Phone).First(&phone).Error, gorm.ErrRecordNotFound) {
		return errors.New("手机号已录入"), phoneInter
	}
	err = global.GW_DB.Create(&p).Error
	return err, p
}

func (searchService *SearchService) GetPhoneInfo(phone int) (err error, p *search.SearchPhone) {
	var phoneInfo search.SearchPhone
	err = global.GW_DB.Where("`phone` = ?", phone).First(&phoneInfo).Error
	return err, &phoneInfo
}
