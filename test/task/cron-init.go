package task

import (
	"github.com/robfig/cron"
)

func PrepareCron() {
	spec := "0, 10, 1, *, *, *" // 每天6:40
	c := cron.New()
	c.AddFunc(spec, taskEveryDay())
	c.Start()
}
