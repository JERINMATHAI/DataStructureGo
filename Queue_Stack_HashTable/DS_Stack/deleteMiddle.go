package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) push(num int) {
	newNode := new(Node)
	newNode.data = num

	if s.top == nil {
		s.top = newNode
	} else {
		newNode.next = s.top
		s.top = newNode
	}
}

func (s *Stack) pop() {
	if s.top == nil {
		fmt.Println("Stack is empty")
	} else {
		s.top = s.top.next
	}
}

func (s *Stack) DeleteMiddile() {

}
func (s *Stack) Display() {
	temp := s.top
	for temp != nil {
		fmt.Printf("[%d]\n", temp.data)
		temp = temp.next
	}
}
func main() {
	fmt.Println("Stack Implementation")
	stack := Stack{}
	var choice int
	for true {
		fmt.Println("Enter your choice")
		fmt.Println("1. Push")
		fmt.Println("2. Pop")
		fmt.Println("3. Display")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			var num int
			fmt.Println("Enter a number")
			fmt.Scan(&num)
			stack.push(num)
		case 2:
			stack.pop()
		case 3:
			stack.Display()
		}
	}
}
