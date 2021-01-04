package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

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

	rpc.Register(rect)

	rpc.HandleHTTP()

	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("listen err")
	}
}
