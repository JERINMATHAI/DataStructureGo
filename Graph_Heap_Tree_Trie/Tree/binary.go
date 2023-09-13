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

func main() {
	fmt.Println("Binary Tree")
	root := createNode(1)
	root.left = createNode(2)
	root.right = createNode(3)
	root.left.left = createNode(4)
	root.left.right = createNode(5)
	root.right.left = createNode(6)
	root.right.right = createNode(7)
	root.preOrder()
}
