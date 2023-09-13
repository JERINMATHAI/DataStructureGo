package main

import (
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

type bst struct {
	root *Node
}

func (b *bst) insert(num int) {
	newNode := new(Node)
	newNode.data = num

	if b.root == nil {
		b.root = newNode
	} else {
		temp := b.root
		for temp != nil {
			if newNode.data < temp.data {
				if temp.left == nil {
					temp.left = newNode
					return
				} else {
					temp = temp.left
				}
			} else {
				if temp.right == nil {
					temp.right = newNode
					return
				} else {
					temp = temp.right
				}
			}
		}
	}
}

func (b *bst) preOrder(node *Node) {

	if node == nil {
		return
	}

	fmt.Printf("%d\n", node.data)
	b.preOrder(node.left)
	b.preOrder(node.right)
}

func (b *bst) postOrder(node *Node) {

	if node == nil {
		return
	}

	b.postOrder(node.left)
	b.postOrder(node.right)
	fmt.Printf("%d\n", node.data)
}

func (b *bst) inOrder(node *Node) {
	if node == nil {
		return
	}
	b.inOrder(node.left)
	fmt.Printf("%d\n", node.data)
	b.inOrder(node.right)
}
func main() {

	fmt.Println("Binary Search Tree")
	fmt.Println("'''''''''''''''''''")

	binary := bst{}
	binary.insert(8)
	binary.insert(3)
	binary.insert(1)
	binary.insert(6)
	binary.insert(4)
	binary.insert(7)
	binary.insert(10)
	binary.insert(14)
	binary.insert(13)

	fmt.Println("Pre Order Traversal")
	binary.preOrder(binary.root)
	fmt.Println("Post Order Traversal")
	binary.postOrder(binary.root)
	fmt.Println("Inorder Traversal")
	binary.inOrder(binary.root)
}
