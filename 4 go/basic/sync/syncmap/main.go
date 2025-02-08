package main

import (
	"fmt"
	"sync"
)

// sync.Map适用读多写少场景，或者键值对数据量较大的场景
func main() {
	var m sync.Map
	m.Store("one", 1)
	m.Store("two", 2)
	v, ok := m.Load("one")
	fmt.Println(v, ok)

	v, ok = m.LoadOrStore("one", 3) //读到的是更新前的数据
	fmt.Println(v, ok)

	m.Range(func(k, v interface{}) bool {
		fmt.Println("range map", k, v)
		return true
	})
}
