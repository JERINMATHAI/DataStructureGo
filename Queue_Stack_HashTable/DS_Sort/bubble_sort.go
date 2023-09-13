package main

import "fmt"

func bubbleSort(array []int) {
	n := len(array)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

func main() {
	fmt.Println("Bubble Sort")
	array := []int{2, 8, 1, 6, 5}
	fmt.Println("Before sorting", array)
	bubbleSort(array)
	fmt.Println("After sorting", array)
}
