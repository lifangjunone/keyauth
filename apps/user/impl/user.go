package impl

import (
	"context"
	"github.com/lifangjunone/keyauth/apps/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *impl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.User, error) {
	ins := user.NewUser(req)
	err := save(i.col, ctx, ins)
	if err != nil {
		return nil, err
	}
	return ins, err
}

func (i *impl) QueryUser(ctx context.Context, req *user.QueryUserRequest) (*user.UserSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryUser not implemented")
}

func (i *impl) DescribeUser(ctx context.Context, req *user.DescribeUserRequest) (*user.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeUser not implemented")
}

func (i *impl) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}

func (i *impl) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
