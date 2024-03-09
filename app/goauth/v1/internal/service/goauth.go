package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"goauth/app/goauth/v1/internal/biz"

	pb "goauth/api/goauth/v1"
)

type GoAuthService struct {
	pb.UnimplementedGoAuthServer
	auth *biz.AuthUseCase
	log  *log.Helper
}

func NewGoAuthService(auth *biz.AuthUseCase, logger log.Logger) *GoAuthService {
	return &GoAuthService{
		auth: auth,
		log:  log.NewHelper(logger),
	}
}

func (s *GoAuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	user, err := s.auth.RegisterBiz(ctx, req)
	if err != nil {
		panic("call failed")
	}
	return &pb.RegisterReply{
		Id:       user.Id,
		Username: user.Username,
		Gender:   user.Gender,
		Age:      user.Age,
	}, nil
}
func (s *GoAuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	return &pb.LoginReply{}, nil
}
