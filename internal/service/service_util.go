package service

import (
	"jim_service/global"
	"jim_service/pkg"
)

func Register()  {
	endPoints:=global.Config.Etcd.EndPoints
	client:=pkg.NewClientV3(global.Config.Etcd.EndPoints,5)
}

func UnRegister()  {
	
}



