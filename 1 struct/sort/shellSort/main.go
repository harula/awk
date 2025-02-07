package main

import "fmt"

func shellSort(arr []int) {

	n := len(arr)
	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {

			key := arr[i]
			j := i
			for j >= gap && arr[j-gap] > key {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = key
		}
	}
}

func main() {
	arr := []int{9, 5, 1, 4, 3}
	shellSort(arr)
	fmt.Println("Sorted array:", arr)
}
