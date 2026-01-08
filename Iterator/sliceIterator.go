package Iterator

type SliceIterator struct {
	slice    []string
	position int
}

func NewSliceIterator(slice []string) *SliceIterator {
	return &SliceIterator{
		slice:    slice,
		position: 0,
	}
}

func (si *SliceIterator) HasNext() bool {
	return si.position < len(si.slice)
}

func (si *SliceIterator) Next() interface{} {
	if si.HasNext() {
		value := si.slice[si.position]
		si.position++
		return value
	}
	return nil
}
