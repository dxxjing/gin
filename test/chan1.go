package main

import "time"

//var c = make(chan int) //无缓冲通道  读写都会阻塞
var c = make(chan int,1)
var a string

func f () {
	a = "hello word"
	time.Sleep(3*time.Second)
	<- c
}

func main() {
	go f()
	c <- 0
	println("111")
	println(a)
}

