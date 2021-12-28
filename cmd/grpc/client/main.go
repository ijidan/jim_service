package main

import (
	"github.com/sirupsen/logrus"
	"jim_service/internal/call"
	"jim_service/internal/jim_proto/proto_build"
)

func main() {
	basic:=call.NewBasicCall()
	defer basic.Close()
	client := proto_build.NewPingServiceClient(basic.Conn)
	rsp, err := client.Ping(basic.DialCtx, &proto_build.PingRequest{})
	if err != nil {
		logrus.Fatalf("could not greet: %v", err)
	}
	logrus.Println(rsp.GetMessage())
}
