package main

import (
	"fmt"
	"reflect"
)

func main() {
	//函数是一类公民
	//函数类型跟int、string等类型一样，可以被赋值给变量
	max := maxF
	fmt.Println(max(1, 2, 3, 4, 5, 0, 2, 3, 10, -1))
	fmt.Println(max([]int{}...))
	fmt.Println("********I am split line********")

	//定义一个函数值min
	//参数类型为...int，参数名为nums
	//返回值 第一个类型为int，返回值名为min；第二个类型为error，返回值名为err
	min := func(nums ...int) (min int, err error) {
		if len(nums) == 0 {
			err = fmt.Errorf("nums is empty")
			return
		}
		min = nums[0]
		for _, i := range nums {
			if i < min {
				min = i
			}
		}
		return
	}
	//匿名函数
	fmt.Println(min([]int{-1, -2, -3, -4, -5, 1, -7, 0}...))
	fmt.Println("********I am split line********")

	nextNumFunc := nextNum()
	for i := 0; i < 10; i++ {
		fmt.Print(nextNumFunc(), "  ")
	}
}

//定义一个函数max
//参数类型为...int，参数名为nums
//返回值类型为int, error两个
func maxF(nums ...int) (int, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("nums is empty")
	}
	fmt.Println("nums: ", nums)
	fmt.Println("nums len: ", len(nums))
	fmt.Println("nums cap: ", cap(nums))
	fmt.Println("nums type: ", reflect.TypeOf(nums))
	var max int = nums[0]
	for _, i := range nums {
		if i > max {
			max = i
		}
	}
	return max, nil
}

//函数闭包
//返回值类型为函数
//返回的函数是一个无参返回值为int的函数
func nextNum() (func() int) {
	i, j := 0, 1
	return func() int {
		tmp := i + j
		i, j = j, tmp
		return i
	}
}
