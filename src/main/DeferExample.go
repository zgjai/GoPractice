package main

import (
	"os"
	"log"
	"io/ioutil"
	"fmt"
	"time"
)

func main() {
	openFile("/Users/zhangguijiang/tmpDoc/newFile")
	//openFile("/Users/zhangguijiang/tmpDoc/noFile")
	fmt.Println("********I am split line********")

	slowOperation()
}

func openFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("open file err", err)
	}
	//在退出函数之前定会执行
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("read file err", err)
	}

	fmt.Printf("%s", b)
}


func slowOperation() {
	//最后的()不能丢！！！
	defer trace("slowOperation")()

	//sleep 3s
	time.Sleep(3*time.Second)
}
//一个通用的追踪函数
//可以用于记录函数的进入退出以及执行时间
func trace(msg string) func() {
	start := time.Now()
	log.Println("enter ", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

