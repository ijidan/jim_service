package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"jim_service/global"
	"jim_service/internal/jim_proto/proto_build"
	"jim_service/internal/service"
	"testing"
)


func TestPing(t *testing.T) {
	defer global.Close()
	pingService := service.NewPingService(global.Config)
	rsp, err := pingService.Ping(context.Background(), &proto_build.PingRequest{})
	assert.Nil(t, err)
	t.Log(rsp,err)
}



func BenchmarkPing(b *testing.B) {
	defer global.Close()
	for i:=0;i<b.N;i++{
		pingService := service.NewPingService(global.Config)
		_, _ = pingService.Ping(context.Background(), &proto_build.PingRequest{})
	}
}
