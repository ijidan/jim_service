package call

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"jim_service/pkg"
	"time"
)

type BasicCall struct {
	Conn       *grpc.ClientConn
	ConnCancel context.CancelFunc
	DialCtx    context.Context
	DialCancel context.CancelFunc
}

func (c *BasicCall) GetClientConn(host string,port uint) (*grpc.ClientConn, context.CancelFunc) {
	defer func() {
		pkg.Close()
	}()
	address := fmt.Sprintf("%s:%d", host, port)
	connCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	clientToken := pkg.NewClientToken("")
	conn, err := grpc.DialContext(connCtx, address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithPerRPCCredentials(clientToken),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
		)
	if err != nil {
		logrus.Fatalf("did not connect: %v", err)
	}
	c.Conn = conn
	c.ConnCancel = cancel
	return conn, cancel
}

func (c *BasicCall) GetContext() (context.Context, context.CancelFunc) {
	ctx:=context.Background()
	newCtx:=metadata.AppendToOutgoingContext(ctx,"X-Request-Source","grpc/client")
	c.DialCtx, c.DialCancel = context.WithTimeout(newCtx, 5*time.Second)
	return c.DialCtx, c.DialCancel
}

func (c *BasicCall) Close() {
	_ = c.Conn.Close()
	c.ConnCancel()
	c.DialCancel()
}

func NewBasicCall(host string,port uint) *BasicCall  {
	basic:=&BasicCall{}
	basic.GetClientConn(host,port)
	basic.GetContext()
	return basic
}