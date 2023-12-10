package collection

type FilterFunc[T any] func(T) bool
type ItemFunc[T any] func(item T) error
type IdentifierFunc[K comparable, T any] func(T) K

type Collection[T any] struct {
	items []T
}

func NewCollection[T any](items ...T) Collection[T] {
	return Collection[T]{items: items}
}

func (c Collection[T]) Items() []T {
	return c.items
}

func (c Collection[T]) Count() int {
	return len(c.items)
}

func (c Collection[T]) Empty() bool {
	return len(c.items) == 0
}

func (c Collection[T]) First() T {
	if len(c.items) > 0 {
		return c.items[0]
	}

	return *new(T)
}

func (c Collection[T]) Last() T {
	if len(c.items) > 0 {
		return c.items[len(c.items)-1]
	}

	return *new(T)
}

func (c Collection[T]) Filter(fn FilterFunc[T]) Collection[T] {
	items := make([]T, 0)

	for _, item := range c.items {
		if fn(item) {
			items = append(items, item)
		}
	}

	return NewCollection[T](items...)
}

func (c Collection[T]) Each(fn ItemFunc[T]) error {
	beforeChange := c.Items()
	for _, item := range c.items {
		if err := fn(item); err != nil {
			c.items = beforeChange
			return err
		}
	}

	return nil
}

func (c Collection[T]) Search(fn FilterFunc[T]) T {
	for _, item := range c.items {
		if fn(item) {
			return item
		}
	}

	return *new(T)
}
