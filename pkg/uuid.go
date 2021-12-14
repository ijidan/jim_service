package pkg

import (
	"github.com/gofrs/uuid"
)

func GetUUId() string {
	u2, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return u2.String()
}
