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
	if cli, err = rpc.DialHTTP("tcp", ":8081"); err != nil {
		fmt.Println("dial http err")
		return
	}

	if err = cli.Call("Rect.Area", ParamsC{12, 5}, &ret); err != nil {
		fmt.Println("rpc call err")
		return
	}
	fmt.Println("area:", ret)
}
