package main

import (
	"context"
	"embed"
	"github.com/fatih/color"
	"jim_service/internal/dispatch"
	"jim_service/internal/server"
	"jim_service/pkg"
	"os"
	"os/signal"
	"syscall"
)

//go:embed config.yaml
var config embed.FS

func main() {
	defer pkg.Close()
	//server.RunFunc()
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		err := server.RunHttp(*pkg.Conf, ctx)
		if err != nil {
			cancel()
			pkg.Logger.Fatalf("run http server:%s", err.Error())
		}
	}()
	go func() {
		err := server.RunPprof(*pkg.Conf, ctx)
		if err != nil {
			cancel()
			pkg.Logger.Fatalf("run pprof server:%s", err.Error())
		}
	}()
	go func() {
		err:= server.RunGrpc(*pkg.Conf, ctx)
		if err != nil {
			cancel()
			pkg.Logger.Fatalf("run grpc server:%s", err.Error())
		}
	}()

	go func() {
		err:=server.RunGoPs(*pkg.Conf,ctx)
		if err!=nil{
			cancel()
			pkg.Logger.Fatalf("run gops:%s", err.Error())
		}
	}()
	go func() {
		err:=dispatch.SubscribeMessage(ctx)
		if err!=nil{
			cancel()
			pkg.Logger.Fatalf("run kafka client:%s", err.Error())
		}
	}()

	color.Green("begin stop")
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGILL, syscall.SIGQUIT, syscall.SIGTERM)
	<-ch
	color.Red("closing ...")
	cancel()
	color.Red("closed")
}
