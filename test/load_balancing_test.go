package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"jim_service/global"
	"jim_service/internal/jim_proto/proto_build"
	"jim_service/pkg"
	"testing"
)


func TestGetServiceServerList(t *testing.T) {
	defer global.Close()
	pkg.LoadBalancingRegister(global.ClientV3,global.Config.App.Name)
	conn, err := pkg.GetRpcConn(global.Config.App.Name,"ping_service")
	assert.Errorf(t, err,err.Error())
	t.Log(conn)
}

func TestPingService(t *testing.T)  {
	defer global.Close()
	pkg.LoadBalancingRegister(global.ClientV3,global.Config.App.Name)
	conn, err := pkg.GetRpcConn(global.Config.App.Name,"ping_service")
	assert.Errorf(t, err,err.Error())
	client := proto_build.NewPingServiceClient(conn)
	rsp, err1 := client.Ping(context.Background(), &proto_build.PingRequest{})
	assert.Errorf(t, err1,err1.Error())
	t.Log(rsp)
}
