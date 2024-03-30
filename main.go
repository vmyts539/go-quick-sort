package main

import (
	"fmt"
	"reflect"
)

func partition(arr []int, lo int, hi int) int {
	pivot := arr[hi]
	idx := lo - 1

	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			idx++
			arr[i], arr[idx] = arr[idx], arr[i]
		}
	}

	idx++
	arr[hi], arr[idx] = arr[idx], arr[hi]

	return idx
}

func quickSort(arr []int, lo int, hi int) {
	if lo >= hi {
		return
	}

	pivotIdx := partition(arr, lo, hi)
	quickSort(arr, lo, pivotIdx-1)
	quickSort(arr, pivotIdx+1, hi)
}

func main() {
	arr := []int{9, 3, 7, 4, 69, 420, 42}
	res := []int{3, 4, 7, 9, 42, 69, 420}

	quickSort(arr, 0, len(arr)-1)

	if reflect.DeepEqual(arr, res) {
		fmt.Printf("Sorted: %v", arr)
	} else {
		fmt.Printf("Failed to sort: %v", arr)
	}
}
