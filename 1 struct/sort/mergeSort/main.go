package main

import "fmt"

var B []int

func Merge(A []int, low int, mid int, high int) {
	for i := low; i <= high; i++ {
		B[i] = A[i]
	}
	i := low
	j := mid + 1
	k := i
	for i <= mid && j <= high {
		if B[i] <= B[j] {
			A[k] = B[i]
			i++
		} else {
			A[k] = B[j]
			j++
		}
		k++
	}

	for i <= mid {
		A[k] = B[i]
		k++
		i++
	}

	for j <= high {
		A[k] = B[j]
		k++
		j++
	}

}

func MergeSort(arr []int, low int, high int) {
	if low < high {
		mid := (low + high) / 2
		MergeSort(arr, low, mid)
		MergeSort(arr, mid+1, high)
		Merge(arr, low, mid, high)
	}
}

func main() {
	arr := []int{49, 38, 65, 97, 76, 13, 27}
	B = make([]int, len(arr))
	MergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
