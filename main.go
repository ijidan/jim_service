package main

import (
	"embed"
	"jim_service/global"
	"jim_service/internal/server"
)
//go:embed config.yaml
var config embed.FS

func main() {
	defer global.Close()
	err := server.RunServer(global.Config)
	if err != nil {
		global.Logger.Fatalf("failed to listen：%v", err)
	}
}


