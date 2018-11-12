package collections

type Collection interface {
	Put(val interface{})
	Insert(index int, val interface{})
	Get(index int) interface {}
	Remove(index int) interface{}
	Size() int
}
