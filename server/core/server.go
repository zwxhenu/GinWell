package core

import (
	"GinWell-Server/global"
	"GinWell-Server/initialize"
	"fmt"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	address := fmt.Sprintf(":%d", global.GW_CONFIG.System.Addr)
	Router := initialize.Routers()
	s := initServer(address, Router)
	s.ListenAndServe()
}
