package repository

import (
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
