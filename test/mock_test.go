package test

import (
	"jim_service/test/mocks"
	"strings"
	"testing"
)

func TestStringer(t *testing.T)  {
	mock:=&mocks.Stringer{}
	mock.On("String").Return(func(string2 string) string{
		return strings.ToUpper(string2)
	})
	str:=mock.String("jidan")
	t.Logf(str)
}

