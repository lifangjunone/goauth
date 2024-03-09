package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	pb "goauth/api/goauth/v1"
	userv1 "goauth/api/user/v1"
)

type UserInfo struct {
	Id       int64
	Username string
	Password string
	Age      int64
	Gender   string
}

type LoginInfo struct {
	Username string
	Password string
}

type AuthUseCase struct {
	log *log.Helper
	uc  userv1.UserClient
}

func NewAuthUseCase(uc userv1.UserClient, logger log.Logger) *AuthUseCase {
	return &AuthUseCase{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func NewUserClient(r registry.Discovery) userv1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///user.rpc"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return userv1.NewUserClient(conn)
}

func (s *AuthUseCase) RegisterBiz(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	user, err := s.uc.GetUser(ctx, &userv1.GetUserRequest{Id: 1})
	if err != nil {
		panic("call failed")
	}
	return &pb.RegisterReply{
		Id:       user.Id,
		Username: user.Username,
		Age:      user.Age,
		Gender:   user.Gender,
	}, nil
}
func (s *AuthUseCase) LoginBiz(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	return &pb.LoginReply{}, nil
}
