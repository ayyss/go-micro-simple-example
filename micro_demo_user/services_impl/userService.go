package services_impl

import (
	"context"
	"micro_demo_user/services"
)

var userMap = make(map[string]string)

type UserService struct {
}

func (*UserService) UserLogin(ctx context.Context, req *services.UserRequest, resp *services.UserDetailResponse) error {
	if _, ok := userMap[req.UserName]; !ok {
		resp.Code = 1
		resp.Msg = "用户不存在"
	} else if userMap[req.UserName] == req.Password {
		resp.Code = 0
		resp.UserDetail = &services.UserModel{UserName: req.UserName, Password: req.Password}
		resp.Msg = "登录成功"
	} else {
		resp.Code = 1
		resp.Msg = "密码错误"
	}
	return nil
}

func (*UserService) UserRegister(ctx context.Context, req *services.UserRequest, resp *services.UserDetailResponse) error {
	if _, ok := userMap[req.UserName]; ok {
		resp.Code = 1
		resp.Msg = "用户名已存在"
	} else {
		userMap[req.UserName] = req.Password
		resp.Code = 0
		resp.UserDetail = &services.UserModel{UserName: req.UserName, Password: req.Password}
		resp.Msg = "注册成功"
	}
	return nil
}
