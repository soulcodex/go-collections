package collection

import "fmt"

type UniqueCollection[K comparable, T any] struct {
	index   map[K]int
	indexer IdentifierFunc[K, T]
	Collection[T]
}

func buildIndex[K comparable, T any](indexer IdentifierFunc[K, T], items []T) (map[K]int, error) {
	index := make(map[K]int, len(items))

	for idx, item := range items {
		if _, ok := index[indexer(item)]; ok {
			return make(map[K]int, 0), fmt.Errorf("duplicated item found on unique collection")
		}

		index[indexer(item)] = idx
	}

	return index, nil
}

func NewUniqueCollection[K comparable, T any](indexer IdentifierFunc[K, T], items ...T) (UniqueCollection[K, T], error) {
	index, err := buildIndex[K, T](indexer, items)
	if err != nil {
		return UniqueCollection[K, T]{}, err
	}

	return UniqueCollection[K, T]{index: index, indexer: indexer, Collection: NewCollection[T](items...)}, nil
}

func (uc UniqueCollection[K, T]) Filter(fn FilterFunc[T]) (UniqueCollection[K, T], error) {
	filteredItems := uc.Collection.Filter(fn)

	index, err := buildIndex[K, T](uc.indexer, filteredItems.Items())
	if err != nil {
		return UniqueCollection[K, T]{}, err
	}

	return UniqueCollection[K, T]{index: index, indexer: uc.indexer, Collection: filteredItems}, nil
}

func (uc UniqueCollection[K, T]) Item(id K) T {
	if idx, ok := uc.index[id]; ok {
		return uc.items[idx]
	}

	return *new(T)
}
