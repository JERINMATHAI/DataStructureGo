package main

import "fmt"

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
				}
				temp = temp.right
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
func (b *bst) Contain(num int) {
	temp := b.root
	for temp != nil {
		if temp.data == num {
			fmt.Println("Element found")
			return
		} else if temp.data > num {
			temp = temp.left
		} else {
			temp = temp.right
		}
	}
	fmt.Println("Element not found")
}
func main() {
	fmt.Println("Contain")
	B := bst{}
	B.insert(8)
	B.insert(3)
	B.insert(1)
	B.insert(6)
	B.insert(4)
	B.insert(7)
	B.insert(10)
	B.insert(14)
	B.insert(13)

	fmt.Println("Pre Order Traversal")
	B.preOrder(B.root)
	fmt.Println("Contain")
	B.Contain(100)
}
