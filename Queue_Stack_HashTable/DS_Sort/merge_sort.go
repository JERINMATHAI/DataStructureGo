package main

import "fmt"

func mergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	middle := len(array) / 2
	leftArray := mergeSort(array[:middle])
	rightArray := mergeSort(array[middle:])

	return merge(rightArray, leftArray)
}

func merge(leftArray, rightArray []int) []int {
	result := make([]int, 0, len(leftArray)+len(rightArray))
	i, j := 0, 0

	for i < len(leftArray) && j < len(rightArray) {
		if leftArray[i] <= rightArray[j] {
			result = append(result, leftArray[i])
			i++
		} else {
			result = append(result, rightArray[j])
			j++
		}
	}
	for i < len(leftArray) {
		result = append(result, leftArray[i])
		i++
	}
	for j < len(rightArray) {
		result = append(result, rightArray[j])
		j++
	}
	return result

}

func main() {
	fmt.Println("Merge sort")
	fmt.Println("```````````")
	array := []int{9, 10, 1, 2, 3, 7, 0}
	fmt.Println("Array before sorting: ", array)
	sortedArray := mergeSort(array)
	fmt.Println("After Sorting", sortedArray)
}
