package main

import (
	"fmt"
	"gin-test/Logger"
	"gin-test/routers"
	"github.com/facebookgo/grace/gracehttp"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"syscall"
	"time"
)



func main() {
	//初始化zap日志
	Logger.InitZapLogger()
	//将缓存中的日志刷入文件中 程序退出前应调用
	defer zap.L().Sync()
	fmt.Println("zap log start...")

	//定时任务
	//go crontab.Crontest()

	r := routers.InitRouter()

	/*Logger.Debug("jdx_test", Logger.CommFields{
		Uid: 3688,
	}, zap.Int("uu", 111), zap.String("kkk", "123"))*/
	//r.Run() //默认监听 0.0.0.0：8080
	//测试 http://localhost:8080/ping

	srv := &http.Server{
		Addr: ":8080",
		Handler: r,
		ReadTimeout: 600 * time.Second,
		WriteTimeout: 600 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("http srver start...")
	zap.L().Info("http server started")
	//平滑重启 必须要go install 然后再执行 GOPATH/bin/main
	//kill -SIGUSR2  PID 触发重启  SIGINT/SIGTERM 关闭服务
	WritePidToFile("master-test.pid", os.Getpid())
	gracehttp.Serve(srv)
	fmt.Println("master pid:", ReadPidFromFile("master-test.pid"))
	/*if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		zap.L().Error("start http srv err:")
	}*/
	fmt.Println("http server stop")
	/*//捕获信号 停止服务 不要有 监听SIGHUP信号 否则重启的SIGUSR2会导致 服务关闭
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	sig := <-quit
	fmt.Println("get signal shutdown server:", sig.String(), sig)
	zap.L().Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error("Server Shutdown err")
	}
	zap.L().Info("Server exiting")*/


}

func WritePidToFile (fileName string, pid int) {
	fileName = "./" + fileName
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	fmt.Println("write file fd:", file.Fd())
	if err != nil {
		Logger.ErrorMsg("open file err:" + err.Error())
		return
	}
	defer file.Close()
	if err = syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		fmt.Println("server is already running:" + err.Error())
		return
	}

	//pid写入文件
	pidStr := strconv.Itoa(pid)
	n, err := file.WriteString(pidStr)
	if n != len([]rune(pidStr)) || err != nil {
		fmt.Println("write err:" + err.Error())
		return
	}
	syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
}

func ReadPidFromFile (fileName string) int {
	fileName = "./" + fileName
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("server has stoped or read file err:" + err.Error())
		return -1
	}
	pid, _ := strconv.ParseInt(string(buf[:]), 10, 64)
	return int(pid)
}