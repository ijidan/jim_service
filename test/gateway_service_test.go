package test

import (
	"github.com/stretchr/testify/require"
	"jim_service/internal/call"
	"jim_service/internal/jim_proto/proto_build"
	"jim_service/pkg"
	"testing"
)

func TestSendToAll(t *testing.T) {
	basic:=call.NewBasicCall(pkg.Conf.Rpc.Host,pkg.Conf.Rpc.Port)
	defer basic.Close()

	content:=`{"type":"text","data":{"sender_id":"10000","sender_name":"管理员","sender_avatar":"https://www.baidu.com/img/flexible/logo/pc/result.png","receiver_id":"","content":"hello client 2"}}`
	client:=proto_build.NewGatewayServiceClient(basic.Conn)
	rsp,err:=client.SendToAll(basic.DialCtx,&proto_build.SendToAllRequest{
		Data:     []byte(content),
		SenderId: "10000",
	})
	require.Nil(t, err,err)
	t.Log(rsp.String())
}