package repository

import (
	"fmt"
	"jim_service/pkg"
)

func SendEmail(receiver map[string]string, cc map[string]string, subject string, content string) error {
	conf := pkg.Conf.Email
	email := pkg.NewEmail(conf.Smtp, conf.Port, conf.Ssl, conf.Account, conf.Password)
	return email.Send(receiver, cc, subject, content)
}

func SendNewUserEmail(userId uint64 ,extra string) error {
	email := pkg.NewEmail(pkg.Conf.Email.Smtp, pkg.Conf.Email.Port, pkg.Conf.Email.Ssl, pkg.Conf.Email.Account, pkg.Conf.Email.Password)
	subject := "新用户注册"
	content := fmt.Sprintf("<p>新用户注册成功，用户Id：%d,附件信息：%s</p>", userId,extra)
	return email.Send(pkg.Conf.Manager.Email, nil, subject, content)
}

func SendErrorEmail(err error) error  {
	conf := pkg.Conf.Email
	email := pkg.NewEmail(conf.Smtp, conf.Port, conf.Ssl, conf.Account, conf.Password)
	subject:="程序错误信息"
	content := fmt.Sprintf("<p>错误信息：%+v</p>", err)
	return email.Send(pkg.Conf.Manager.Email, nil, subject, content)
}
