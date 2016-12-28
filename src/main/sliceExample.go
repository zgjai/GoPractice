package main

import (
	"fmt"
)

func main() {
	//切片(slice)是变长序列，底层是数组，多个切片可以共享同一个底层数组
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(arr)

	//前闭后开
	s1 := arr[3:6]
	s2 := arr[1:4]
	fmt.Println(s1)
	fmt.Println(s2)

	//底层数组的修改会导致上层slice的改变
	arr[3] = 40
	fmt.Println(s1)
	fmt.Println(s2)
}
