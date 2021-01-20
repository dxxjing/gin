package main

import (
	"fmt"
	"gin-test/Logger"
	"gin-test/crontab"
	"gin-test/routers"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func main() {
	//初始化zap日志
	Logger.InitZapLogger()
	//将缓存中的日志刷入文件中 程序退出前应调用
	defer zap.L().Sync()

	//定时任务
	go crontab.Crontest()



	r := routers.InitRouter()

	Logger.Debug("jdx_test", Logger.CommFields{
		Uid: 3688,
	}, zap.Int("uu", 111), zap.String("kkk", "123"))
	//r.Run() //默认监听 0.0.0.0：8080
	//测试 http://localhost:8080/ping

	srv := &http.Server{
		Addr: ":8080",
		Handler: r,
		ReadTimeout: 30 * time.Second,
		WriteTimeout: 30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("http srver start...")
	zap.L().Info("http server started")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		zap.L().Error("start http srv err:")
	}

	/*//捕获信号 停止服务
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("get signal shutdown server")
	zap.L().Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error("Server Shutdown err")
	}
	zap.L().Info("Server exiting")*/


}