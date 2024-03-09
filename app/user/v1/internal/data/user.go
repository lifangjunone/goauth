package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"goauth/app/user/v1/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u userRepo) Create(ctx context.Context, user *biz.User) (*biz.UserReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepo) Update(ctx context.Context, user *biz.User) (*biz.UserReply, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepo) Delete(ctx context.Context, i int64) (*biz.UserDelReply, error) {
	//TODO implement me
	return &biz.UserDelReply{
		Ok: true,
	}, nil
}

func (u userRepo) Get(ctx context.Context, i int64) (*biz.UserReply, error) {
	return &biz.UserReply{
		Id:       i,
		Username: "张三",
		Age:      18,
		Gender:   "男",
	}, nil
}

func (u userRepo) List(ctx context.Context) ([]*biz.UserReply, error) {
	users := make([]*biz.UserReply, 0)
	users = append(users, &biz.UserReply{
		Id:       1,
		Username: "张三",
		Age:      18,
		Gender:   "男",
	})
	users = append(users, &biz.UserReply{
		Id:       2,
		Username: "高圆圆",
		Age:      28,
		Gender:   "女",
	})
	return users, nil
}
