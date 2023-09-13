package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func createNode(num int) *Node {
	newNode := new(Node)
	newNode.data = num
	newNode.left = nil
	newNode.right = nil
	return newNode
}

func (node *Node) preOrder() {
	if node == nil {
		return
	}
	fmt.Printf("%d\n", node.data)
	node.left.preOrder()
	node.right.preOrder()
}

func validateBST(root *Node, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.data <= min || root.data >= max {
		return false
	}
	return validateBST(root.left, min, root.data) && validateBST(root.right, max, root.data)
}

func main() {
	fmt.Println("Validate BST")
	root := createNode(8)
	root.left = createNode(3)
	root.right = createNode(10)
	root.left.right = createNode(6)
	root.left.left = createNode(1)

	root.preOrder()
	if validateBST(root, -1000, 1000) {
		fmt.Println("Valid BST")
	} else {
		fmt.Println("Invalid BST")
	}
}
