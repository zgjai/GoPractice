package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go transferData(ch)
	msg := <-ch
	fmt.Println(msg)
	close(ch)

	fmt.Println("********I am split line********")

	seriesTransfer()
}

//无缓冲channel的基本使用
func transferData(ch chan string) {
	fmt.Println("wait for transfer...")
	time.Sleep(1 * time.Second)
	ch <- "Hello"
}

//一个串联的channel
func seriesTransfer() {
	numCh := make(chan int)
	squaCh := make(chan int)
	go counter(numCh)
	go squarer(squaCh, numCh)
	printer(squaCh)
}

//chan<- 表示这个channel只能用于发送数据
func counter(out chan<- int) {
	for i := 1; i <= 100; i++ {
		time.Sleep(1 * time.Second)
		out <- i
	}
	close(out)
}

//<-chan 表示这个channel只能用于接收数据，这种类型的channel是不能主动关闭的
//range 可以在channel被关闭后，遍历完channel发出的所有数据后，退出循环
func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

func printer(in <-chan int) {
	for s := range in {
		fmt.Println("s: ", s)
	}
}
