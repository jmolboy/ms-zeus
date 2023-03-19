package service

import (
	"context"

	"github.com/jmolboy/ms-adm-pkg/pmtjwt"
	pb "github.com/jmolboy/ms-zeus/api/zeus/v1"
)

type OpUserService struct {
	pb.UnimplementedOpUserServer
}

func NewOpUserService() *OpUserService {
	return &OpUserService{}
}

func (s *OpUserService) Current(ctx context.Context, req *pb.CurrentRequest) (rply *pb.CurrentReply, err error) {
	// 获取当前登录用户
	jwtUsr, err := pmtjwt.FromAuthContext(ctx)
	if err != nil {
		return
	}

	rply = &pb.CurrentReply{
		Data: &pb.CurrentUser{
			Id:       jwtUsr.Id,
			Uid:      jwtUsr.Id,
			UserId:   jwtUsr.UserId,
			UserName: jwtUsr.UserName,
			RoleId:   jwtUsr.RoleId,
		},
	}
	return
}

func (s *OpUserService) Logout(context.Context, *pb.LogOutRequest) (*pb.LogOutReply, error) {
	// 获取
	return &pb.LogOutReply{
		Data: "",
	}, nil
}
