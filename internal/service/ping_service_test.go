package service

import (
	"context"
	"jim_service/internal/jim_proto/proto_build"
	"testing"
)

func TestPing(t *testing.T)  {
	ping:=NewPingService()
	_,err:=ping.Ping(context.Background(),&proto_build.PingRequest{})
	if err!=nil{
		t.Errorf("error:%v",err)
	}
}

func BenchmarkPing(b *testing.B) {
	ping:=NewPingService()
	_,_=ping.Ping(context.Background(),&proto_build.PingRequest{})
}