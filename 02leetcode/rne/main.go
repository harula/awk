package main

import "fmt"

// 1 2 3 4 5
//移除链表中倒数第n个节点
type ListNode struct {
	val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	dummy := &ListNode{0, head}
	first := head
	second := dummy

	for i := 0; i < n; i++ {
		if first == nil {
			return head
		}
		first = first.Next
	}
	//first = 2

	printList(first)
	printList(second)

	for first != nil {
		first = first.Next
		second = second.Next
	}

	// first = 5->next nil
	//second =
	//删除second后一个节点
	second.Next = second.Next.Next
	return dummy.Next
}

func printList(h *ListNode) {
	for h != nil {
		fmt.Printf("%d -> ", h.val)
		h = h.Next
	}
	fmt.Println("nil")
}

func main() {
	h := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	printList(h)

	h = removeNthFromEnd(h, 2)
	printList(h)
}
