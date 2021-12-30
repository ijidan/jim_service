package repository

import (
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"jim_service/internal/jim_proto/proto_build"
	"jim_service/internal/model_gormt"
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

func ConvertUserToProtoUser(user model_gormt.User) *proto_build.User {
	protoUser := &proto_build.User{
		Id:        user.ID,
		Nickname:  user.Nickname,
		Gender:    GenderDbToProto[user.Gender],
		AvatarUrl: user.AvatarURL,
		Extra:     user.Extra,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
		DeletedAt: timestamppb.New(user.DeletedAt),
	}
	return protoUser
}

func CreateUser(db *gorm.DB, nickName string, password string, gender string, avatarUrl string, extra string) *model_gormt.User {
	passwordHash, _ := pkg.HashPassword(password)
	userMgr := model_gormt.UserMgr(db)
	user := &model_gormt.User{
		Nickname:  nickName,
		Password:  passwordHash,
		Gender:    gender,
		AvatarURL: avatarUrl,
		Extra:     extra,
	}
	userMgr.Create(user)
	return user
}
func GetProtoUserByUserId(db *gorm.DB, userId uint64) (*proto_build.User,error) {
	userMgr := model_gormt.UserMgr(db)
	user, err := userMgr.GetFromID(userId)
	if err != nil {
		return nil,errors.New("not found")
	}
	protoUser := ConvertUserToProtoUser(user)
	return protoUser,nil
}

func QueryProtoUser(db *gorm.DB, keyword string, page uint64, pageSize uint64) ([]*proto_build.User, *proto_build.Pager, error) {
	where := model_gormt.UserColumns.Nickname + " like ?"
	whereValue := []interface{}{"%" + keyword + "%"}

	var model model_gormt.User
	var rows []model_gormt.User
	order:=model_gormt.UserColumns.ID+" desc"
	pager, _ := QueryPager(db, &model, &rows, order, page, pageSize, where, whereValue)
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
