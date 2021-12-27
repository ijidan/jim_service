package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/resolver"
	"jim_service/global"
	"jim_service/internal/jim_proto/proto_build"
	"jim_service/pkg"
)

func main()  {
	defer global.Close()
	builder:=pkg.NewJResolverBuilder(global.ClientV3)
	resolver.Register(builder)
	conn, err := pkg.GetRpcConn("service_ping")
	if err!=nil{
		logrus.Fatalf(err.Error())
	}
	client := proto_build.NewPingServiceClient(conn)
	rsp, err1 := client.Ping(context.Background(), &proto_build.PingRequest{})
	if err1!=nil{
		logrus.Fatalf(err1.Error())
	}
	logrus.Print(rsp.GetMessage())
}
