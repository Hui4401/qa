package cron

import (
	"github.com/Hui4401/gopkg/logs"
	"github.com/robfig/cron"
)

func StartSchedule() {
	c := cron.New()

	// 每30分钟将redis数据同步到mysql
	addCronFunc(c, "@every 30m", func() {})

	// 每30分钟同步热榜信息
	addCronFunc(c, "@every 30m", func() {})

	c.Start()
}

func addCronFunc(c *cron.Cron, sep string, cmd func()) {
	err := c.AddFunc(sep, cmd)
	if err != nil {
		logs.PanicKvs("cron job error", err)
	}
}
