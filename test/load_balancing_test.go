package test

import (
	"context"
	"flag"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/resolver"
	"jim_service/global"
	"jim_service/internal/jim_proto/proto_build"
	"jim_service/pkg"
	"testing"
)


var serviceName=flag.String("service_ping","service_ping","service name required!")

func TestGetServiceServerList(t *testing.T) {
	defer global.Close()
	flag.Parse()
	serviceDiscovery:=pkg.NewServiceDiscovery(global.ClientV3)
	serverList:=serviceDiscovery.GetServerList(*serviceName)
	t.Log(serverList)
}

func TestPingService(t *testing.T)  {
	defer global.Close()
	builder:=pkg.NewJResolverBuilder(global.ClientV3)
	resolver.Register(builder)
	conn, err := pkg.GetRpcConn("service_ping")
	assert.Nil(t, err)
	client := proto_build.NewPingServiceClient(conn)
	rsp, err1 := client.Ping(context.Background(), &proto_build.PingRequest{})
	assert.Nil(t, err1)
	t.Log(rsp.GetMessage())
}
