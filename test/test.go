package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)

	go func(ch chan bool) {
		time.Sleep(5*time.Second)
		fmt.Println("gourtine done")
		ch <- true
	}(ch)
	 <- ch
	 fmt.Println("done")
}
