package collection

type IndexAccessFunc[K comparable, T any] func(id K) (T, bool)
type NonUniqueIndexFunc[T any] func(T) error

func uniqueIndex[K comparable, T any](collection []T, idFunc IdentifierFunc[K, T], onError NonUniqueIndexFunc[T]) (map[K]int, error) {
	index := make(map[K]int, len(collection))

	for pos, item := range collection {
		id := idFunc(item)
		if _, ok := index[id]; ok {
			return nil, onError(item)
		}

		index[id] = pos
	}

	return index, nil
}

func indexAccessFunc[K comparable, T any](collection []T, index map[K]int) IndexAccessFunc[K, T] {
	return func(id K) (T, bool) {
		var (
			needle T
			found  = false
		)

		if pos, ok := index[id]; ok {
			needle, found = collection[pos], true
		}

		return needle, found
	}
}
