package main

import "fmt"

type Node struct {
	key    int
	height int
	left   *Node
	right  *Node
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func height(n *Node) int {
	if n == nil {
		return 0
	}
	return n.height
}

func rightRotate(y *Node) *Node {
	x := y.left
	t2 := x.right

	x.right = y
	y.left = t2

	y.height = max(height(y.left), height(y.right)) + 1
	x.height = max(height(x.left), height(x.right)) + 1

	return x
}

func leftRotate(x *Node) *Node {
	y := x.right
	t2 := y.left

	y.left = x
	x.right = t2

	x.height = max(height(x.left), height(x.right)) + 1
	y.height = max(height(y.left), height(x.right)) + 1

	return y
}

func getBalance(n *Node) int {
	if n == nil {
		return 0
	}

	return height(n.left) - height(n.right)
}

func Insert(node *Node, key int) *Node {

	if node == nil {
		return &Node{key: key, height: 1}
	}

	if key < node.key {
		node.left = Insert(node.left, key)
	} else if key > node.key {
		node.right = Insert(node.right, key)
	} else {
		return node
	}

	node.height = max(height(node.left), height(node.right)) + 1

	balance := getBalance(node)

	//LL型，右旋
	if balance > 1 && key < node.left.key {
		return rightRotate(node)
	}

	//RR型左旋
	if balance < -1 && key > node.right.key {
		return leftRotate(node)
	}

	//LR型
	if balance > 1 && key > node.left.key {
		//左子树左旋变成LL型
		node.left = leftRotate(node.left)
		//L型右旋
		return rightRotate(node)
	}

	//RL型
	if balance < -1 && key < node.right.key {
		//右子树右旋，变成RR型
		node.right = rightRotate(node.right)
		//RR型左旋
		return leftRotate(node)
	}

	return node
}

func delete(node *Node,key)*Node{
	if root == nil {
		return root
	}

	if key < node.key {
		node.left = delete(node.left,key)
	}else if key >  node.key {
		node.right = delete(node.right,key)
	}else {

		if node.left == nil || node.right == nil {
			var tmp *Node
			if node.left != nil {
				tmp = node.left
			}else if node.right != nil {
				tmp = node.right
			}

			if tmp == nil {
				//左右子树均为空，直接删除节点（节点置空）
				//tmp = node
				node = nil
			}else {
				//更优雅的写法？
				*node = *tmp
			}
		}else{
			tmp := minValueNode(node.right)
			node.key = tmp.key
			node.right = delete(node.right,tmp.key)

		}
	}

	if root == nil {
		return nil
	}

	
}


func minValueNode(node *Node)*Node{
	curr := node
	for curr != nil {
		curr = curr.left
	}
	return curr
}
func preOrder(node *Node) {
	if node != nil {
		fmt.Printf("%d ", node.key)
		preOrder(node.left)
		preOrder(node.right)
	}
}

func main() {
	var root *Node

	keys := []int{10, 20, 30, 40, 50, 25}

	for _, key := range keys {
		root = Insert(root, key)
	}

	fmt.Println("preorder traversal:", root.key)
	preOrder(root)
}
