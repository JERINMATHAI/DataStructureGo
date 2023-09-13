package main

import "fmt"

func mergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	middle := len(array) / 2
	leftArray := mergeSort(array[:middle])
	rightArray := mergeSort(array[middle:])
	return merge(leftArray, rightArray)
}

func merge(leftArray, rightArray, []int) []int {
	result := make([]int, 0, len(leftArray+len(rightArray)))
	i, j := 0, 0
}

func main() {
	array := []int{3, 4, 5, 1, 9, 6}
	fmt.Println("Before: ", array)
	mergeSort(array)
	fmt.Println("After: ", array)
}
