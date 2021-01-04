package main

import "fmt"

func main () {
	//切片并不存储任何数据，它只是描述了底层数组中的一段。
	//
	//更改切片的元素会修改其底层数组中对应的元素。
	//
	//与它共享底层数组的切片都会观测到这些修改。
	var arr = [10]int{0,1,2,3,4,5,6,7,8,9}
	s1 := arr[1:6]
	s2 := arr[2:7]
	fmt.Println(s1, s2, arr, cap(s1), cap(s2)) //[1 2 3 4 5] [2 3 4 5 6] [0 1 2 3 4 5 6 7 8 9]

	s1[3] = 99
	fmt.Println(s1, s2, arr) //[1 2 3 99 5] [2 3 99 5 6] [0 1 2 3 99 5 6 7 8 9]

	ss := make([]int, 4)
	ss[0] = 1
	fmt.Println(ss)

}
