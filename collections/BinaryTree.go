package collections

import "fmt"

type TreeItem struct {
	val    interface{}
	left   *TreeItem
	right  *TreeItem
	parent *TreeItem
}

type BinaryTree struct {
	root  *TreeItem
	count int
}

func (t *BinaryTree) init(val interface{}) {
	if t.root == nil {
		t.root = &TreeItem{val, nil, nil, nil}
		t.count = 1
	}
}

func (t *BinaryTree) getRoot() interface{} {
	return t.root
}

func (t *BinaryTree) Put(val interface{}) {
	if t.root == nil {
		t.init(val)
	} else {
		insert := &TreeItem{val, nil, nil, nil}
		finger := t.root
		for finger != nil {
			if insert.val.(int) < finger.val.(int) {
				// Go left
				if finger.left == nil {
					insert.parent = finger
					finger.left = insert
					t.count = t.count + 1
					break
				} else {
					finger = finger.left
				}

			} else {
				// Go right
				if finger.right == nil {
					insert.parent = finger
					finger.right = insert
					t.count = t.count + 1
					break
				} else {
					finger = finger.right
				}
			}
		}
	}
}

type ResultHolder struct {
	result []interface{}
	index  int
}

func (t *BinaryTree) GetAscending() []interface{} {
	fmt.Println("\nCount ", t.count)
	holder := &ResultHolder{
		make([]interface{}, t.count),
		0,
	}
	result := t.getAscending(t.root, holder).result
	return result
}

func (t *BinaryTree) getAscending(node *TreeItem, holder *ResultHolder) *ResultHolder {
	if node.left != nil {
		t.getAscending(node.left, holder)
	}

	holder.result[holder.index] = node.val
	holder.index = holder.index + 1

	if node.right != nil {
		t.getAscending(node.right, holder)
	}
	return holder
}

func (t *BinaryTree) GetDescending() []interface{} {
	fmt.Println("\nCount ", t.count)
	holder := &ResultHolder{
		make([]interface{}, t.count),
		0,
	}
	result := t.getDescending(t.root, holder).result
	return result
}

func (t *BinaryTree) getDescending(node *TreeItem, holder *ResultHolder) *ResultHolder {
	if node.right != nil {
		t.getDescending(node.right, holder)
	}

	holder.result[holder.index] = node.val
	holder.index = holder.index + 1

	if node.left != nil {
		t.getDescending(node.left, holder)
	}
	return holder
}

func (t *BinaryTree) PrintAscending(node *TreeItem) {
	if node == nil {
		node = t.root
	}
	if node.left != nil {
		t.PrintAscending(node.left)
	}
	fmt.Print(node.val, ",")
	if node.right != nil {
		t.PrintAscending(node.right)
	}
}

func (t *BinaryTree) PrintDescending(node *TreeItem) {
	if node == nil {
		node = t.root
	}
	if node.right != nil {
		t.PrintDescending(node.right)
	}
	fmt.Print(node.val, ",")
	if node.left != nil {
		t.PrintDescending(node.left)
	}
}
