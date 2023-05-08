package user

import (
	"github.com/Hui4401/gopkg/errors"
	"github.com/Hui4401/gopkg/logs"
	"github.com/gin-gonic/gin"

	"github.com/Hui4401/qa/constdef"
	"github.com/Hui4401/qa/model"
	"github.com/Hui4401/qa/service/user"
	"github.com/Hui4401/qa/storage"
)

// Register 用户注册
func Register(ctx *gin.Context) (interface{}, error) {
	req := model.UserRegisterRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		return nil, errors.NewCodeError(constdef.CodeParam)
	}

	res, err := user.Register(ctx, &req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Login 用户登录
func Login(ctx *gin.Context) (interface{}, error) {
	req := model.UserLoginRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		return nil, errors.NewCodeError(constdef.CodeParam)
	}

	res, err := user.Login(ctx, &req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Logout 退出登录
func Logout(ctx *gin.Context) (interface{}, error) {
	token, ok := ctx.Get(constdef.CtxUserToken)
	if !ok {
		return nil, errors.NewCodeError(constdef.CodeTokenNotFound)
	}
	if err := storage.BanUserToken(ctx, token.(string), constdef.UserTokenExpiredTime); err != nil {
		return nil, err
	}

	return nil, nil
}

// Profile 查看当前登录用户个人资料
func Profile(ctx *gin.Context) (interface{}, error) {
	v, ok := ctx.Get(constdef.CtxUserID)
	if !ok {
		return nil, errors.NewCodeError(constdef.CodeTokenExpired)
	}
	userID, ok := v.(uint)
	if !ok {
		logs.CtxErrorKvs(ctx, "userID covert fail, v", v)
		return nil, errors.NewCodeError(constdef.CodeUnknown)
	}

	userInfo, err := storage.GetUserInfoByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if userInfo == nil {
		return nil, errors.NewCodeError(constdef.CodeUserNotExist)
	}

	return &model.UserProfileResponse{
		Username:    userInfo.Username,
		Email:       userInfo.Email,
		Nickname:    userInfo.Nickname,
		Avatar:      userInfo.Avatar,
		Description: userInfo.Description,
	}, err
}
