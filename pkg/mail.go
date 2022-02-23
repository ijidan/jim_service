package pkg

import (
	"fmt"
	"github.com/spf13/cast"
	"gopkg.in/gomail.v2"
)

type Email struct {
	Smtp     string
	Port     uint64
	Ssl      bool
	Account  string
	Password string
}

func (e *Email) Send(receiver map[string]string, cc map[string]string, subject string, content string) error {
	m := gomail.NewMessage()
	for name, address := range receiver {
		item := fmt.Sprintf("%s<%s>", name, address)
		m.SetHeader("From", item)
	}
	if len(cc) > 0 {
		for name, address := range cc {
			item := fmt.Sprintf("%s<%s>", name, address)
			m.SetHeader("To", item)
		}
	}
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)

	d := gomail.NewDialer(e.Smtp, cast.ToInt(e.Port), e.Account, e.Password)
	if e.Ssl == true {
		d.SSL = true
	}
	return d.DialAndSend(m)
}

func NewEmail(smtp string, port uint64, ssl bool, account string, password string) *Email {
	email := &Email{
		Smtp:     smtp,
		Port:     port,
		Ssl:      ssl,
		Account:  account,
		Password: password,
	}
	return email
}
