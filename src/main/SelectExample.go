package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)

	//并发获取数据
	go fetchData(ch)

	select {
	//等待数据返回
	case i := <-ch:
		fmt.Println("wait for ", i, " s")

	//设置3s的超时
	case <-time.After(3 * time.Second):
		fmt.Println("too long...")
	}
}

//模拟一个获取数据的函数，随机sleep 0-5s
func fetchData(ch chan<- int) {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(5)
	sleepST(i)
	ch <- i
	fmt.Println("done")
}

func sleepST(t int) {
	fmt.Println("sleep ", t, " s")
	time.Sleep(time.Duration(t) * time.Second)
	fmt.Println("wake up")
}
