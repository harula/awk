package main

import (
	"fmt"
	"sync"
	"time"
)

var count, count2 int
var mutex sync.Mutex

func increase() {
	mutex.Lock()
	defer mutex.Unlock()
	count++
}

func increase2() {
	count2++
}

func main() {
	count = 0

	for i := 0; i < 1000; i++ {
		go increase()
	}
	for i := 0; i < 1000; i++ {
		go increase2()
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println(count, count2)
}
