package main

type sub struct {
	s string
	i int
}

type father struct {
	son *sub
	i   int
}

func main() {
	f := father{son: &sub{s: "a", i: 2}, i: 1}
	println(f.son.i)
}
