package call

import "C"
import (
	"github.com/sirupsen/logrus"
	"jim_service/internal/jim_proto/proto_build"
)

type PingCall struct {
	*BasicCall
}

func (c *PingCall) Ping()  (*proto_build.PingResponse, error){
	client := proto_build.NewPingServiceClient(c.Conn)
	rsp, err := client.Ping(c.DialCtx, &proto_build.PingRequest{})
	if err != nil {
		logrus.Fatalf("could not greet: %v", err)
	}
	return rsp,err
}

func NewPingCall(host string,port uint) *PingCall  {
	basic:=NewBasicCall(host,port)
	return &PingCall{BasicCall:basic}
}
