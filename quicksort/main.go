package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	left, right := 0, len(array)-1
	pivotIndex := rand.Int() % len(array)

	array[pivotIndex], array[right] = array[right], array[pivotIndex]

	for i, _ := range array {
		if array[i] < array[right] {
			array[i], array[left] = array[left], array[i]
			left++
		}
	}

	array[left], array[right] = array[right], array[left]

	quickSort(array[:left])
	quickSort(array[left+1:])

	return array
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	arr := []int{10, 5, 3, 9, 2, 11, 6, 8, 7}
	fmt.Println("Unsorted:", arr)
	fmt.Println("Sorted:", quickSort(arr))
}
