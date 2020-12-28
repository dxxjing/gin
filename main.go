package main

import (
	"gin-test/routers"
)

func main() {
	r := routers.InitRouter()

	r.Run() //默认监听 0.0.0.0：8080
	/*go func() {

	}()*/

	//测试 http://localhost:8080/ping
}