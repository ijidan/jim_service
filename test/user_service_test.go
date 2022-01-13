package test

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang-module/carbon/v2"
	"github.com/spf13/afero"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/require"
	"io"
	"jim_service/internal/call"
	"jim_service/internal/jim_proto/proto_build"
	"jim_service/pkg"
	"testing"
)

var basic *call.BasicCall
var client proto_build.UserServiceClient
var userNum = flag.Int("user_num", 1, "user num")
var userId = flag.Int("user_id", 0, "user id")
var userImg = flag.String("user_img", "", "img path required")

func init() {
	basic = call.NewBasicCall(pkg.Conf.Rpc.Host, pkg.Conf.Rpc.Port)
	client = proto_build.NewUserServiceClient(basic.Conn)
}

func TestCreateUser(t *testing.T) {
	defer basic.Close()
	flag.Parse()
	for i := 0; i < *userNum; i++ {
		req := &proto_build.CreateUserRequest{
			Nickname:    fmt.Sprintf("jidan%d", carbon.Now().Timestamp()),
			Gender:      proto_build.Gender_Male,
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
	req := &proto_build.GetUserRequest{
		Id: cast.ToUint64(userId),
	}
	rsp, err := client.GetUser(context.Background(), req)
	require.Nil(t, err, err)
	t.Log(rsp.User)
}

func TestQueryUser(t *testing.T) {
	defer basic.Close()
	req := &proto_build.QueryUserRequest{
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

func TestUploadAvatar(t *testing.T) {
	defer basic.Close()
	flag.Parse()
	osFs := afero.NewOsFs()
	content, err := afero.ReadFile(osFs, *userImg)
	require.Nil(t, err, err)

	uploadAvatarContent := &proto_build.UploadAvatarRequest_Content{Content: content}
	req := &proto_build.UploadAvatarRequest{Data: uploadAvatarContent}
	uploadAvatarClient, err1 := client.UpdateAvatar(context.Background())
	require.Nil(t, err1, err1)

	err2 := uploadAvatarClient.Send(req)
	if err2 != nil {
		if err2 == io.EOF {
			err3 := uploadAvatarClient.CloseSend()
			require.Nil(t, err3, err3)
		} else {
			require.Nil(t, err2, err2)
		}
	}

	rsp, err4 := uploadAvatarClient.CloseAndRecv()
	require.Nil(t, err4, err4)
	t.Log(rsp)
}
