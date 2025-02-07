package main

import (
	"sync"
	"time"
)

var count int
var rwMutex sync.RWMutex

func read() {
	rwMutex.RLock()
	defer rwMutex.RUnlock()

	print(count)
}

func Write() {
	rwMutex.Lock()
	defer rwMutex.Unlock()

	count++
}

func main() {
	for i := 0; i < 10; i++ {
		go read()
		go Write()
	}

	time.Sleep(10 * time.Millisecond)
	read()
}
