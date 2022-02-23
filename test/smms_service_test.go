package test

import (
	"flag"
	"github.com/stretchr/testify/require"
	"jim_service/pkg"
	"testing"
)

var SmMsImg = flag.String("ss_img", "", "img path required")

func TestSmMsGetToken(t *testing.T) {
	defer basic.Close()
	smMs := pkg.NewSmMs()
	token, err := smMs.GetToken(pkg.Conf.Smms.User, pkg.Conf.Smms.Password)
	require.Nil(t, err,err)
	t.Log(token)
}

func TestSmMsGetUserProfile(t *testing.T) {
	defer basic.Close()
	smMs := pkg.NewSmMs()
	token, err := smMs.GetToken(pkg.Conf.Smms.User, pkg.Conf.Smms.Password)
	require.Nil(t, err,err)
	t.Log(token)

	userProfile, err1 := smMs.GetUserProfile(token)
	require.Nil(t, err1,err1)
	t.Log(userProfile)
}

func TestSmMsUploadImage(t *testing.T) {
	defer basic.Close()
	flag.Parse()
	smMs := pkg.NewSmMs()
	token, err := smMs.GetToken(pkg.Conf.Smms.User, pkg.Conf.Smms.Password)
	require.Nil(t, err,err)

	img, err1 := smMs.UploadImage(token, *SmMsImg)
	require.Nil(t, err1,err1)
	t.Log(img)
}
