package Iterator

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type Collection interface {
	CreateIterator() Iterator
}
