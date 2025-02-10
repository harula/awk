package main

import "fmt"

type Heap struct {
	arr []int
}

func NewHeap() *Heap {
	return &Heap{
		arr: []int{},
	}
}
func parrent(n int) int {
	return (n - 1) / 2
}

func leftChild(i int) int {
	return 2*i + 1
}
func rightChild(i int) int {
	return 2*i + 2
}
func (h *Heap) HeapUp(idx int) {
	for parrent(idx) >= 0 && h.arr[parrent(idx)] < h.arr[idx] {
		h.arr[parrent(idx)], h.arr[idx] = h.arr[idx], h.arr[parrent(idx)]
		idx = parrent(idx)
	}
}

func (h *Heap) HeapInsert(v int) {
	h.arr = append(h.arr, v)
	h.HeapUp(len(h.arr) - 1)
}

func (h Heap) Down(idx int) {
	last := len(h.arr) - 1
	lc := leftChild(idx)
	rc := rightChild(idx)

	childToCom := lc
	for lc <= last {
		if lc == last {
			childToCom = lc
		} else if h.arr[lc] > h.arr[rc] {
			childToCom = lc
		} else {
			childToCom = rc
		}

		if h.arr[idx] < h.arr[childToCom] {
			h.arr[idx], h.arr[childToCom] = h.arr[childToCom], h.arr[idx]
			idx = childToCom
			lc, rc = leftChild(idx), rightChild(idx)
		} else {
			return
		}
	}

}
func (h *Heap) Extract() int {

	if len(h.arr) == 0 {
		return -1
	}
	extracted := h.arr[0]
	lengh := len(h.arr) - 1
	if lengh >= 0 {
		h.arr[0] = h.arr[lengh]
		h.arr = h.arr[:lengh]
		h.Down(0)
	}

	return extracted
}

func heapUp(arr []int, k int) {
	for i := leftChild(k); i < len(arr); i = leftChild(i) {
		if i < len(arr)-1 && arr[i] < arr[i+1] {
			i++
		}
		if arr[k] >= arr[i] {
			break //跳出循环
		} else {
			arr[k], arr[i] = arr[i], arr[k] //交换数据
			k = i
		}
	}
	return
}

func buildHeap(arr []int) {
	lastIdx := len(arr) - 1
	lastParrent := (lastIdx - 1) / 2
	for i := lastParrent; i >= 0; i-- {
		heapUp(arr, i)
	}
}

func extract(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	extracted := arr[0]

	last := len(arr) - 1

	arr[0] = arr[last]
	arr = arr[:last]

	//fmt.Println("extreact", len(arr), arr)
	heapUp(arr, 0)

	return extracted
}

func main() {
	h := NewHeap()
	arr := []int{6, 1, 3, 7, 9, 12, 4, 0}
	for _, v := range arr {
		h.HeapInsert(v)
	}
	fmt.Println(h.arr)

	for i := 0; i < len(arr); i++ {
		extreact := h.Extract()
		fmt.Println(extreact, h.arr)
	}

	buildHeap(arr)
	lengh := len(arr)
	fmt.Println(arr)
	for i := 0; i < lengh; i++ {
		e := extract(arr)
		fmt.Println(e, arr)
	}

	arr2 := []int{6, 1, 3, 7, 9, 12, 4, 0}
	arr2 = arr2[:4]
	fmt.Println(arr2)
}
