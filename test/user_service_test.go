package test

import (
	"context"
	"flag"
	"github.com/bxcodec/faker/v3"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/require"
	"jim_service/internal/call"
	"jim_service/internal/jim_proto/proto_build"
	"jim_service/pkg"
	"math/rand"
	"testing"
)

var basic *call.BasicCall
var client proto_build.UserServiceClient
var userNum = flag.Int("user_num", 1, "user num")
var userId = flag.Int("user_id", 0, "user id")
var userImg = flag.String("user_img", "", "img url required")

func init() {
	basic = call.NewBasicCall(pkg.Conf.Rpc.Host, pkg.Conf.Rpc.Port)
	client = proto_build.NewUserServiceClient(basic.Conn)
}

func TestCreateUser(t *testing.T) {
	defer basic.Close()
	flag.Parse()
	gender := []proto_build.Gender{proto_build.Gender_Unknown, proto_build.Gender_Male, proto_build.Gender_Female}
	for i := 0; i < *userNum; i++ {
		name:=faker.Name()
		t.Log(name)
		req := &proto_build.UserCreateRequest{
			Nickname:    name,
			Gender:      gender[rand.Intn(len(gender))],
			AvatarUrl:   "https://cdn.libravatar.org/static/img/nobody/80.png",
			Password:    "jidan123456",
			PasswordRpt: "jidan123456",
		}
		rsp, err2 := client.CreateUser(context.Background(), req)
		require.Nil(t, err2, err2)
		t.Log(rsp.User)
	}
}

func TestGetUser(t *testing.T) {
	defer basic.Close()
	flag.Parse()
	req := &proto_build.UserGetRequest{
		Id: cast.ToUint64(userId),
	}
	rsp, err := client.GetUser(context.Background(), req)
	require.Nil(t, err, err)
	t.Log(rsp.User)
}

func TestQueryUser(t *testing.T) {
	defer basic.Close()
	req := &proto_build.UserQueryRequest{
		Keyword:  "jidan",
		Page:     1,
		PageSize: 10,
	}
	rsp, err := client.QueryUser(context.Background(), req)
	require.Nil(t, err, err)
	t.Log(rsp.GetUser())
	t.Logf("-------------------")
	t.Log(rsp.GetPager())
}

func TestUpdateAvatar(t *testing.T) {
	defer basic.Close()
	flag.Parse()

	req := &proto_build.UpdateAvatarRequest{Url: *userImg}
	rsp, err := client.UpdateAvatar(context.Background(), req)
	require.Nil(t, err, err)

	t.Log(rsp)
}
