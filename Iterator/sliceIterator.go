package Iterator

// SliceIterator реализует итератор для слайса строк
type SliceIterator struct {
	slice    []string
	position int
}

// NewSliceIterator создает новый итератор для слайса
func NewSliceIterator(slice []string) *SliceIterator {
	return &SliceIterator{
		slice:    slice,
		position: 0,
	}
}

// HasNext проверяет, есть ли следующий элемент
func (si *SliceIterator) HasNext() bool {
	return si.position < len(si.slice)
}

// Next возвращает следующий элемент
func (si *SliceIterator) Next() interface{} {
	if si.HasNext() {
		value := si.slice[si.position]
		si.position++
		return value
	}
	return nil
}

