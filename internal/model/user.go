package model

import (
	"gorm.io/gorm"
	"jim_service/internal/gen_model"
)

type User gen_model.User

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	return
}