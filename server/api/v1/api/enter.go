package api

import "GinWell-Server/service"

type ApiGroup struct {
	SearchApi
}

var (
	searchService = service.SearchServiceGroupApp.SearchServiceGroup.SearchService
)
