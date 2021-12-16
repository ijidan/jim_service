package main

import (
	"github.com/sirupsen/logrus"
	"jim_service/global"
	"jim_service/internal/server"
)

func main() {
	defer func() {
		global.Close()
	}()
	rpc := global.Config.Rpc
	err := server.RunServer(rpc.Host, rpc.Port)
	if err != nil {
		logrus.Fatalln("failed to listen：", err)
	}
}


