package main

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"jim_service/config"
	"jim_service/internal/dispatch"
	"jim_service/internal/server"
	"jim_service/pkg"
	"os"
	"os/signal"
	"syscall"
)

////go:embed config.yaml
//var config embed.FS

func buildTable(config config.Config) *tablewriter.Table {
	httpAddress := fmt.Sprintf("%s:%d", config.Http.Host, config.Http.Port)
	pprofAddress := fmt.Sprintf("%s:%d", config.Pprof.Host, config.Pprof.Port)
	grpcAddress := fmt.Sprintf("%s:%d", config.Rpc.Host, config.Rpc.Port)
	goPsAddress := fmt.Sprintf("%s:%d", config.Ps.Host, config.Ps.Port)

	data := [][]string{
		[]string{"1", "Application", "Jim_Service"},
		[]string{"2", "Pprof", pprofAddress},
		[]string{"3", "Grpc", grpcAddress},
		[]string{"4", "GoPs", goPsAddress},
		[]string{"5", "Http", httpAddress},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Info", "Desc"})
	table.AppendBulk(data)
	table.SetAlignment(tablewriter.ALIGN_LEFT) // Set Alignment
	return table
}

func main() {
	defer pkg.Close()
	table := buildTable(*pkg.Conf)
	table.Render()
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
		err := server.RunGrpc(*pkg.Conf, ctx)
		if err != nil {
			cancel()
			pkg.Logger.Fatalf("run grpc server:%s", err.Error())
		}
	}()

	go func() {
		err := server.RunGoPs(*pkg.Conf, ctx)
		if err != nil {
			cancel()
			pkg.Logger.Fatalf("run gops:%s", err.Error())
		}
	}()
	go func() {
		err := dispatch.SubscribeMessage(ctx)
		if err != nil {
			cancel()
			pkg.Logger.Fatalf("run kafka client:%s", err.Error())
		}
	}()


	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGILL, syscall.SIGQUIT, syscall.SIGTERM)
	<-ch
	color.Red("closing ...")
	cancel()
	color.Red("closed")
}
