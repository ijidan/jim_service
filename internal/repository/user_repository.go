package repository

import (
	"github.com/spf13/cast"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"jim_service/internal/jim_proto/proto_build"
	"jim_service/internal/model"
	"jim_service/pkg"
)

const GenderMale = "1"
const GenderFemale = "2"

var GenderDbToProto = map[string]proto_build.Gender{
	GenderMale:   proto_build.Gender_Male,
	GenderFemale: proto_build.Gender_Female,
}

var GenderProtoToDb = map[proto_build.Gender]string{
	proto_build.Gender_Male:   GenderMale,
	proto_build.Gender_Female: GenderFemale,
}

func ConvertUserToProtoUser(user model.User) *proto_build.User {
	protoUser := &proto_build.User{
		Id:        cast.ToUint64(user.ID),
		Nickname:  user.Nickname,
		Gender:    GenderDbToProto[user.Gender],
		AvatarUrl: user.AvatarURL,
		Extra:     user.Extra,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
		DeletedAt: timestamppb.New(user.DeletedAt.Time),
	}
	return protoUser
}

func CreateUser(db *gorm.DB, nickName string, password string, gender string, avatarUrl string, extra string) (*model.User, error) {
	passwordHash, err := pkg.HashPassword(password)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	user := &model.User{
		Nickname:  nickName,
		Password:  passwordHash,
		Gender:    gender,
		AvatarURL: avatarUrl,
		Extra:     extra,
	}
	err1 := db.Create(user).Error
	if err1 != nil {
		return nil, status.Error(codes.Internal, err1.Error())
	}
	return user, nil
}

func GetProtoUser(db *gorm.DB,nickName string,password string)(*proto_build.User,error)  {
	passwordHash,err:=pkg.HashPassword(password)
	if err!=nil{
		return nil,status.Error(codes.Internal,err.Error())
	}
	var user model.User
	err1:=db.Where("nickname=? and password=?",nickName,passwordHash).First(&user).Error
	if err1!=nil{
		return nil,ConvertModelQueryErrorToGrpcStatusError(err1)
	}
	protoUser:=ConvertUserToProtoUser(user)
	return protoUser,nil

}
func GetProtoUserByUserId(db *gorm.DB, userId uint64) (*proto_build.User, error) {
	var user model.User
	err := db.Where("id=?", userId).First(&user).Error
	if err != nil {
		return nil,ConvertModelQueryErrorToGrpcStatusError(err)
	}
	protoUser := ConvertUserToProtoUser(user)
	return protoUser, nil
}

func QueryProtoUser(db *gorm.DB, keyword string, page uint64, pageSize uint64) ([]*proto_build.User, *proto_build.Pager, error) {
	where := "nickname like ?"
	whereValue := []interface{}{"%" + keyword + "%"}

	var user model.User
	var rows []model.User
	order :=  "id desc"
	pager, err := model.QueryPager(db, &user, &rows, order, page, pageSize, where, whereValue)
	if err!=nil{
		return nil, nil, ConvertModelQueryErrorToGrpcStatusError(err)
	}
	var protoUserList []*proto_build.User
	if len(rows) > 0 {
		for _, v := range rows {
			protoUser := ConvertUserToProtoUser(v)
			protoUserList = append(protoUserList, protoUser)
		}
	}
	protoPager := ConvertPagerToProtoPager(pager)
	return protoUserList, protoPager, nil
}

func UpdateUserPassword(db *gorm.DB,userId uint64,password string) error {
	err:=db.Model(&model.User{}).Where("id=?",userId).UpdateColumn("password",password).Error
	return ConvertModelQueryErrorToGrpcStatusError(err)
}

func UpdateUserAvatar(db *gorm.DB,userId uint64,avatarUrl string)  error{
	err:=db.Model(&model.User{}).Where("id=?",userId).UpdateColumn("avatar_url",avatarUrl).Error
	return ConvertModelQueryErrorToGrpcStatusError(err)
}

