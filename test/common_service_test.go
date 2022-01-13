package test

import (
	"context"
	"flag"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
	"io"
	"jim_service/internal/call"
	"jim_service/internal/jim_proto/proto_build"
	"jim_service/pkg"
	"testing"
)

var commonServiceBasic *call.BasicCall
var CommonServiceClient proto_build.CommonServiceClient
var commonImg = flag.String("common_img", "", "img path required")

func init() {
	commonServiceBasic = call.NewBasicCall(pkg.Conf.Rpc.Host, pkg.Conf.Rpc.Port)
	CommonServiceClient=proto_build.NewCommonServiceClient(commonServiceBasic.Conn)
}

func TestUploadImage(t *testing.T) {
	defer basic.Close()
	flag.Parse()

	osFs := afero.NewOsFs()
	content, err := afero.ReadFile(osFs, *commonImg)
	require.Nil(t, err, err)

	uploadImageClient,err1:=CommonServiceClient.UploadImage(context.Background())
	require.Nil(t, err1,err1)

	req:=&proto_build.UploadImageRequest{Content: content}
	err2 := uploadImageClient.Send(req)
	if err2 != nil {
		if err2 == io.EOF {
			err3 := uploadImageClient.CloseSend()
			require.Nil(t, err3, err3)
		} else {
			require.Nil(t, err2, err2)
		}
	}
	rsp, err4 := uploadImageClient.CloseAndRecv()
	require.Nil(t, err4, err4)
	t.Log(rsp)
}
