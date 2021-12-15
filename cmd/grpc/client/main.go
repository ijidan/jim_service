package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"jim_service/global"
	"jim_service/internal/jim_proto/proto_build"
	"jim_service/pkg"
	"time"
)

func main() {
	defer func() {
		global.Close()
	}()
	rpc := global.Config.Rpc
	address := fmt.Sprintf("%s:%d", rpc.Host, rpc.Port)
	connCtx,cancel:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	clientToken:=pkg.NewClientToken("")
	conn, err := grpc.DialContext(connCtx,address, grpc.WithInsecure(), grpc.WithBlock(),grpc.WithPerRPCCredentials(clientToken))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	client := proto_build.NewPingServiceClient(conn)
	ctx, cancel1 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel1()
	rsp, err1 := client.Ping(ctx, &proto_build.PingRequest{})
	if err1 != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("rsp: %s", rsp.GetMessage())
}
