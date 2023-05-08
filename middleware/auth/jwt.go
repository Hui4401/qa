package auth

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/Hui4401/qa/constdef"
	"github.com/Hui4401/qa/storage"
	"github.com/Hui4401/qa/util"
)

const (
	JwtSecretKey = "a random key"
)

type JwtClaim struct {
	jwt.StandardClaims
	UserID uint
}

// JwtAuthRequired 通过jwt秘钥来验证用户身份
func JwtAuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从请求头获得token
		userToken := ctx.Request.Header.Get("Authorization")

		// 判断请求头中是否有token
		if userToken == "" {
			ctx.JSON(http.StatusOK, util.ErrorResponseByCode(constdef.CodeTokenNotFound))
			ctx.Abort()
			return
		}

		// 解码token值
		token, err := jwt.ParseWithClaims(userToken, &JwtClaim{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(JwtSecretKey), nil
		})
		if err != nil || token.Valid != true {
			// 过期或者token不正确
			ctx.JSON(http.StatusOK, util.ErrorResponseByCode(constdef.CodeTokenExpired))
			ctx.Abort()
			return
		}

		// 判断token是否有效
		if !storage.CheckUserToken(ctx, token.Raw) {
			ctx.JSON(http.StatusOK, util.ErrorResponseByCode(constdef.CodeTokenExpired))
			ctx.Abort()
			return
		}

		// context保存用户信息
		if jwtStruct, ok := token.Claims.(*JwtClaim); ok {
			ctx.Set(constdef.CtxUserID, jwtStruct.UserID)
		}
		ctx.Set(constdef.CtxUserToken, token.Raw)
	}
}
