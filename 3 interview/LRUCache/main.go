package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type entry struct {
	key   int
	value int
}

func Constructor(cap int) LRUCache {
	return LRUCache{
		capacity: cap,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	if ele, ok := c.cache[key]; ok {
		c.list.MoveToFront(ele)
		return ele.Value.(*entry).value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	//元素存在列表，将元素移到队首，重置map中元素得值
	if ele, ok := c.cache[key]; ok {
		c.list.MoveToFront(ele)
		ele.Value.(*entry).value = value
		return
	}

	//超出容量删除队尾元素
	if c.list.Len() == c.capacity {
		oldest := c.list.Back()
		c.list.Remove(oldest)
		delete(c.cache, oldest.Value.(*entry).key)
	}

	//插入新元素
	newEntry := &entry{key: key, value: value}
	ele := c.list.PushFront(newEntry)
	c.cache[key] = ele
}

func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1))

	lru.Put(3, 3)
	fmt.Println(lru.Get(2))
	lru.Put(4, 4)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(3))
	fmt.Println(lru.Get(4))
}
