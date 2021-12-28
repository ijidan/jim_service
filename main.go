package main

import (
	"embed"
	"jim_service/internal/server"
	"jim_service/internal/service"
	"jim_service/pkg"
)
//go:embed config.yaml
var config embed.FS

func main() {
	defer pkg.Close()
	clientV3:=service.NewClientV3(pkg.Conf.Etcd.Host,pkg.Conf.Etcd.Timeout)
	err := server.RunServer(clientV3,pkg.Conf)
	if err != nil {
		pkg.Logger.Fatalf("failed to listen：%v", err)
	}
}


