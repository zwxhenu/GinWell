/**
 @author: zangl
 @date: 2022/6/9
 @note:
**/
package main

import (
	"GinWell-Server/service"
	"fmt"
)

func main() {
	var phone = "13146701366"
	//var phoneSearch = new(search.SearchService)
	var searchService = service.SearchServiceGroupApp.SearchServiceGroup.SearchService
	pr, err := searchService.GetPhoneInfoByDat(phone)
	if err != nil {
		panic(err)
	}
	fmt.Print(pr)
}
