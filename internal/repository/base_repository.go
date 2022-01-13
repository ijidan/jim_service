package repository

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"jim_service/internal/jim_proto/proto_build"
	"jim_service/pkg"
)

func ConvertPagerToProtoPager(pager *pkg.Pager) *proto_build.Pager {
	protoPager := &proto_build.Pager{
		Page:       pager.Page,
		PageZie:    pager.PageSize,
		TotalRows:  pager.TotalRows,
		TotalPages: pager.TotalPages,
	}
	return protoPager
}

func ConvertModelQueryErrorToGrpcStatusError(err error) error {
	if err==nil{
		return nil
	}
	var code codes.Code
	if errors.Is(err, gorm.ErrRecordNotFound) {
		code = codes.NotFound
	} else {
		code = codes.Internal
	}
	return status.Error(code, err.Error())
	
}
