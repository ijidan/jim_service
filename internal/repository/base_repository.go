package repository

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"jim_service/internal/jim_proto/proto_build"
	"jim_service/pkg"
)

const (
	ReviewInit = 0 //未提交审核
	ReviewTo   = 1 //已提交，待审核
	ReviewPass = 2 //审核通过
	ReviewFail = 3 //审核未通过
)

const (
	NotEnable = 0 //未启用
	Enable    = 1 //启用
)

var ReviewStatusDbToProto = map[int32]proto_build.ReviewStatus{
	ReviewInit: proto_build.ReviewStatus_ReviewInit,
	ReviewTo:   proto_build.ReviewStatus_ReviewTo,
	ReviewPass: proto_build.ReviewStatus_ReviewPass,
	ReviewFail: proto_build.ReviewStatus_ReviewFail,
}
var ReviewStatusProtoToDb = map[proto_build.ReviewStatus]int32{
	proto_build.ReviewStatus_ReviewInit: ReviewInit,
	proto_build.ReviewStatus_ReviewTo:   ReviewTo,
	proto_build.ReviewStatus_ReviewPass: ReviewPass,
	proto_build.ReviewStatus_ReviewFail: ReviewFail,
}

var EnableDbToProto = map[int32]proto_build.IsEnable{
	NotEnable: proto_build.IsEnable_NotEnable,
	Enable:    proto_build.IsEnable_Enable,
}

var EnableProtoToDb = map[proto_build.IsEnable]int32{
	proto_build.IsEnable_NotEnable: NotEnable,
	proto_build.IsEnable_Enable:    Enable,
}

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
	if err == nil {
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
