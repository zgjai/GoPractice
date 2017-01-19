package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {

}

//单例模式lazySet多个版本的实现
var m map[string]int

//initialize
func initMap() {
	fmt.Println("initialization starting")

	//为了稳定复现并发问题
	time.Sleep(1 * time.Second)

	m = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}

	fmt.Println("initialization done")
}

//concurrency unsafe
func unsafeGet(k string) (int, bool) {
	if m == nil {
		initMap()
	}
	return m[k]
}


//通过加锁保证线程安全，但失去并发性，单线程访问
var mu sync.Mutex

func lockGet(k string) (int, bool) {
	mu.Lock()
	defer mu.Unlock()
	if m == nil {
		initMap()
	}
	return m[k]
}

//通过读写锁来提高并发度，多个读不会阻塞
var rmu sync.RWMutex

func rLockGet(k string) (int, bool) {
	rmu.RLock()
	if m != nil {
		rmu.RUnlock()
		return m[k]
	}
	rmu.RUnlock()

	rmu.Lock()
	defer rmu.Unlock()
	//必须再次检查
	if m == nil {
		initMap()
	}
	return m[k]
}

//clean code
var one sync.Once

func onceGet(k string) (int, bool) {
	one.Do(initMap)
	return m[k]
}