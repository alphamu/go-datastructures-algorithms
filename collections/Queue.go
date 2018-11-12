package collections

import "fmt"

type Queue struct {
	head *QueueNode
	tail *QueueNode
}

type QueueNode struct {
	val interface{}
	prev *QueueNode
	next *QueueNode
}

func (q * Queue) init() {
	if q.head == nil {
		q.head = &QueueNode{}
		q.tail = &QueueNode{}
		q.head.next = q.tail
		q.tail.prev = q.head
	}
}

func (q * Queue) push(val interface{}) {
	if q.head == nil || q.tail == nil || q.head.next == q.tail {
		q.init()
		newNode := &QueueNode{val, q.head, q.tail}
		q.head.next = newNode
		q.tail.prev = newNode
	} else {
		newNode := &QueueNode{val, q.head, q.head.next}
		q.head.next.prev = newNode
		q.head.next = newNode
	}
}

func (q * Queue) pop() interface{} {
	if q.tail == nil || q.tail.prev == q.head {
		return nil
	}
	finger := q.tail.prev
	finger.prev.next = q.tail
	q.tail.prev = finger.prev
	finger.next = nil
	finger.prev = nil
	return finger.val
}

func (q * Queue) printQueue() {
	finger := q.head
	for finger.next != q.tail {
		fmt.Print(finger.next.val, ",")
		finger = finger.next
	}
}
