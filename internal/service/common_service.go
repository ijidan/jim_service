package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"jim_service/config"
	"jim_service/internal/jim_proto/proto_build"
	"jim_service/pkg"
)

type CommonService struct {
	BasicService
	proto_build.UnimplementedCommonServiceServer
}


func (s *CommonService) UploadImage(stream proto_build.CommonService_UploadImageServer) error  {
	fileUpload:=pkg.NewFileUpload()
	defer fileUpload.Close()

	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return status.Error(codes.Internal, err.Error())
		}
		content := req.GetContent()
		_,err1:=fileUpload.WriteContent(content)
		if err1 != nil {
			return status.Error(codes.Internal, "write file failed")
		}
	}

	fileName:=fileUpload.GetFileName()
	img,err2:=pkg.SmMsUploadImage(pkg.Conf.Smms.User,pkg.Conf.Smms.Password,fileName)
	if err2!=nil{
		return status.Error(codes.Internal,err2.Error())
	}
	rsp := &proto_build.UploadImageResponse{Url: img.Url}
	err3 := stream.SendAndClose(rsp)
	if err3 != nil {
		return status.Error(codes.Internal, err3.Error())
	}
	return nil
}

func NewCommonService(config *config.Config) *CommonService {
	instance := &CommonService{BasicService: BasicService{
		Name: "service_common",
		Host: config.Rpc.Host,
		Port: config.Rpc.Port,
		Ttl:  config.Rpc.Ttl,
	}}
	return instance
}
