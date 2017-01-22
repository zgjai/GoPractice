package main

import (
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//为了让main等待
	wg.Add(3)
	go unsafeGet("one")
	go unsafeGet("two")
	go unsafeGet("five")
	wg.Wait()

	wg.Add(3)
	go lockGet("one")
	go lockGet("two")
	go lockGet("five")
	wg.Wait()

	wg.Add(3)
	go rLockGet("one")
	go rLockGet("two")
	go rLockGet("five")
	wg.Wait()

	wg.Add(3)
	go onceGet("one")
	go onceGet("two")
	go onceGet("five")
	wg.Wait()
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
func unsafeGet(k string) int {
	defer wg.Done()

	if m == nil {
		initMap()
	}
	return m[k]
}


//通过加锁保证线程安全，但失去并发性，单线程访问
var mu sync.Mutex

func lockGet(k string) int {
	defer wg.Done()

	mu.Lock()
	defer mu.Unlock()
	if m == nil {
		initMap()
	}
	return m[k]
}

//通过读写锁来提高并发度，多个读不会阻塞
var rmu sync.RWMutex

func rLockGet(k string) int {
	defer wg.Done()

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

func onceGet(k string) int {
	defer wg.Done()

	one.Do(initMap)
	return m[k]
}