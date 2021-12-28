package test

import (
	"context"
	"fmt"
	"github.com/golang-module/carbon/v2"
	"github.com/stretchr/testify/assert"
	"jim_service/internal/call"
	"jim_service/internal/jim_proto/proto_build"
	"jim_service/pkg"
	"testing"
)

var serviceUser = "service_user"

func TestCreateUser(t *testing.T) {
	basic:=call.NewBasicCall(pkg.Conf.Rpc.Host,pkg.Conf.Rpc.Port)
	defer basic.Close()
	t.Log(basic)
	client := proto_build.NewUserServiceClient(basic.Conn)
	req := &proto_build.CreateUserRequest{
		Nickname:    fmt.Sprintf("jidan%d",carbon.Now().Timestamp()),
		Gender:      proto_build.Gender_Male,
		AvatarUrl:   "https://cdn.libravatar.org/static/img/nobody/80.png",
		Password:    "123456",
		PasswordRpt: "123456",
	}
	rsp, err2 := client.CreateUser(context.Background(), req)
	assert.Nil(t, err2)
	t.Log(rsp.User.String())

}
