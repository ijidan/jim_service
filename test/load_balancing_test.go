package test

import (
	"context"
	"flag"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/resolver"
	"jim_service/internal/jim_proto/proto_build"
	"jim_service/internal/service"
	"jim_service/pkg"
	"testing"
)


var serviceName=flag.String("service_ping","service_ping","service name required!")

func TestGetServiceServerList(t *testing.T) {
	defer pkg.Close()
	flag.Parse()

	clientV3:=service.NewClientV3(pkg.Conf.Etcd.Host,pkg.Conf.Etcd.Timeout)
	serviceDiscovery:=service.NewServiceDiscovery(clientV3)
	serverList:=serviceDiscovery.GetServerList(*serviceName)
	t.Log(serverList)
}

func TestPingService(t *testing.T)  {
	defer pkg.Close()
	clientV3:=service.NewClientV3(pkg.Conf.Etcd.Host,pkg.Conf.Etcd.Timeout)
	builder:=service.NewJResolverBuilder(clientV3)
	resolver.Register(builder)
	conn, err := service.GetRpcConn("service_ping")
	require.Nil(t, err,err)
	client := proto_build.NewPingServiceClient(conn)
	rsp, err1 := client.Ping(context.Background(), &proto_build.PingRequest{})
	require.Nil(t, err1,err1)
	t.Log(rsp.GetMessage())
}
