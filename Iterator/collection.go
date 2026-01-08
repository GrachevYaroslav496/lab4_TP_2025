package Iterator

type IntArrayCollection struct {
	items []int
}

func NewIntArrayCollection(items []int) *IntArrayCollection {
	return &IntArrayCollection{
		items: items,
	}
}

func (c *IntArrayCollection) CreateIterator() Iterator {
	return NewArrayIterator(c.items)
}

type StringSliceCollection struct {
	items []string
}

func NewStringSliceCollection(items []string) *StringSliceCollection {
	return &StringSliceCollection{
		items: items,
	}
}

func (c *StringSliceCollection) CreateIterator() Iterator {
	return NewSliceIterator(c.items)
}
