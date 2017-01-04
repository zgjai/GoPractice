package main

import "fmt"

func main() {
	//声明初始化一个变量
	var x int = 100
	var str string = "hello world"
	//声明初始化多个变量
	var i, j, k int = 1, 2, 3

	//不用指明类型，通过初始化值来推导
	var b = true //bool型

	fmt.Printf("%d\n%s\n%d\n%d\n%d\n%t\n", x, str, i, j, k, b)

	//x := 100 //等价于 var x int = 100;
	z := 111
	fmt.Println(z)

	//常量
	const y = "string"
	fmt.Println(y)
	//y = "newString"

	//指针
	var p *int
	p = &x
	fmt.Println(p)
	*p = 2
	fmt.Println(x)
}
