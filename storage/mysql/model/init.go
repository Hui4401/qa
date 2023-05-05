package model

import (
	"github.com/Hui4401/gopkg/logs"

	"github.com/Hui4401/qa/storage/mysql"
)

// AutoMigrate 注册所有model的自动迁移
func AutoMigrate() {
	client := mysql.GetClient()
	if err := client.AutoMigrate(&User{}, &UserProfile{}); err != nil {
		logs.PanicKvs("auto migration error", err)
	}
}
