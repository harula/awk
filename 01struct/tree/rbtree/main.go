package main

import "fmt"

const (
	RED   = true
	BLACK = false
)

type Node struct {
	color   bool
	value   int
	left    *Node
	right   *Node
	parrent *Node
}

type RBTree struct {
	root *Node
}

func NewRBTree() *RBTree {
	return &RBTree{}
}

func (t *RBTree) RotateLeft(x *Node) {
	y := x.right

	x.right = y.left

	if y.left != nil {
		y.left.parrent = x
	}

	y.parrent = x.parrent
	if x.parrent == nil {
		t.root = y
	} else if x == x.parrent.left {
		x.parrent.left = y
	} else {
		x.parrent.right = y
	}

	y.left = x
	x.parrent = y
}

func (t *RBTree) RotateRight(y *Node) {
	x := y.left

	y.left = x.right
	if x.right != nil {
		x.right.parrent = y
	}

	x.parrent = y.parrent
	if y.parrent == nil {
		t.root = x
	} else if y == y.parrent.left {
		y.parrent.left = x
	} else {
		y.parrent.right = x
	}

	x.right = y
	y.parrent = x
}

func (t *RBTree) fixInsert(node *Node) {
	//新增节点为红色，新增节点的父节点如果也为红色，则需要递推维护红黑树的颜色
	for node != t.root && node.parrent.color == RED {
		//1. 父节点为左子树
		if node.parrent == node.parrent.parrent.left {
			uncle := node.parrent.parrent.right
			//1.1 叔节点为红色
			if uncle.color == RED {
				//1.将父节点和叔节点变为黑色，将爷节点变为红色，将当前节点变为爷节点进行递推
				node.parrent.color = BLACK
				uncle.color = BLACK
				node.parrent.parrent.color = RED
				node = node.parrent.parrent

			} else {
				//1.2 叔节点为黑色
				//1.1.1 新增节点为左子树 - 将父节点变为黑色，爷节点变为红色，以爷节点进行右旋
				//1.1.2 新增节点为右子树 - 以父节点为中心进行左旋，变为1.1 结构，执行1.1.1相同变换

				if node == node.parrent.right {
					//左旋后父节点变为左子树，新增节点变为父节点，形态与1.1.1一致
					node = node.parrent
					t.RotateLeft(node)
				}

				node.parrent.color = BLACK
				node.parrent.parrent.color = RED
				t.RotateRight(node.parrent.parrent)
			}
		} else if node.parrent == node.parrent.parrent.right {
			//2.新增节点为右子树
			uncle := node.parrent.parrent.left
			if uncle.color == RED {
				//2.1 叔节点为红色
				//将父节点和叔节点变为黑色，爷节点变为红色，当前节点变为爷节点，进行递推
				uncle.color = BLACK
				node.parrent.color = BLACK
				node.parrent.parrent.color = RED
				node = node.parrent.parrent

			} else {
				//2.2 叔节点为黑色
				//2.2.1 新增节点为左子树 - 以父节点为中心进行右旋，结构变为2.2.2，执行2.2.2相同变换
				//2.2.2 新增节点为右子树 - 将父节点变为黑色，爷节点变为红色，以爷节点为中心进行左旋
				if node == node.parrent.left {
					node = node.parrent
					t.RotateRight(node)
				}
				node.parrent.color = BLACK
				node.parrent.parrent.color = RED
				t.RotateLeft(node.parrent.parrent)

			}
		}
	}

	//递推结束后根节点可能为红色，则将其变为黑色
	t.root.color = BLACK
}

func (t *RBTree) Insert(value int) {
	newNode := &Node{value: value, color: RED}
	if t.root == nil {
		t.root = newNode
	} else {
		curr := t.root
		for curr != nil {
			if value < curr.value {
				if curr.left == nil {
					curr.left = newNode
					break
				}
				curr = curr.left
			} else if value > curr.value {
				if curr.right == nil {
					curr.right = newNode
					break
				}
				curr = curr.right
			}
		}

		newNode.parrent = curr
	}

	t.fixInsert(newNode)

}

func printInOrder(node *Node) {
	if node != nil {
		printInOrder(node.left)
		fmt.Println(node.value)
		printInOrder(node.right)
	}

}
func main() {
	t := NewRBTree()
	values := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	for _, v := range values {
		t.Insert(v)
	}

	printInOrder(t.root)
}
