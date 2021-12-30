package repository

import (
	"github.com/spf13/cast"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"jim_service/internal/jim_proto/proto_build"
	"jim_service/pkg"
	"math"
)

func Pager(page uint64, pageSize uint64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(cast.ToInt(offset)).Limit(cast.ToInt(pageSize))
	}
}

func QueryPager(db *gorm.DB, model interface{}, rows interface{}, order string, page uint64, pageSize uint64, where string, whereValue ...interface{}) (*pkg.Pager, error) {
	var totalRows int64
	var totalPages int64

	err:=db.Scopes(Pager(page, pageSize)).Where(where, whereValue).Order(order).Find(rows).Error
	if err!=nil{
		return nil,status.Error(codes.Internal,err.Error())
	}
	err1:=db.Model(&model).Where(where, whereValue).Count(&totalRows).Error
	if err1!=nil{
		return nil,status.Error(codes.Internal,err.Error())
	}
	totalPages = cast.ToInt64(math.Ceil(cast.ToFloat64(totalRows) / cast.ToFloat64(pageSize)))
	pager := &pkg.Pager{
		Page:       page,
		PageSize:   pageSize,
		TotalRows:  cast.ToUint64(totalRows),
		TotalPages: cast.ToUint64(totalPages),
		Rows:       rows,
	}
	return pager, nil

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
