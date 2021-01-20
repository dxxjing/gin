package crontab

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func Crontest() {
	//cron.WithSeconds()  	 秒 | 分 | 时 | 每月的那一天(1-31) | 哪一月 | 每周哪一天
	//无cron.WithSeconds()   分 | 时 | 每月的那一天(1-31) | 哪一月 | 每周哪一天
	c := cron.New(cron.WithSeconds())

	//	"* * * * * *"  每秒执行
	//	"*/10 * * * * *" 每10秒
	//  "0 5 10 * * *"  每天早上 10:05:00执行
	//  "* * 10 * * *"  每天早上10点执行
	//  "* 1-59/10 * * * *"  每天1分-59分 每分钟执行一次
	//  "CRON_TZ=Asia/Tokyo 0 30 04 * * *"  指定时区下 每天上午 4:30:00执行一次


	//每分钟执行一次
	id, err := c.AddFunc("0 */1 * * * *", func() {
		fmt.Println("new crontab:", time.Now().Format("2006-01-02 15:04:05"))
	})
	if err != nil {
		fmt.Println(err, id)
		return
	}
	c.Start()

	//这是在http服务中运行 main.go主协程 不退出, 该协程start后内部执行死循环也不会退出，故此方法无需再后面加 死循环
	/*for{

	}*/
}
