package main

import (
	"time"
	"fmt"
)

func main() {
	go sleep(2 * time.Second, "go1")
	go sleep(100 * time.Millisecond, "go2")
	sleep(1 * time.Second, "main")

	//main函数的退出会导致goroutine的终止
}

func sleep(t time.Duration, msg string) {
	fmt.Println(msg, " sleep start")
	time.Sleep(t)
	fmt.Println(msg, " sleep end")
}
