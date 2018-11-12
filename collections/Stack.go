package collections

import "fmt"

type Stack struct {
	head *StackNode
}

type StackNode struct {
	val interface{}
	next *StackNode
}

func (s * Stack) init() {
	s.head = &StackNode{}
}

func (s * Stack) push(val interface{})  {
	newHead := &StackNode{}
	newHead.val = val
	if s.head ==  nil {
		s.init()
		s.head.next = newHead
	} else {
		finger := s.head
		newHead.next = finger.next
		finger.next = newHead
	}
}

func (s * Stack) pop() interface{} {
	if s.head.next != nil {
		finger := s.head.next
		s.head.next = finger.next
		finger.next = nil
		return finger.val
	} else {
		return nil
	}
}

func (s * Stack) printStack() {
	finger := s.head
	for finger.next != nil {
		fmt.Print(finger.next.val, ",")
		finger = finger.next
	}
}
