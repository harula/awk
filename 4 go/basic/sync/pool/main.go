package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const (
	datalen = 1024
)

type Buffer struct {
	data []byte
	name string
	age  int
}

var numCalcsCreated int32

func createBuffer() interface{} {

	atomic.AddInt32(&numCalcsCreated, 1) //自定义pool的new函数，非线程安全
	buffer := &Buffer{
		data: make([]byte, datalen),
	}

	return buffer
}
func main() {
	bufferPool := &sync.Pool{
		New: createBuffer,
	}

	numWorkers := 1024 * 1024
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			v := bufferPool.Get()
			defer bufferPool.Put(v) //get与put必须成对出现，内存没有回收，池无限增大

			_, ok := v.(*Buffer)
			if !ok {
				fmt.Println("get buf failed")
			}

			return
		}()

	}
	wg.Wait()

	fmt.Printf("%d buffer created\n", numCalcsCreated)
}
