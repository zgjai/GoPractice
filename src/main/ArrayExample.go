package main

import "fmt"

func main() {
	//数组是固定大小的
	//定义一个容量为5的数组
	var arr [5]int
	fmt.Println("array:", arr)
	fmt.Println("********我是分割线********")

	//通过下标访问数组元素
	arr[0] = 10
	arr[4] = 50
	fmt.Println(arr)
	fmt.Println("********我是分割线********")

	//创建数组并初始化
	b := []int{1, 2, 3}
	fmt.Println(b)
	//输出数组的长度
	fmt.Println(len(b))
	fmt.Println("********我是分割线********")

	//数组指针
	arrP := &arr
	(*arrP)[1] = 20
	fmt.Println(arr)
	fmt.Println(*arrP)
	fmt.Println("********我是分割线********")
}
