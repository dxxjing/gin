package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)
//参数结构体 因为 rpc.esClient.Call(serviceMethod string, args interface{}, reply interface{})
//只接收三个参数 依次为 远程方法名  参数列表  返回值
type Params struct {
	Width, Height int
}

type Rect struct {

}

func (r *Rect) Area(p Params, ret *int) (err error) {
	*ret = p.Width * p.Height
	return nil
}

func main() {
	rect := new(Rect)
	//注册服务
	rpc.Register(rect)
	//绑定到http
	rpc.HandleHTTP()
	//监听
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("listen err")
	}
}
