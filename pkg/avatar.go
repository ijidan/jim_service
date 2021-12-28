package pkg

import (
	"fmt"
	"strk.kbt.io/projects/go/libravatar"
)

func GetAvatar(nickName string) (string,error) {
	avt := libravatar.New()
	email:=fmt.Sprintf("%s@%s",nickName,"jim.com")
	avatarUrl,err:= avt.FromEmail(email)
	return avatarUrl,err
}