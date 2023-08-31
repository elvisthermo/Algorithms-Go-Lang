package main

import (
	"fmt"
)

func insertionSort(array []int) {
	lenght := len(array)
	for i := 1; i < lenght; i++ {
		key := array[i]
		j := i - 1
		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = key
	}
}

func main() {
	arr := []int{12, 213, 45, 435, 43, 3, 45, 9, 8, 7, 5, 4, 3, 1}
	fmt.Println("Array before sort:", arr)
	insertionSort(arr)
	fmt.Println("Array after sort:", arr)
}
