package main

import (
	"log"
	"sliceexample"
)

func main() {
	arr := [...]int{1, 2, 3}
	x := make([]int, 5)
	copy(x, arr[:])
	x = sliceexample.AppendInt(x, 4)
	log.Print(x)
	y := []int{5, 6, 7, 8}
	x = sliceexample.AppendSlice(x, y...)
	log.Print(x)
}
