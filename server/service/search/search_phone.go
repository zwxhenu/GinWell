/**
 @author: zangl
 @date: 2022/5/31
 @note:
**/
package search

import (
	"GinWell-Server/global"
	"GinWell-Server/model/search"
	"bytes"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"io/ioutil"
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

func (searchService *SearchService) GetPhoneInfo(phone string) (err error, p *search.SearchPhone) {
	var phoneInfo search.SearchPhone
	err = global.GW_DB.Where("`phone` = ?", phone).First(&phoneInfo).Error
	return err, &phoneInfo
}

type PhoneRecord struct {
	PhoneNum string
	Province string
	City     string
	ZipCode  string
	AreaZone string
	CardType string
}

func (SearchService *SearchService) GetPhoneInfoByDat(phone string) (err error, pr *PhoneRecord) {
	var (
		intLen  int32 = 4
		charLen int32 = 1
	)
	const (
		CMCC   byte = iota + 0x01 //中国移动
		CUCC                      //中国联通
		CTCC                      //中国电信
		CTCC_v                    //电信虚拟运营商
		CUCC_v                    //联通虚拟运营商
		CMCC_v
	)
	var (
		CardTypemap = map[byte]string{
			CMCC:   "中国移动",
			CUCC:   "中国联通",
			CTCC:   "中国电信",
			CTCC_v: "中国电信虚拟运营商",
			CUCC_v: "中国联通虚拟运营商",
			CMCC_v: "中国移动虚拟运营商",
		}
	)

	datFile := global.GW_CONFIG.Dat.Dir + "phone.dat"
	//datFile := "./resource/dat/phone.dat"
	content, err := ioutil.ReadFile(datFile)
	if err != nil {
		global.GW_LOG.Error("打开phone.dat文件失败")
		return err, nil
	}
	totalLen := int32(len(content))
	firstOffset := GetFourByte(content[intLen : intLen*2])
	if len(phone) < 7 || len(phone) > 11 {
		global.GW_LOG.Error("输入手机号格式错误")
		return err, nil
	}
	var left int32
	topSevenMobile, err := GetN(phone[0:7])
	if err != nil {
		return err, nil
	}
	right := (totalLen - firstOffset) / 9

	for {
		if left > right {
			break
		}
		mid := (left + right) / 2
		offset := firstOffset + mid*9
		if offset >= totalLen {
			break
		}
		curPhone := GetFourByte(content[offset : offset+intLen])
		recordOffset := GetFourByte(content[offset+intLen : offset+intLen*2])
		cartType := content[offset+intLen*2 : offset+intLen*2+charLen][0]
		topSevenMobileInt := int32(topSevenMobile)
		switch {
		case curPhone > topSevenMobileInt:
			right = mid - 1
		case curPhone < topSevenMobileInt:
			left = mid + 1
		default:
			cByte := content[recordOffset:]
			endOffset := int32(bytes.Index(cByte, []byte("\000")))
			data := bytes.Split(cByte[:endOffset], []byte("|"))
			cardStr, ok := CardTypemap[cartType]
			if !ok {
				cardStr = "未知电信运营商"
			}
			pr = &PhoneRecord{
				PhoneNum: phone,
				Province: string(data[0]),
				City:     string(data[1]),
				ZipCode:  string(data[2]),
				AreaZone: string(data[3]),
				CardType: cardStr,
			}
			return err, pr
		}
	}
	return errors.New("手机号不存在"), nil
}

func GetFourByte(b []byte) int32 {
	if len(b) < 4 {
		return 0
	}
	fmt.Println(int32(b[1]) << 8)
	return int32(b[0]) | int32(b[1])<<8 | int32(b[2])<<16 | int32(b[3])<<24
}

func GetN(s string) (uint32, error) {
	var n, cutOff, maxVal uint32
	i := 0
	base := 10
	cutOff = (1<<32-1)/10 + 1
	maxVal = 1<<uint(32) - 1
	for ; i < len(s); i++ {
		var v byte
		d := s[i]

		switch {
		case '0' <= d && d <= '9':
			v = d - '0'
		case 'a' <= d && d <= 'z':
			v = d - 'a' + 10
		case 'A' <= d && d <= 'Z':
			v = d - 'A' + 10
		default:
			return 0, errors.New("invalid syntax")
		}
		if v >= byte(base) {
			return 0, errors.New("invalid syntax")
		}
		if n >= cutOff {
			n = 1<<3 - 1
			return n, errors.New("value out of range")
		}
		n *= uint32(base)
		n1 := n + uint32(v)
		if n1 < n || n1 > maxVal {
			n = 1<<32 - 1
			return n, errors.New("value out of range")
		}
		n = n1
	}
	return n, nil
}
