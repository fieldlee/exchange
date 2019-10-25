package main

import (
	"h-exchange/httpserver"
	"h-exchange/libs/types"
	_ "h-exchange/libs/utils"
	"h-exchange/match_engine"
	"log"
	"runtime"
)


func main()  {
	runtime.GOMAXPROCS(runtime.NumCPU())

	log.Println("撮合引擎正在启动")
	go match_engine.ReceiveLimitOrderByHttp(htypes.GlobalMsg.PutOrderChanByHttp)
	match_engine.InitMarket()
	httpserver.AccessHttpInit()

}
