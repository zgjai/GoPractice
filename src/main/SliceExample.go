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
	s2 := arr[:4]
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	//底层数组的修改会导致上层slice的改变
	arr[3] = 40
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	//slice的长度
	fmt.Println("s1 len:", len(s1))
	fmt.Println("s2 len:", len(s2))

	//slice的容量
	fmt.Println("s1 cap:", cap(s1))
	fmt.Println("s2 cap:", cap(s2))
	fmt.Println(s1[2])
	fmt.Println("********我是分割线********")



	//直接创建一个slice
	s3 := make([]int, 10)
	fmt.Println("s3:", s3)
	fmt.Println("s3 len:", len(s3))
	fmt.Println("s3 cap:", cap(s3))

	s4 := make([]int, 3, 10)        // s4 = make([]int, 10)[:3]
	fmt.Println("s4:", s4)
	fmt.Println("s4 len:", len(s4))
	fmt.Println("s4 cap:", cap(s4))
	fmt.Println("********我是分割线********")

	//new 的方式创建slice，基本不用
	s5 := new([]int)
	fmt.Println("s5:", s5)
	fmt.Println("*s5:", *s5)
	fmt.Println("*s5==nil?:", *s5 == nil)

	*s5 = append(*s5, 1)
	fmt.Println("s5:", s5)
	fmt.Println("s5:", *s5)
	fmt.Println("********我是分割线********")


	//slice的增长
	fmt.Println("s1 before append:", s1)
	fmt.Println("s1 len before append:", len(s1))
	fmt.Println("s1 cap before append:", cap(s1))

	s1 = append(s1, s3...)

	fmt.Println("s1 after append:", s1)
	fmt.Println("s1 len after append:", len(s1))
	fmt.Println("s1 cap after append:", cap(s1))


	fmt.Println("s2 after append:", s2)
	fmt.Println("s2 len after append:", len(s2))
	fmt.Println("s2 cap after append:", cap(s2))

	s2[3] = 4
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

}
