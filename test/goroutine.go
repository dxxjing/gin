package main

import "fmt"

func main() {

	var num = 1

	c1 := make(chan int, 2)
	c2 := make(chan int)

	go func(num int){
		c1 <- num+1
	}(num)

	go func(num int){
		c2 <- num+2

	}(num)


	fmt.Println(<-c1,<-c2)
}
