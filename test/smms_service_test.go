package test

import (
	"flag"
	"jim_service/pkg"
	"testing"
)

var img = flag.String("img", "", "img path required")

func TestSmMsGetToken(t *testing.T) {
	defer basic.Close()
	smMs := pkg.NewSmMs()
	token, err := smMs.GetToken(pkg.Conf.Smms.User, pkg.Conf.Smms.Password)
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(token)
}

func TestSmMsGetUserProfile(t *testing.T) {
	defer basic.Close()
	smMs := pkg.NewSmMs()
	token, err := smMs.GetToken(pkg.Conf.Smms.User, pkg.Conf.Smms.Password)
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(token)
	userProfile, err1 := smMs.GetUserProfile(token)
	if err1 != nil {
		t.Fatal(err1)
	}
	t.Log(userProfile)
}

func TestSmMsUploadImage(t *testing.T) {
	defer basic.Close()
	flag.Parse()
	smMs := pkg.NewSmMs()
	token, err := smMs.GetToken(pkg.Conf.Smms.User, pkg.Conf.Smms.Password)
	if err!=nil{
		t.Fatal(err)
	}
	img, err1 := smMs.UploadImage(token, *img)
	if err1!=nil{
		t.Fatal(err1)
	}
	t.Log(img)
}
