package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type Queue struct {
	front *Node
	rear  *Node
}

func (q *Queue) enQueue(num int) {
	newNode := new(Node)
	newNode.data = num

	if q.front == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
}

func (q *Queue) Display() {
	temp := q.front

	for temp != nil {
		fmt.Printf("[%d]\n", temp.data)
		temp = temp.next
	}
}

func (q *Queue) deQueue() {
	if q.front == nil {
		fmt.Println("Queue is empty")
	} else {
		q.front = q.front.next
	}
}
func main() {
	fmt.Println("Queue implementation")
	queue := Queue{}
	var choice int
	for true {
		fmt.Println("Enter your choice")
		fmt.Println("1. enQueue")
		fmt.Println("2. Display")
		fmt.Println("3. deQueue")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			var num int
			fmt.Println("Enter the number")
			fmt.Scan(&num)
			queue.enQueue(num)
		case 2:
			queue.Display()
		case 3:
			queue.deQueue()
		}
	}
}
