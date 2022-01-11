package model

import (
	"gorm.io/gorm"
	"jim_service/internal/gen_model"
)

type User gen_model.User

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	return nil
}

func (u *User) AfterUpdate(tx *gorm.DB) (err error) {
	return nil
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	return nil
}

func (u *User) AfterSave(tx *gorm.DB) (err error) {
	return nil
}

func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	return nil
}

func (u *User) AfterDelete(tx *gorm.DB) (err error) {
	return nil
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	return nil
}
