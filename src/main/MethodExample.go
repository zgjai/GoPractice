package main

import (
	"fmt"
)

type coder struct {
	name      string
	languages []string
}

func sayName(c coder) {
	fmt.Println("my name is: ", c.name)
}

//coder类型的method，接收者类型为coder
func (c coder) sayName() {
	fmt.Println("my name is: ", c.name)
}

//coder类型的method，接收者类型为coder
func (c coder) learnLanguageUnknown(language string) {
	fmt.Println("learn more language: ", language)
	c.languages = append(c.languages, language)
}

//coder类型的method，接收者类型为*coder
func (c *coder) learnLanguage(language string) {
	fmt.Println("learn more language: ", language)
	c.languages = append(c.languages, language)
}

func main() {
	c := coder{"Tom", []string{"java"}}
	fmt.Println(c)
	fmt.Println("********I am split line********")

	sayName(c)
	c.sayName()
	fmt.Println("********I am split line********")

	c.learnLanguageUnknown("Go")
	fmt.Println(c)
	fmt.Println("********I am split line********")

	c.learnLanguage("Go")
	fmt.Println(c)
	fmt.Println("********I am split line********")

	i := newInt(1)
	j := newInt(2)
	fmt.Println(i.compareNewInt(j))
}

//为基本类型添加自定义method
type newInt int

func (i newInt) compareNewInt(j newInt) bool {
	if i > j {
		return true
	} else {
		return false
	}
}
