package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var wg sync.WaitGroup

func worker() {
	defer wg.Done()
	fmt.Println("i am a wg worker")
}

func worker2() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("i am not a wg worker")
}

func main() {
	wg.Add(1)
	go worker()
	wg.Wait()

	go worker2() //主协程退出，子协程可能未执行完成
}
