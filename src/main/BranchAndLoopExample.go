package main

import (
	"fmt"
)

func main() {
	ifExample(8)
	ifExample(7)
	ifExample(6)
	fmt.Println("********我是分割线********")

	switchExample(1)
	switchExample(6)
	switchExample(7)
	fmt.Println("********我是分割线********")

	forExample(10)
	fmt.Println("********我是分割线********")

	forRangeExample()
}

//if条件语句示例
func ifExample(num int) {
	if num%4 == 0 {
		fmt.Println("4")
	} else if num%2 == 0 {
		fmt.Println("2")
	} else {
		fmt.Println("奇数")
	}
}

//switch语句示例
func switchExample(num int) {
	switch num {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4, 5, 6:
		//case 可以包含多个值，逗号隔开
		fmt.Println("4 or 5 or 6")
	default:
		fmt.Println("invalid value")
	}
}

//for循环语句示例，go中没有while
func forExample(num int) {
	//这3种for实现的效果相同
	for i := 0; i < num; i++ {
		fmt.Print(i, "  ")
	}

	i := 0
	for i < num {
		fmt.Print(i, "  ")
		i++
	}

	i = 0
	for {
		if i < num {
			break
		}
		fmt.Print(i, "  ")
		i++
	}
}

//for中使用range内置函数示例
func forRangeExample() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, v := range arr {
		fmt.Println("index: ", i, "value: ", v)
	}

	s := arr[0:10]

	for i, v := range s {
		fmt.Println("index: ", i, "value: ", v)
	}

	//初始化方式，等同于 m := map[string]string{}
	m := make(map[string]string)
	m["1"] = "1"
	m["2"] = "2"

	for k, v := range m{
		fmt.Println("key: ", k, "value: ", v)
	}
}
