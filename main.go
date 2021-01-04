package main

import (
	"context"
	"gin-test/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	r := routers.InitRouter()

	//r.Run() //默认监听 0.0.0.0：8080
	//测试 http://localhost:8080/ping

	srv := &http.Server{
		Addr: ":8080",
		Handler: r,
		ReadTimeout: 30 * time.Second,
		WriteTimeout: 30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen %s \n", err)
		}
	}()
	//捕获信号 停止服务
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}