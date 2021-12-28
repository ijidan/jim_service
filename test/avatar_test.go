package test

import (
	"flag"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strk.kbt.io/projects/go/libravatar"
	"testing"
)

var nickName=flag.String("nick_name","","nick name required")

func TestGenAvatar(t *testing.T) {
	flag.Parse()
	avt := libravatar.New()
	email:=fmt.Sprintf("%s@%s",*nickName,"jim.com")
	avatarUrl,error:= avt.FromEmail(email)
	assert.Nil(t, error)
	t.Log(avatarUrl)

}

