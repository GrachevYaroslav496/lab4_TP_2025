package Iterator

// IntArrayCollection представляет коллекцию целых чисел
type IntArrayCollection struct {
	items []int
}

// NewIntArrayCollection создает новую коллекцию целых чисел
func NewIntArrayCollection(items []int) *IntArrayCollection {
	return &IntArrayCollection{
		items: items,
	}
}

// CreateIterator создает итератор для коллекции
func (c *IntArrayCollection) CreateIterator() Iterator {
	return NewArrayIterator(c.items)
}

// StringSliceCollection представляет коллекцию строк
type StringSliceCollection struct {
	items []string
}

// NewStringSliceCollection создает новую коллекцию строк
func NewStringSliceCollection(items []string) *StringSliceCollection {
	return &StringSliceCollection{
		items: items,
	}
}

// CreateIterator создает итератор для коллекции
func (c *StringSliceCollection) CreateIterator() Iterator {
	return NewSliceIterator(c.items)
}
