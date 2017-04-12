package test

import (
	"fmt"
	"testing"
)

func printCommon(i int) {
	fmt.Print(i)
}

func printClosure(i int) (func()) {
	return func() {
		fmt.Print(i)
	}
}

func BenchmarkCommon(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printCommon(i)
	}
}

func BenchmarkClosure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printClosure(i)()
	}
}
