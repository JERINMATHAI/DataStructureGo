package main

func main() {
	var temp int
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 3, 4}
	for i := 0; i < len(array2); i++ {
		array1 = append(array1, array2[i])
	}

	size := len(array1)
	for i := 0; i < size-1; i++ {
		for j := 0; j < size; j++ {
			if array1[i] > array1[j] {
				temp = array1[i]
				array1[i] = array1[j]
				array1[j] = temp
			}
		}
	}
}
