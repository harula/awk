package main

type List struct {
	head *Node
	tail *Node
}
type Node struct {
	data int
	next *Node
}

// 如何实现单向链表从1/2处反转 只遍历一次的情况下
// 思路：使用快遍历（一次遍历2个节点）和慢遍历（一次遍历一个节点）两种方式
// 1 ->2-> 3 ->4 ->5
// 3->4->5->1->2
func revert(l List) List {
	len := 0

	firstTail := l.tail
	for n := l.head; n != nil; {

		l.tail.next = n
		l.tail = n
		l.tail.next = nil
		n = n.next
		l.head = n
	}

	return l
}
