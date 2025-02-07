package main

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) Insert(value int) {
	if value <= n.Value {
		if n.Left == nil {
			n.Left = &TreeNode{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

func (n *TreeNode) Find(val int) bool {
	if n == nil {
		return false
	}

	if val < n.Value {
		return n.Left.Find(val)
	} else if val > n.Value {
		return n.Right.Find(val)
	}

	return true
}

func (n *TreeNode) InOrderTraverse() {
	fmt.Println("enter InOrderTraverse", n)
	if n != nil {
		n.Left.InOrderTraverse()
		fmt.Println(n.Value)
		n.Right.InOrderTraverse()
	}
}

func (n *TreeNode) PreOrderTraverse() {
	if n != nil {
		fmt.Println(n.Value)
		n.Left.PreOrderTraverse()
		n.Right.PreOrderTraverse()
	}
}

func (n *TreeNode) PostOrderTraverse() {
	fmt.Println("enter PostOrderTraverse", n)
	if n != nil {
		n.Left.PostOrderTraverse()
		n.Right.PostOrderTraverse()
		fmt.Println(n.Value)
	}
}

func main() {
	root := &TreeNode{Value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(12)

	root.InOrderTraverse()
}
