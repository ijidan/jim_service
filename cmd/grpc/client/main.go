package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"jim/global"
	"jim/internal/rpc/proto_build"
	"time"
)

func main() {
	defer func() {
		global.Close()
	}()
	rpc := global.Config.Rpc
	address := fmt.Sprintf("%s:%d", rpc.Host, rpc.Port)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)
	client := proto_build.NewPingServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rsp, err1 := client.Ping(ctx, &proto_build.PingRequest{})
	if err1 != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("rsp: %s", rsp.GetMessage())
}
