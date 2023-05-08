package storage

import (
	"context"
	"time"

	sqlModel "github.com/Hui4401/qa/storage/mysql/model"
	redisModel "github.com/Hui4401/qa/storage/redis/model"
)

type CreateUserReq struct {
	Username    string // 用户名
	Password    string // 密码（明文）
	Nickname    string // 昵称
	Email       string // 邮箱
	Avatar      string // 头像
	Description string // 个人描述
}

type UserInfo struct {
	UserID      uint      // 用户ID
	Username    string    // 用户名
	Email       string    // 邮箱
	Nickname    string    // 昵称
	Avatar      string    // 头像
	Description string    // 个人描述
	CreatedAt   time.Time // 创建时间
}

func CreateUser(ctx context.Context, userReq *CreateUserReq) (*UserInfo, error) {
	user := &sqlModel.User{
		Username: userReq.Username,
		Nickname: userReq.Nickname,
		Avatar:   userReq.Avatar,
	}
	// 加密密码
	if err := user.SetPassword(userReq.Password); err != nil {
		return nil, err
	}
	// 创建用户
	ud := sqlModel.NewUserDao()
	if err := ud.CreateUser(user); err != nil {
		return nil, err
	}

	return &UserInfo{
		UserID:      user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Nickname:    user.Nickname,
		Avatar:      user.Avatar,
		Description: user.Description,
		CreatedAt:   user.CreatedAt,
	}, nil
}

// IsUserExistByUsername 判断用户是否存在，发生错误时由调用方决定结果
func IsUserExistByUsername(ctx context.Context, username string, defaultVal bool) bool {
	user, err := GetUserInfoByUsername(ctx, username)
	if err != nil {
		return defaultVal
	}

	return user != nil
}

// GetUserInfoByID 获取用户信息，用户不存在时返回 nil，nil
func GetUserInfoByID(ctx context.Context, userID uint) (*UserInfo, error) {
	ud := sqlModel.NewUserDao()
	user, err := ud.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}

	return &UserInfo{
		UserID:      user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Nickname:    user.Nickname,
		Avatar:      user.Avatar,
		Description: user.Description,
		CreatedAt:   user.CreatedAt,
	}, nil
}

// GetUserInfoByUsername 获取用户信息，用户不存在时返回 nil，nil
func GetUserInfoByUsername(ctx context.Context, username string) (*UserInfo, error) {
	ud := sqlModel.NewUserDao()
	user, err := ud.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}

	return &UserInfo{
		UserID:      user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Nickname:    user.Nickname,
		Avatar:      user.Avatar,
		Description: user.Description,
		CreatedAt:   user.CreatedAt,
	}, nil
}

// CheckUserPassword 验证密码，发生内部错误默认不通过检查
func CheckUserPassword(ctx context.Context, username string, password string) bool {
	ud := sqlModel.NewUserDao()
	user, err := ud.GetUserByUsername(username)
	if err != nil || user == nil {
		return false
	}

	return user.CheckPassword(password)
}

// BanUserToken 下线用户token
func BanUserToken(ctx context.Context, token string, expireTime time.Duration) error {
	jd := redisModel.NewJwtDao()
	return jd.BanToken(ctx, token, expireTime)
}

// CheckUserToken 判断用户token是否有效
func CheckUserToken(ctx context.Context, token string) bool {
	jd := redisModel.NewJwtDao()
	return !jd.IsBanedToken(ctx, token)
}
