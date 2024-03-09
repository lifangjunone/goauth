package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "goauth/api/user/v1"
)

type User struct {
	Id       int64
	Username string
	Password string
	Age      int64
	Gender   string
}

type UserReply struct {
	Id       int64
	Username string
	Age      int64
	Gender   string
}

type UserDelReply struct {
	Ok bool
}

type UserRepo interface {
	Create(context.Context, *User) (*UserReply, error)
	Update(context.Context, *User) (*UserReply, error)
	Delete(context.Context, int64) (*UserDelReply, error)
	Get(context.Context, int64) (*UserReply, error)
	List(context.Context) ([]*UserReply, error)
}

func NewUserUseCase(ur UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		ur:  ur,
		log: log.NewHelper(logger),
	}
}

// UserUseCase Create UserRepo object
type UserUseCase struct {
	ur  UserRepo
	log *log.Helper
}

func (s *UserUseCase) CreateUserBiz(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	return &pb.CreateUserReply{}, nil
}
func (s *UserUseCase) UpdateUserBiz(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}
func (s *UserUseCase) DeleteUserBiz(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}
func (s *UserUseCase) GetUserBiz(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	user, err := s.ur.Get(ctx, 1)
	if err != nil {
		panic("call failed")
	}
	return &pb.GetUserReply{
		Id:       user.Id,
		Username: user.Username,
		Age:      user.Age,
		Gender:   user.Gender,
	}, nil
}
func (s *UserUseCase) ListUserBiz(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}
