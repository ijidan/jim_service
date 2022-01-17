package repository

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"jim_service/internal/jim_proto/proto_build"
	"jim_service/internal/model"
)

func ConvertGroupToProtoGroup(group model.Group) *proto_build.Group {
	protoGroup := &proto_build.Group{
		Id:           group.ID,
		Name:         group.Name,
		Introduction: group.Introduction,
		AvatarUrl:    group.AtavarURL,
		Extra:        group.Extra,
		UserId:       group.UserID,
		CreatedAt:    timestamppb.New(group.CreatedAt),
		UpdatedAt:    timestamppb.New(group.UpdatedAt),
		DeletedAt:    timestamppb.New(group.DeletedAt.Time),
	}
	return protoGroup
}

func CreateGroup(db *gorm.DB, userId uint64, name string, introduction string, avatarUrl string) (*model.Group, error) {
	group := &model.Group{
		Name:         name,
		Introduction: introduction,
		Extra:        "",
		AtavarURL:    avatarUrl,
		UserID:       userId,
	}
	err := db.Create(group).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return group, nil
}

func UpdateGroup(db *gorm.DB, id uint64, name string, introduction string, avatarUrl string) error {
	var group model.Group
	err1 := db.Where("id=?", id).First(&group).Error
	if err1 != nil {
		return ConvertModelQueryErrorToGrpcStatusError(err1)
	}
	upData := map[string]interface{}{"name": name, "introduction": introduction, "atavar_url": avatarUrl}
	err2 := db.Where("id=?", id).UpdateColumns(upData).Error
	if err2 != nil {
		return status.Error(codes.Internal, err2.Error())
	}
	return nil
}

func GetProtoGroupByGroupId(db *gorm.DB, id uint64) (*proto_build.Group, error) {
	var group model.Group
	err := db.Where("id=?", id).First(&group).Error
	if err != nil {
		return nil,ConvertModelQueryErrorToGrpcStatusError(err)
	}
	protoGroup := ConvertGroupToProtoGroup(group)
	return protoGroup, nil
}

func GetProtoGroup(db *gorm.DB, id uint64) (*proto_build.Group, error) {
	var group model.Group
	err := db.Where("id=?", id).First(&group).Error
	if err != nil {
		return nil, ConvertModelQueryErrorToGrpcStatusError(err)
	}
	protoGroup := ConvertGroupToProtoGroup(group)
	return protoGroup, nil
}

func QueryProtoGroup(db *gorm.DB, keyword string, page uint64, pageSize uint64) ([]*proto_build.Group, *proto_build.Pager, error) {
	where := "name like ?"
	whereValue := []interface{}{"%" + keyword + "%"}

	var group model.Group
	var rows []model.Group
	order := "id desc"
	pager, err := model.QueryPager(db, &group, &rows, order, page, pageSize, where, whereValue)
	if err != nil {
		return nil, nil, ConvertModelQueryErrorToGrpcStatusError(err)
	}
	var protoGroupList []*proto_build.Group
	if len(rows) > 0 {
		for _, v := range rows {
			protoGroup := ConvertGroupToProtoGroup(v)
			protoGroupList = append(protoGroupList, protoGroup)
		}
	}
	protoPager := ConvertPagerToProtoPager(pager)
	return protoGroupList, protoPager, nil
}

func DeleteGroup(db *gorm.DB, id uint64) error {
	err:=db.Transaction(func(tx *gorm.DB) error {
		var group model.Group
		err1 := db.Where("id=?", id).Delete(&group).Error
		if err1 != nil {
			return ConvertModelQueryErrorToGrpcStatusError(err1)
		}
		var groupUser model.GroupUser
		err2:=db.Where("group_id=?",id).Delete(&groupUser).Error
		if err2!=nil{
			return ConvertModelQueryErrorToGrpcStatusError(err2)
		}
		return nil
	})
	if err != nil {
		return ConvertModelQueryErrorToGrpcStatusError(err)
	}
	return nil
}

func JoinGroup(db *gorm.DB, id uint64, userId uint64) error {
	var group model.Group
	err := db.Where("id=?", id).First(&group).Error
	if err != nil {
		return ConvertModelQueryErrorToGrpcStatusError(err)
	}
	var user model.User
	err2 := db.Where("id=?", id).First(&user).Error
	if err2 != nil {
		return ConvertModelQueryErrorToGrpcStatusError(err)
	}
	var groupUser model.GroupUser
	err3 := db.Where("group_id=? and user_id=?", id, userId).First(&groupUser).Error
	if err3 == nil {
		return errors.New("record exists")
	}
	if errors.Is(err3, gorm.ErrRecordNotFound) {
		groupUser := &model.GroupUser{
			ID:           0,
			GroupID:      id,
			UserID:       userId,
			UserShowName: user.Nickname,
		}
		err4 := db.Create(groupUser).Error
		if err4 != nil {
			return status.Error(codes.Internal, err4.Error())
		}
		return nil
	}
	return status.Error(codes.Internal, err3.Error())
}

func QuitGroup(db *gorm.DB, id uint64, userId uint64) error {
	var groupUser model.GroupUser
	err := db.Where("group_id=? and user_id=?", id, userId).First(&groupUser).Error
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}
