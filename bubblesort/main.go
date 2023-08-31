package main

import (
	"fmt"
)

func bubbleSort(array []int) {
	lenght := len(array)
	for i := 0; i < lenght-1; i++ {
		swapped := false
		for j := 0; j < lenght-i-1; j++ {
			if array[j] > array[j+1] {
				//change elements
				array[j], array[j+1] = array[j+1], array[j]
				swapped = true
			}
		}

		// if list not change elements
		if !swapped {
			break
		}
	}

}

func main() {
	arr := []int{12, 213, 45, 435, 43, 3, 45, 54, 54, 3, 2, 4, 5, 43, 3}
	fmt.Println("Array before sort:", arr)
	bubbleSort(arr)
	fmt.Println("Array after sort:", arr)
}
