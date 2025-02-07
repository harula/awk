package main

import "math/rand"

const (
	Probability = 0.5
	MaxLevel    = 16
)

type SlNode struct {
	key   int
	value interface{}
	next  []*SlNode
}

type SkipList struct {
	head  *SlNode
	level int
}

func NewSlNode(key int, value interface{}, level int) *SlNode {
	return &SlNode{key: key, value: value, next: make([]*SlNode, level)}
}

func NewSkipList() *SkipList {
	return &SkipList{head: NewSlNode(-1, nil, MaxLevel), level: 1}
}

func randomLevel() int {
	lv := 1
	for rand.Float64() < Probability && lv < MaxLevel {
		lv++
	}
	return lv
}
func (sl *SkipList) Search(key int) *SlNode {
	curr := sl.head
	for i := sl.level - 1; i >= 0; i++ {
		for curr != nil && curr.next[i].key < key {
			curr = curr.next[i]
		}
	}
	curr = curr.next[0]
	if curr != nil && curr.key == key {
		return curr
	}
	return nil
}

func (sl *SkipList) Insert(key int, v interface{}) {
	forward := make([]*SlNode, MaxLevel)
	curr := sl.head

	//第一步：自顶向下，从前向后遍历跳表
	for i := sl.level - 1; i >= 0; i-- {
		for curr.next[i] != nil && curr.next[i].key < key {
			curr = curr.next[i]
		}
		//记录每一层的待插入节点的前向节点
		forward[i] = curr
	}

	//第二步：判断节点是否存在
	curr = curr.next[0]
	if curr != nil && curr.key == key {
		//命中，则更新value
		curr.value = v
	} else {
		//未命中，则插入节点
		randomLevel := randomLevel()
		newNode := NewSlNode(key, v, randomLevel)

		if randomLevel > sl.level {
			for i := sl.level; i < randomLevel; i++ {
				//补充前向节点
				forward[i] = sl.head
			}
			sl.level = randomLevel
		}

		//逐层插入节点
		for i := 0; i < randomLevel; i++ {
			newNode.next[i] = forward[i].next[i]
			forward[i].next[i] = newNode
		}

	}
}

func (sl *SkipList) Delete(key int) {
	forword := make([]*SlNode, MaxLevel)
	curr := sl.head

	for i := sl.level - 1; i >= 0; i-- {
		for curr.next[i] != nil && curr.next[i].key < key {
			curr = curr.next[i]
		}
		forword[i] = curr
	}

	curr = curr.next[0]

	if curr != nil && curr.key == key {
		for i := 0; i < sl.level; i++ {

			if forword[i].next[i] != curr {
				break
			}
			forword[i].next[i] = curr.next[i]
		}

		for sl.level > 1 && sl.head.next[sl.level-1] == nil {
			sl.level--
		}
	}
}
