package main

import (
	"os"

	"github.com/Hui4401/gopkg/errors"
	"github.com/Hui4401/gopkg/logs"
	"github.com/gin-gonic/gin"

	"github.com/Hui4401/qa/conf"
	"github.com/Hui4401/qa/handler"
	"github.com/Hui4401/qa/handler/user"
	"github.com/Hui4401/qa/middleware/auth"
	"github.com/Hui4401/qa/middleware/wrapper"
	"github.com/Hui4401/qa/storage/mysql"
	sqlModel "github.com/Hui4401/qa/storage/mysql/model"
	"github.com/Hui4401/qa/storage/redis"
	"github.com/Hui4401/qa/util/error_code"
)

func main() {
	conf.Init()
	mysql.InitMySQL(os.Getenv("MYSQL_URL"))
	sqlModel.AutoMigrate()
	redis.InitRedis(os.Getenv("REDIS_URL"))
	errors.SetCode2MsgMap(error_code.MsgMap)
	defer func() {
		logs.Sync()
	}()

	r := gin.Default()
	r.GET("/", wrapper.HandlerFuncWrapper(handler.Index))
	userGroup := r.Group("/user")
	// 注册
	userGroup.POST("/register", wrapper.HandlerFuncWrapper(user.Register))
	// 登录
	userGroup.POST("/login", wrapper.HandlerFuncWrapper(user.Login))
	// 需要登录权限
	userAuthGroup := userGroup.Group("/", auth.JwtAuthRequired())
	{
		// 退出登录
		userAuthGroup.POST("/logout", wrapper.HandlerFuncWrapper(user.Logout))
		// 查看当前登录用户个人基本信息
		userAuthGroup.GET("/profile", wrapper.HandlerFuncWrapper(user.Profile))
	}

	if err := r.Run(":8080"); err != nil {
		logs.PanicKvs("run server error", err)
	}
}
