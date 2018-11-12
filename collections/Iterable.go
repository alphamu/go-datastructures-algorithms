package collections

type Iterable interface {
	Iterator() Iterator
}

type Iterator interface {
	Next() interface{}
	Previous() interface{}
	HasNext() bool
	HasPrevious() bool
	Remove()
}