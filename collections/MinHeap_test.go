package collections

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)


func TestMinHeap(t *testing.T) {
	heap := &MinHeap{}
	heap.init()
	heap.Add(1)
	heap.Add(2)
	heap.Add(3)
	heap.Add(4)
	heap.Add(5)
	heap.Add(6)
	heap.Add(7)
	heap.Add(8)
	heap.Add(9)
	heap.Add(10)
	fmt.Println(heap.heap)

	heap2 := &MinHeap{}
	heap2.init()
	heap2.Add(15)
	heap2.Add(41)
	heap2.Add(66)
	heap2.Add(42)
	heap2.Add(18)
	fmt.Println(heap2.heap)

	assert.True(t, heap.Pop() == 1, fmt.Sprintf("heap: first pop should be %v", 1))
	assert.True(t, heap.Peek() == 2, fmt.Sprintf("heap: first peek should be %v", 2))

	assert.True(t, heap2.Pop() == 15, fmt.Sprintf("heap2: first pop should be %v", 1))
	assert.True(t, heap2.Peek() == 18, fmt.Sprintf("heap2: first peek should be %v", 2))

	fmt.Println(heap.heap)
	fmt.Println(heap2.heap)

}
