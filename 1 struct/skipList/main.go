package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MaxLevel    int     = 16
	Probability float64 = 0.5
)

type Node struct {
	key   int
	value interface{}
	next  []*Node
}

type SkipList struct {
	head  *Node
	level int
}

func NewNode(key int, value interface{}, level int) *Node {
	return &Node{key: key, value: value, next: make([]*Node, level)}
}

func NewSkipList() *SkipList {
	return &SkipList{
		head:  NewNode(-1, nil, MaxLevel),
		level: 1,
	}
}

func (sl *SkipList) randomLevel() int {
	lv1 := 1
	for rand.Float64() < Probability && lv1 < MaxLevel {
		lv1++
	}
	return lv1
}

// 跳跃表搜索
func (sl *SkipList) Search(key int) *Node {
	curr := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for curr.next != nil && curr.next[i].key < key {
			curr = curr.next[i]
		}
	}

	curr = curr.next[0]
	if curr != nil && curr.key == key {
		return curr
	}

	return nil
}

// 输入8
// lv2 : 1   5   9
// 1v1 : 1 3 5 7 9
func (sl *SkipList) Insert(key int, value interface{}) {
	update := make([]*Node, MaxLevel)
	curr := sl.head

	//记录待插入节点前向节点
	for i := sl.level - 1; i >= 0; i-- {
		for curr.next[i] != nil && curr.next[i].key < key {
			curr = curr.next[i]
		}

		update[i] = curr
	}

	curr = curr.next[0]

	//未命中key
	if curr == nil || curr.key != key {
		randomLevel := sl.randomLevel()

		//randomLevel大于当前最大level，更新level
		if randomLevel > sl.level {
			for i := sl.level; i < randomLevel; i++ {
				update[i] = sl.head
			}

			sl.level = randomLevel
		}

		//构造新节点
		newNode := NewNode(key, value, randomLevel)

		//插入新节点
		for i := 0; i < randomLevel; i++ {
			newNode.next[i] = update[i].next[i]
			update[i].next[i] = newNode
		}
	} else {
		//命中key，则更新节点
		curr.value = value
	}
}

func (sl *SkipList) Delete(key int) {
	update := make([]*Node, MaxLevel)
	curr := sl.head

	//记录待删除节点前向节点
	for i := sl.level - 1; i >= 0; i-- {
		for curr.next[i] != nil && curr.next[i].key < key {
			curr = curr.next[i]
		}
		update[i] = curr
	}

	curr = curr.next[0]
	//命中key则删除节点
	if curr != nil && curr.key == key {
		for i := 0; i < sl.level; i++ {
			if update[i].next[i] != curr {
				//待删除节点层数小于sl的层数，高于节点层不处理
				break
			}
			//删除当前节点
			update[i].next[i] = curr.next[i]
		}

		//从顶向下遍历，如果该层head指向空则level向下偏移
		for sl.level > 1 && sl.head.next[sl.level-1] == nil {
			sl.level--
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	sl := NewSkipList()

	sl.Insert(1, "one")

	fmt.Println("sl:", sl)
	sl.Insert(2, "two")
	fmt.Println("sl:", sl)
	sl.Insert(3, "tree")
	fmt.Println("sl:", sl)

	node := sl.Search(2)
	if node != nil {
		fmt.Printf("found key %d with value:%v\n", node.key, node.value)
	} else {
		fmt.Printf("key not found")
	}

	sl.Delete(2)
	fmt.Printf("sl:%v", sl)
	node = sl.Search(2)
	if node != nil {
		fmt.Printf("found key %d with value:%v\n", node.key, node.value)
	} else {
		fmt.Printf("key not found")
	}
}
