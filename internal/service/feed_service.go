package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"jim_service/internal/jim_proto/proto_build"
	"jim_service/internal/model"
	"jim_service/internal/repository"
	"jim_service/pkg"
)

type FeedService struct {
	BasicService
	proto_build.UnimplementedFeedServiceServer
}

func (s *FeedService) FeedCreate(c context.Context, req *proto_build.FeedCreateRequest) (*proto_build.FeedCreateResponse, error) {
	userId:=s.GetLoginUserId()
	var feed *model.Feed
	var err error
	switch req.Type {
	case proto_build.FeedType_Txt:
		feed,err=repository.CreateTextFeed(pkg.Db, userId, req.Content)
	case proto_build.FeedType_Image:
		feed,err=repository.CreateImageFeed(pkg.Db, userId, req.Content,req.Resource)
	case proto_build.FeedType_Video:
		feed,err=repository.CreateVideoFeed(pkg.Db, userId, req.Content,req.Resource)
	}
	if err!=nil{
		return nil,status.Error(codes.Internal, err.Error())
	}
	protoFeed:=repository.ConvertFeedToProtoFeed(*feed)
	rsp:=&proto_build.FeedCreateResponse{Feed: protoFeed}
	return rsp,nil
}

func (s *FeedService) FeedEdit(c context.Context, req *proto_build.FeedEditRequest) (*proto_build.FeedEditResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method FeedEdit not implemented")
}

func (s *FeedService) FeedLike(c context.Context, req *proto_build.FeedLikeRequest) (*proto_build.FeedLikeResponse, error) {
	userId:=s.GetLoginUserId()
	err:=repository.LikeFeed(pkg.Db,req.Id,userId)
	if err!=nil{
		return nil,status.Error(codes.Internal, err.Error())
	}
	rsp:=&proto_build.FeedLikeResponse{}
	return rsp, nil
}

func (s *FeedService) FeedUnLike(c context.Context, req *proto_build.FeedUnLikeRequest) (*proto_build.FeedUnLikeResponse, error) {
	userId:=s.GetLoginUserId()
	err:=repository.UnLikeFeed(pkg.Db,req.Id,userId)
	if err!=nil{
		return nil,status.Error(codes.Internal,err.Error())
	}
	rsp:=&proto_build.FeedUnLikeResponse{}
	return rsp, nil
}

func (s *FeedService) FeedGet(c context.Context, req *proto_build.FeedGetRequest) (*proto_build.FeedGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedGet not implemented")
}

func (s *FeedService) FeedDelete(c context.Context, req *proto_build.FeedDeleteRequest) (*proto_build.FeedDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedDelete not implemented")
}

func (s *FeedService) FeedOwn(c context.Context, req *proto_build.FeedOwnRequest) (*proto_build.FeedOwnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedOwn not implemented")
}

func (s *FeedService) FeedSearch(c context.Context, req *proto_build.FeedSearchRequest) (*proto_build.FeedSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedSearch not implemented")
}

func (s *FeedService) FeedFollow(c context.Context, req *proto_build.FeedFollowRequest) (*proto_build.FeedFollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedFollow not implemented")
}
