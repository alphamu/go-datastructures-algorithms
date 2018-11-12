package collections

import (
	"fmt"
)

type Node struct {
	next *Node
	prev *Node
	val  interface{}
}

type LinkedList struct {
	head  *Node
	tail  *Node
	count int

	Collection
	Iterable
}

type IteratorImpl struct {
	list   *LinkedList
	finger *Node
	Iterator
}

func (l *LinkedList) Put(val interface{}) {
	l.init()
	// Set prev as tail.prev
	// Set next as tail
	// Set tail.prev.next as insert
	insert := &Node{l.tail, l.tail.prev, val}
	l.tail.prev.next = insert
	l.tail.prev = insert
	l.count = l.count + 1
}

func (l *LinkedList) Get(index int) interface{} {
	l.init()
	finger := l.head
	i := 0
	for finger != l.tail && i <= index {
		finger = finger.next
		i = i + 1
	}
	if finger != l.tail {
		return finger.val
	}
	return nil
}

func (l *LinkedList) Insert(index int, val interface{}) {
	l.init()
	if index > l.count || index < 0 {
		// Handle error
		return
	}

	if index == l.count {
		// Insert at tail, same as Put
		l.Put(val)
	} else {
		finger := l.head
		for i := 0; finger != l.tail; i++ {
			if i == index {
				insert := &Node{finger.next, finger, val}
				finger.next.prev = insert
				finger.next = insert
				l.count = l.count + 1
				break
			}
			finger = finger.next
		}
	}
}

func (l *LinkedList) Remove(index int) interface{} {
	l.init()
	if index < 0 || index > l.count {
		// Handle error
		return nil
	}
	finger := l.head.next
	for i := 0; finger != l.tail; i++ {
		if i == index {
			finger.prev.next = finger.next
			finger.next.prev = finger.prev
			finger.next = nil
			finger.prev = nil
			l.count = l.count - 1
			return finger.val
		}
		finger = finger.next
	}
	return nil
}

func (l *LinkedList) Size() int {
	l.init()
	return l.count
}

func (l *LinkedList) Iterator() Iterator {
	return &IteratorImpl{list:l, finger:l.head}
}

func (it *IteratorImpl) HasNext() bool {
	return it.finger.next != nil && it.finger.next != it.list.tail
}

func (it *IteratorImpl) HasPrevious() bool {
	return it.finger.prev != nil && it.finger.prev != it.list.head
}

func (it *IteratorImpl) Next() interface{} {
	if it.finger == it.list.tail {
		return nil
	}

	it.finger = it.finger.next
	val := it.finger.val

	return val
}

func (it *IteratorImpl) Previous() interface{} {
	if it.finger == it.list.head {
		return nil
	}

	it.finger = it.finger.prev
	val := it.finger.val

	return val
}

func (it *IteratorImpl) Remove() {
	if it.finger != nil && it.finger != it.list.head && it.finger != it.list.tail {
		prev := it.finger.prev
		it.finger.next.prev = it.finger.prev
		it.finger.prev.next = it.finger.next
		it.finger.next = nil
		it.finger.prev = nil
		it.finger = prev
		it.list.count = it.list.count - 1
	}
}

func (l *LinkedList) init() {
	if l.head == nil && l.tail == nil {
		fmt.Println("Init")
		head := &Node{nil, nil, "head"}
		tail := &Node{nil, nil, "tail"}
		head.next = tail
		tail.prev = head
		l.head = head
		l.tail = tail
	}
}

func (l *LinkedList) Println() {
	finger := l.head
	index := 0
	for finger != nil {
		if finger == l.head {
			fmt.Printf("H: {val: %v}\n", finger.val)
		} else  if finger == l.tail {
			fmt.Printf("T: {val: %v}\n", finger.val)
		} else {
			fmt.Printf("%v: {val: %v}\n", index, finger.val)
			index = index + 1
		}
		finger = finger.next
	}
	fmt.Printf("Size: %v\n", l.count)
}
