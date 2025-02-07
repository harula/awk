package main

import "fmt"

const t = 2 //b树的度

type BTreeNode struct {
	keys     []int
	children []*BTreeNode
	leaf     bool
}

func NewBTreeNode(leaf bool) *BTreeNode {
	return &BTreeNode{
		keys:     make([]int, 0),
		children: make([]*BTreeNode, 0),
		leaf:     leaf,
	}
}

func main() {

	arr := []int{1, 2, 3, 4}
	fmt.Println(arr[:2])
	fmt.Println(arr[2:])

	arr2 := []int{}

	arr2 = append(arr2, arr[0])
	fmt.Println(arr)
}
