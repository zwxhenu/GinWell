/**
 @author: zangl
 @date: 2022/5/31
 @note:
**/
package service

import "GinWell-Server/service/search"

type SearchServiceGroup struct {
	SearchServiceGroup search.SearchGroupService
}

var SearchServiceGroupApp = new(SearchServiceGroup)
