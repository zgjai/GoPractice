package main

import (
	"math"
	"fmt"
)

//定义一个接口
type shape interface {
	//计算面积
	area() float64
	//计算周长
	perimeter() float64
}

//定义长方形
type rect struct {
	width, height float64
}

//求面积
func (r *rect) area() float64 {
	return r.width * r.height
}

//求周长
func (r *rect) perimeter() float64  {
	return 2*(r.width+r.height)
}

//定义圆形
type circle struct {
	radius float64
}

//求面积
func (c *circle) area() float64 {
	return math.Pi * math.Sqrt(c.radius)
}

//求周长
func (c *circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func printArea(s shape)  {
	fmt.Println("s area: ", s.area())
}

func main() {
	r := rect{width:2.0, height:3.5}
	fmt.Println("r area: ", r.area())
	fmt.Println("r perimeter: ", r.perimeter())

	c := circle{3}
	fmt.Println("c area: ", c.area())
	fmt.Println("c perimeter: ", c.perimeter())

	printArea(&r)
	printArea(&c)
}