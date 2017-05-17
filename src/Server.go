package main

import (
	"utils"
	"myserver"
	"common"
	"runtime"
)

func main()  {
	runtime.GOMAXPROCS(6)
	common.HttpToSocket = make(chan string)
	common.SocketToHttp = make(chan string)
	msg := make(chan bool, 0)
	go myserver.StartHttpServer(utils.Cfg.HttpPort) // alexa
	go myserver.StartSocketServer(utils.Cfg.SocketHost, utils.Cfg.BeatingInterval)
	<- msg
}
