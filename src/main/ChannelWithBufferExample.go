package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个buffer size为3的channel
	ch := make(chan int, 3)
	go func(c chan<- int) {
		for i := 0; i < 10; i++ {
			c <- i
			fmt.Println("send: ", i)
		}
		close(c)
	}(ch)

	time.Sleep(1 * time.Second)
	for msg:= range ch{
		fmt.Println("recive: ", msg)
		time.Sleep(1 * time.Second)
	}
}