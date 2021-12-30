package test

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang-module/carbon/v2"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"jim_service/internal/call"
	"jim_service/internal/jim_proto/proto_build"
	"jim_service/pkg"
	"testing"
)


var basic *call.BasicCall
var client proto_build.UserServiceClient
var userId=flag.Int("user_id",0,"user id required")


func init()  {
	basic=call.NewBasicCall(pkg.Conf.Rpc.Host,pkg.Conf.Rpc.Port)
	client= proto_build.NewUserServiceClient(basic.Conn)
}


func TestCreateUser(t *testing.T) {
	defer basic.Close()
	req := &proto_build.CreateUserRequest{
		Nickname:    fmt.Sprintf("jidan%d",carbon.Now().Timestamp()),
		Gender:      proto_build.Gender_Male,
		AvatarUrl:   "https://cdn.libravatar.org/static/img/nobody/80.png",
		Password:    "123456",
		PasswordRpt: "123456",
	}
	rsp, err2 := client.CreateUser(context.Background(), req)
	assert.Nil(t, err2)
	if err2==nil{
		t.Log(rsp.User)
	}
}

func TestGetUser(t *testing.T)  {
	defer basic.Close()
	req:=&proto_build.GetUserRequest{
		Id: cast.ToUint64(userId),
	}
	rsp,err:=client.GetUser(context.Background(),req)
	assert.Nil(t, err)
	if err==nil{
		t.Log(rsp.User)
	}
}

func TestQueryUser(t *testing.T) {
	defer basic.Close()
	req:=&proto_build.QueryUserRequest{
		Keyword:  "jidan",
		Page:     1,
		PageSize: 10,
	}
	rsp,err:=client.QueryUser(context.Background(),req)
	assert.Nil(t, err)
	if err==nil{
		t.Log(rsp.GetUser())
		t.Log(rsp.GetPager())
	}
}
