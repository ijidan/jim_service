package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"jim_service/internal/call"
	"jim_service/internal/jim_proto/proto_build"
	"jim_service/internal/service"
	"jim_service/pkg"
	"testing"
)


func TestPing(t *testing.T) {
	defer pkg.Close()

	pingService := service.NewPingService(pkg.Conf)
	rsp, err := pingService.Ping(context.Background(), &proto_build.PingRequest{})
	assert.Nil(t, err)
	t.Log(rsp.GetMessage())
	t.Log(rsp,err)
}

func TestClientPing(t *testing.T) {
	basic:=call.NewBasicCall(pkg.Conf.Rpc.Host,pkg.Conf.Rpc.Port)
	defer basic.Close()
	client := proto_build.NewPingServiceClient(basic.Conn)
	rsp, err := client.Ping(basic.DialCtx, &proto_build.PingRequest{})
	assert.Nil(t, err)
	t.Log(rsp.GetMessage())
}

func BenchmarkPing(b *testing.B) {
	defer pkg.Close()
	for i:=0;i<b.N;i++{
		pingService := service.NewPingService(pkg.Conf)
		_, _ = pingService.Ping(context.Background(), &proto_build.PingRequest{})
	}
}

