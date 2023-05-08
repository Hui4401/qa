package user

import (
	"context"
	"time"

	"github.com/Hui4401/gopkg/errors"
	"github.com/dgrijalva/jwt-go"

	"github.com/Hui4401/qa/constdef"
	"github.com/Hui4401/qa/middleware/auth"
	"github.com/Hui4401/qa/model"
	"github.com/Hui4401/qa/storage"
)

// Login 用户登录函数
func Login(ctx context.Context, req *model.UserLoginRequest) (*model.UserLoginResponse, error) {
	if !storage.CheckUserPassword(ctx, req.Username, req.Password) {
		return nil, errors.NewCodeError(constdef.CodePasswordError)
	}

	user, err := storage.GetUserInfoByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	token, err := generateToken(user.UserID)
	if err != nil {
		return nil, err
	}

	return &model.UserLoginResponse{
		Token: token,
	}, nil
}

func generateToken(userID uint) (string, error) {
	claim := auth.JwtClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(constdef.UserTokenExpiredTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID: userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString([]byte(auth.JwtSecretKey))
}
