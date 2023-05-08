package conf

import (
	"log"
	"os"

	"github.com/Hui4401/gopkg/errors"
	"github.com/Hui4401/gopkg/logs"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Hui4401/qa/constdef"
	"github.com/Hui4401/qa/storage"
)

func Init() {
	// 从文件读取环境变量，默认读取 .env 文件
	if err := godotenv.Load(); err != nil {
		log.Panicf("Error loading .env file: %v", err)
	}

	ginMode := os.Getenv("GIN_MODE")
	gin.SetMode(ginMode)

	logPath := os.Getenv("LOG_PATH")
	logMode := logs.LogModeDevelopment
	if ginMode == gin.ReleaseMode {
		logMode = logs.LogModeProduct
	}
	logs.Init(logMode, logPath)

	errors.SetUnknownCode(constdef.CodeUnknown)
	errors.SetUnknownMsg(constdef.MsgMap[constdef.CodeUnknown])
	errors.SetCode2MsgMap(constdef.MsgMap)

	storage.InitStorage(os.Getenv("MYSQL_URL"), os.Getenv("REDIS_URL"))
}
