package main

import (
	"fmt"
)

func selectionSort(array []int) {
	lenght := len(array)
	for i := 0; i < lenght-1; i++ {
		minIndex := i
		for j := i + 1; j < lenght; j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}
		array[i], array[minIndex] = array[minIndex], array[i]

	}
}

func main() {
	arr := []int{12, 213, 45, 435, 43, 3, 45, 54, 54, 3, 2, 4, 5, 43, 3}
	fmt.Println("Array before sort:", arr)
	selectionSort(arr)
	fmt.Println("Array after sort:", arr)
}
