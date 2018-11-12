package collections

import "math"

const ONE = float64(1)
const TWO = float64(2)

type Heap interface {
	Add(val int)
	Peek() int
	Pop() int
}

type MinHeap struct {
	heap  []int
	count int

	Heap
}

func (h *MinHeap) init() {
	if h.heap == nil {
		h.heap = []int{-1,-1,-1,-1,-1,-1,-1}
		h.count = 0
	}
}

func (h *MinHeap) Add(val int) {
	// If the array is full, double it
	if h.count == len(h.heap) {
		temp := make([]int, h.count * 2)
		copy(temp[:], h.heap)
		h.heap = temp
	}
	index := h.count
	parentIndex := int(math.Floor((float64(index) - ONE) / TWO))

	for index > 0 && val < h.heap[parentIndex] {
		h.heap[index] = h.heap[parentIndex]
		index = parentIndex
		parentIndex = int(math.Floor((float64(index) - ONE) / TWO))
	}

	h.heap[index] = val
	h.count = h.count + 1
}

func (h *MinHeap) Peek() int {
	if h.count > 0 {
		return h.heap[0]
	} else {
		return -1
	}
}

func (h * MinHeap) Pop() int {
	temp := h.Peek()
	if temp != -1 {
		// Replace the head with the "largest" value
		h.heap[0] = h.heap[h.count - 1]
		h.heap[h.count - 1] = -1
		// Effectively ignore the last value
		// Even though the array is the same size
		h.count = h.count - 1
		h.balance(0)
	}
	return temp
}

func (h * MinHeap) balance(index int) {
	indexOfSmallest := index
	indexOfLeftChild := (index * 2) + 1
	indexOfRightChild := (index * 2) + 2

	if indexOfLeftChild < h.count && h.heap[indexOfLeftChild] < h.heap[indexOfSmallest] {
		indexOfSmallest = indexOfLeftChild
	}
	if indexOfRightChild < h.count && h.heap[indexOfRightChild] < h.heap[indexOfSmallest] {
		indexOfSmallest = indexOfRightChild
	}
	if indexOfSmallest != index {
		//swap
		h.swap(index, indexOfSmallest)
		h.balance(indexOfSmallest)
	}
}

func (h * MinHeap) swap(indexOne int, indexTwo int) {
	temp := h.heap[indexOne]
	h.heap[indexOne] = h.heap[indexTwo]
	h.heap[indexTwo] = temp
}

