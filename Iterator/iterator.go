package Iterator

// Iterator определяет интерфейс для итератора
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// Collection определяет интерфейс для коллекции
type Collection interface {
	CreateIterator() Iterator
}
