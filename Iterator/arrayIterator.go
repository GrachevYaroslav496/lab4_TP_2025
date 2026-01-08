package Iterator

type ArrayIterator struct {
	array []int
	index int
}

func NewArrayIterator(array []int) *ArrayIterator {
	return &ArrayIterator{
		array: array,
		index: 0,
	}
}

func (ai *ArrayIterator) HasNext() bool {
	return ai.index < len(ai.array)
}

func (ai *ArrayIterator) Next() interface{} {
	if ai.HasNext() {
		value := ai.array[ai.index]
		ai.index++
		return value
	}
	return nil
}
