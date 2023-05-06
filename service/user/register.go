package user

import (
	"fmt"
	"math/rand"

	"github.com/Hui4401/gopkg/errors"

	"github.com/Hui4401/qa/constdef"
	"github.com/Hui4401/qa/model"
	sqlModel "github.com/Hui4401/qa/storage/mysql/model"
)

func Register(req *model.UserRegisterRequest) (*model.UserRegisterResponse, error) {
	// 表单验证
	if err := registerValid(req); err != nil {
		return nil, err
	}
	user := &sqlModel.User{
		Username: req.Username,
		UserProfile: &sqlModel.UserProfile{
			Nickname: req.Username,
			Avatar:   fmt.Sprintf("https://images.nowcoder.com/head/%dt.png", rand.Intn(1000)),
		},
	}
	// 加密密码
	if err := user.SetPassword(req.Password); err != nil {
		return nil, err
	}
	// 创建用户
	ud := sqlModel.NewUserDao()
	if err := ud.CreateUser(user); err != nil {
		return nil, err
	}
	// 注册后默认直接登录
	token, err := generateToken(user.ID)
	if err != nil {
		return nil, nil
	}

	return &model.UserRegisterResponse{
		Token: token,
	}, nil
}

func registerValid(req *model.UserRegisterRequest) error {
	// 两次输入密码不一致
	if req.PasswordConfirm != req.Password {
		return errors.NewCodeError(constdef.CodePasswordConfirmError)
	}
	// 用户名已存在
	ud := sqlModel.NewUserDao()
	user, err := ud.GetUserByUsername(req.Username)
	if err != nil {
		return err
	}
	if user != nil {
		return errors.NewCodeError(constdef.CodeUserExist)
	}

	return nil
}
