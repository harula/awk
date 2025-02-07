package main

import "fmt"

func binarySearch(arr []int, key int) int {
	low := 0
	high := len(arr) - 1

	for low < high {
		mid := (low + high) / 2
		if key == arr[mid] {
			return mid
		} else if key < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	//low >= high
	if key > arr[low] {
		return low + 1
	} else {
		return low
	}
}

func binaryInsertionSort(arr []int) {
	for i := 1; i <= len(arr)-1; i++ {
		key := arr[i]
		loc := binarySearch(arr[:i], key)

		for j := i - 1; j >= loc; j-- {
			arr[j+1] = arr[j]
		}

		arr[loc] = key
	}
}

func main() {
	arr := []int{9, 5, 1, 4, 3}
	binaryInsertionSort(arr)
	fmt.Println("Sorted array:", arr)
}
