package user

import (
	"github.com/Hui4401/gopkg/errors"
	"github.com/Hui4401/gopkg/logs"
	"github.com/gin-gonic/gin"

	"github.com/Hui4401/qa/model"
	"github.com/Hui4401/qa/service/user"
	sqlModel "github.com/Hui4401/qa/storage/mysql/model"
	redisModel "github.com/Hui4401/qa/storage/redis/model"
	"github.com/Hui4401/qa/util/error_code"
)

// Register 用户注册
func Register(ctx *gin.Context) (interface{}, error) {
	req := model.UserRegisterRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		return nil, errors.NewCodeError(error_code.CodeParam)
	}

	res, err := user.Register(&req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Login 用户登录
func Login(ctx *gin.Context) (interface{}, error) {
	req := model.UserLoginRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		return nil, errors.NewCodeError(error_code.CodeParam)
	}

	res, err := user.Login(&req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Logout 退出登录
func Logout(ctx *gin.Context) (interface{}, error) {
	token, ok := ctx.Get("token")
	if !ok {
		return nil, errors.NewCodeError(error_code.CodeTokenNotFound)
	}
	jd := redisModel.NewJwtDao()
	if err := jd.BanToken(ctx, token.(string)); err != nil {
		return nil, err
	}

	return nil, nil
}

// Profile 查看当前登录用户个人基本信息
func Profile(ctx *gin.Context) (interface{}, error) {
	v, ok := ctx.Get("user_id")
	if !ok {
		return nil, errors.NewCodeError(error_code.CodeTokenExpired)
	}
	uid, ok := v.(uint)
	if !ok {
		logs.CtxErrorKvs(ctx, "uid covert fail, v", v)
		return nil, errors.NewCodeError(error_code.CodeUnknown)
	}
	userDao := sqlModel.NewUserDao()
	userProfile, err := userDao.GetUserProfileByID(uid)
	if err != nil {
		return nil, err
	}
	if userProfile == nil {
		return nil, errors.NewCodeError(error_code.CodeUserNotExist)
	}

	return &model.UserProfileResponse{
		Id:          userProfile.ID,
		Nickname:    userProfile.Nickname,
		Email:       userProfile.Email,
		Avatar:      userProfile.Avatar,
		Description: userProfile.Description,
	}, err
}
