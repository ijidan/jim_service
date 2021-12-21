package main

import (
	"embed"
	"github.com/sirupsen/logrus"
	"jim_service/global"
	"jim_service/internal/server"
)
//go:embed config.yaml
var config embed.FS

func main() {
	defer func() {
		global.Close()
	}()
	rpc := global.Config.Rpc
	err := server.RunServer(rpc.Host, rpc.Port)
	if err != nil {
		logrus.Fatalf("failed to listen：%v", err)
	}
}


