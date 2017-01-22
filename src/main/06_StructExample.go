package main

import (
	"fmt"
)

type person struct {
	age  int
	name string
	s    sex
}

type sex int

const (
	male sex = iota
	female
	unknown
)

func main() {
	fmt.Println(sex(male))
	fmt.Println(sex(female))
	fmt.Println(sex(unknown))

	//初始化一个struct的两种方式
	//需要按定义的顺序赋值，常用于字段较少的情况
	p1 := person{20, "张三", sex(male)}
	//顺序无关，读者也清晰
	p2 := person{age:18, s:female, name:"莉丝"}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println("********I am split line********")

	//访问struct中的元素
	fmt.Println(p1.age)
	fmt.Println(p2.name)
	p1.name = "张一"
	fmt.Println(p1)
	fmt.Println("********I am split line********")

	//struct的指针
	p3 := &person{age:18, s:female, name:"莉丝"}
	fmt.Println(*p3)
	fmt.Println((*p3).s)
	fmt.Println(p3.s)
	p3.age = 19
	fmt.Println(*p3)
	(*p3).age = 20
	fmt.Println(*p3)

}
