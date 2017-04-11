package main

import (
	"unicode/utf8"
	"fmt"
	"strings"
	"unsafe"
	"reflect"
)

func main() {
	s := "中bcdef"
	fmt.Println(len(s))
	r, size := utf8.DecodeRuneInString(s)
	fmt.Println(r, size)
	i := strings.IndexRune(s, rune('中'))
	fmt.Println(i)

	s1 := make([]int, 10)
	for i := range s1 {
		s1[i] = i
	}
	fmt.Println(unsafe.Pointer(&s1))
	fmt.Printf("%#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&s1)))
	s1 = s1[:]
	fmt.Println(unsafe.Pointer(&s1))
	fmt.Printf("%#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&s1)))
	s1 = s1[1:]
	fmt.Println(unsafe.Pointer(&s1))
	fmt.Printf("%#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&s1)))
	s2 := s1[:]
	fmt.Println(unsafe.Pointer(&s2))
	fmt.Printf("%#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&s2)))
}
