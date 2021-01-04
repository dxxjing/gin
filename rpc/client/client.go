package main

import (
	"fmt"
	"net/rpc"
)

type ParamsC struct {
	Width, Height int
}

func main () {
	var (
		cli *rpc.Client
		err error
		ret int
	)
	//连接远程rpc服务
	if cli, err = rpc.DialHTTP("tcp", ":8081"); err != nil {
		fmt.Println("dial http err")
		return
	}
	//调用方法
	if err = cli.Call("Rect.Area", ParamsC{11, 5}, &ret); err != nil {
		fmt.Println("rpc call err")
		return
	}
	fmt.Println("area:", ret)
}
