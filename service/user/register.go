package user

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/Hui4401/gopkg/errors"

	"github.com/Hui4401/qa/constdef"
	"github.com/Hui4401/qa/model"
	"github.com/Hui4401/qa/storage"
)

func Register(ctx context.Context, req *model.UserRegisterRequest) (*model.UserRegisterResponse, error) {
	// 表单验证
	if err := registerValid(ctx, req); err != nil {
		return nil, err
	}

	createUserReq := &storage.CreateUserReq{
		Username: req.Username,
		Password: req.Password,
		Nickname: req.Username,
		Avatar:   fmt.Sprintf("https://images.nowcoder.com/head/%dt.png", rand.Intn(1000)),
	}
	newUser, err := storage.CreateUser(ctx, createUserReq)
	if err != nil {
		return nil, err
	}

	// 注册后默认直接登录
	token, err := generateToken(newUser.UserID)
	if err != nil {
		return nil, nil
	}

	return &model.UserRegisterResponse{
		Token: token,
	}, nil
}

func registerValid(ctx context.Context, req *model.UserRegisterRequest) error {
	// 两次输入密码不一致
	if req.PasswordConfirm != req.Password {
		return errors.NewCodeError(constdef.CodePasswordConfirmError)
	}
	// 用户已存在
	if isUserExist := storage.IsUserExistByUsername(ctx, req.Username, true); isUserExist {
		return errors.NewCodeError(constdef.CodeUserExist)
	}

	return nil
}
