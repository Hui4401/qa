package storage

import (
	"github.com/Hui4401/qa/storage/mysql"
	sqlModel "github.com/Hui4401/qa/storage/mysql/model"
	"github.com/Hui4401/qa/storage/redis"
)

func InitStorage(mysqlURL string, redisURL string) {
	mysql.InitMySQL(mysqlURL)
	sqlModel.AutoMigrate()
	redis.InitRedis(redisURL)
}
