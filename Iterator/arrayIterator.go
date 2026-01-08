package Iterator

// ArrayIterator реализует итератор для массива целых чисел
type ArrayIterator struct {
	array []int
	index int
}

// NewArrayIterator создает новый итератор для массива
func NewArrayIterator(array []int) *ArrayIterator {
	return &ArrayIterator{
		array: array,
		index: 0,
	}
}

// HasNext проверяет, есть ли следующий элемент
func (ai *ArrayIterator) HasNext() bool {
	return ai.index < len(ai.array)
}

// Next возвращает следующий элемент
func (ai *ArrayIterator) Next() interface{} {
	if ai.HasNext() {
		value := ai.array[ai.index]
		ai.index++
		return value
	}
	return nil
}
