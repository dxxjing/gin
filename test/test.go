package main

import "fmt"

func main() {
	//切片传指针
	s := []int{1,2,3}
	fmt.Println(len(s), cap(s), s)
	Add(&s)
	fmt.Println(len(s),cap(s), s)
}

func Add(s *[]int) {
	*s = append(*s, 4)
}
