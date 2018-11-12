package collections

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func initList() LinkedList {
	ll := LinkedList{}
	ll.Put(0)
	ll.Put(1)
	ll.Put(2)
	ll.Put(3)
	ll.Put(4)

	return ll
}

func TestPutAndGet(t *testing.T) {
	ll := initList()
	assert.Equal(t, 0, ll.Get(0), "Index 0 should be 0")
	assert.Equal(t, 1, ll.Get(1), "Index 1 should be 1")
	assert.Equal(t, 2, ll.Get(2), "Index 2 should be 2")
	assert.Equal(t, 3, ll.Get(3), "Index 3 should be 3")
	assert.Equal(t, 4, ll.Get(4), "Index 4 should be 4")
	// Assert size
	assert.Equal(t, 5, ll.count, "Size should be 5")
}


func TestBuildLLWithInsert(t *testing.T) {
	ll := LinkedList{}
	ll.Insert(0, 4)
	ll.Insert(0, 3)
	ll.Insert(0, 2)
	ll.Insert(0, 1)
	ll.Insert(0, 0)

	assert.Equal(t, 0, ll.Get(0), "Index 0 should be 0")
	assert.Equal(t, 1, ll.Get(1), "Index 1 should be 1")
	assert.Equal(t, 2, ll.Get(2), "Index 2 should be 2")
	assert.Equal(t, 3, ll.Get(3), "Index 3 should be 3")
	assert.Equal(t, 4, ll.Get(4), "Index 4 should be 4")
	// Assert size
	assert.Equal(t, 5, ll.count, "Size should be 5")

	ll = LinkedList{}
	ll.Insert(0, 0)
	ll.Insert(1, 1)
	ll.Insert(2, 2)
	ll.Insert(3, 3)
	ll.Insert(4, 4)

	assert.Equal(t, 0, ll.Get(0), "Index 0 should be 0")
	assert.Equal(t, 1, ll.Get(1), "Index 1 should be 1")
	assert.Equal(t, 2, ll.Get(2), "Index 2 should be 2")
	assert.Equal(t, 3, ll.Get(3), "Index 3 should be 3")
	assert.Equal(t, 4, ll.Get(4), "Index 4 should be 4")
	// Assert size
	assert.Equal(t, 5, ll.count, "Size should be 5")
}

func TestInsertAtIndex0(t *testing.T) {
	ll := initList()

	ll.Insert(0, 100) // Insert at 0, same as prepend.
	// 0 should be 100, everything else moved up 1 spot
	assert.Equal(t, 100, ll.Get(0), "Index 0 should be 100")
	assert.Equal(t, 0, ll.Get(1), "Index 1 should be 0")
	assert.Equal(t, 1, ll.Get(2), "Index 2 should be 1")
	assert.Equal(t, 2, ll.Get(3), "Index 3 should be 2")
	assert.Equal(t, 3, ll.Get(4), "Index 4 should be 3")
	// Assert size
	assert.Equal(t, 6, ll.count, "Size should be 6")
}

func TestInsertAtEnd(t *testing.T) {
	ll := initList()

	ll.Insert(5, 500) // Insert at position 5, which is the length of the list, same as append
	assert.Equal(t, 500, ll.Get(5), "Index 5 should be 500")
	// These should not change position
	assert.Equal(t, 0, ll.Get(0), "Index 0 should be 0")
	assert.Equal(t, 1, ll.Get(1), "Index 1 should be 1")
	assert.Equal(t, 2, ll.Get(2), "Index 2 should be 2")
	assert.Equal(t, 3, ll.Get(3), "Index 3 should be 3")
	assert.Equal(t, 4, ll.Get(4), "Index 4 should be 4")
	// Assert size
	assert.Equal(t, 6, ll.count, "Size should be 6")

}

func TestInsertInMiddle(t *testing.T) {
	ll := initList()

	ll.Insert(3, 300)
	assert.Equal(t, 300, ll.Get(3), "Index 3 should be 300")
	// These should not change position
	assert.Equal(t, 0, ll.Get(0), "Index 0 should be 0")
	assert.Equal(t, 1, ll.Get(1), "Index 1 should be 1")
	assert.Equal(t, 2, ll.Get(2), "Index 2 should be 2")
	// This should move up ONE position
	assert.Equal(t, 4, ll.Get(5), "Index 5 should be 4")
	// Assert size
	assert.Equal(t, 6, ll.count, "Size should be 6")

}

func TestRemove(t *testing.T) {
	ll := initList()

	assert.Equal(t, 0, ll.Get(0), "Index 0 should be 0")
	ll.Remove(0)
	assert.Equal(t, 1, ll.Get(0), "Index 0 should be 1")
	// Assert size
	assert.Equal(t, 4, ll.count, "Size should be 4")


	assert.Equal(t, 4, ll.Get(ll.count - 1), "Last item should be 4")
	ll.Remove(ll.count - 1)
	assert.Equal(t, 3, ll.Get(ll.count - 1), "Last item should be 3")
	// Assert size
	assert.Equal(t, 3, ll.count, "Size should be 3")

}

func TestIterator(t *testing.T) {
	ll := initList()
	it := ll.Iterator()
	var count = 0
	fmt.Println("\n\nTestIterator Next()")
	for it.HasNext() {
		next := it.Next()
		fmt.Println("Index:", count, "Value:", next)
		assert.Equal(t, count, next, fmt.Sprintf("Index %v should be %v", count, count))
		count = count + 1
	}
	it.Next() // force it to the tail
	count = count - 1
	fmt.Println("\n\nTestIterator Previous()")
	for it.HasPrevious() {
		prev := it.Previous()
		fmt.Println("Index:", count, "Value:", prev)
		assert.Equal(t, count, prev, fmt.Sprintf("Index %v should be %v", count, count))
		count = count - 1
	}
}

func TestIteratorRemove(t *testing.T) {
	ll := initList()
	it := ll.Iterator()
	fmt.Println("\n\nTestIterator Next()")
	assert.Equal(t, 2, ll.Get(2), "At index 2, should be value 2")
	assert.Equal(t, 5, ll.Size(), "List size should be 5")
	for it.HasNext() {
		next := it.Next()
		if next == 2 {
			it.Remove()
		}
	}

	assert.Equal(t, 3, ll.Get(2), "At index 2, should be value 3")
	assert.Equal(t, 4, ll.Size(), "List size should be 4")
}

