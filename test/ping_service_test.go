package test

import (
	"context"
	"jim_service/internal/jim_proto/proto_build"
	"jim_service/internal/service"
	"testing"
)

func TestPing(t *testing.T) {
	ping := service.NewPingService()
	_, err := ping.Ping(context.Background(), &proto_build.PingRequest{})
	if err != nil {
		t.Errorf("error:%v", err)
	}
}

func BenchmarkPing(b *testing.B) {
	ping := service.NewPingService()
	_, _ = ping.Ping(context.Background(), &proto_build.PingRequest{})
}