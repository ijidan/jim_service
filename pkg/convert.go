package pkg

import (
	"fmt"
	"strings"
)

func ConvertUintSlice2SepString(data []uint64) string  {
	return strings.Replace(strings.Trim(fmt.Sprint(data), "[]"), " ", ",", -1)
}