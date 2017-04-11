package main

import (
	"strings"
	"fmt"
)

func main() {
	a := "abcdeftjklikklmn"
	b := "gha"
	i := strings.IndexAny(a, b)
	fmt.Println(i)
	a = "中国"
	for _, c := range a {
		fmt.Println(string(c))
	}
	for i = 0; i < len(a); i++ {
		fmt.Println(string(a[i]))
	}
}
