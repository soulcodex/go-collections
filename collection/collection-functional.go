package collection

func Filter[T any](collection []T, fn FilterFunc[T]) []T {
	result := make([]T, 0)
	for _, item := range collection {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}

func Each[T any](collection []T, fn ItemFunc[T]) error {
	for _, item := range collection {
		if err := fn(item); err != nil {
			return err
		}
	}

	return nil
}

func Search[T any](collection []T, fn FilterFunc[T]) (T, bool) {
	var (
		needle T
		found  bool
	)

	for _, item := range collection {
		if fn(item) {
			needle = item
			found = true
			break
		}
	}

	return needle, found
}

func IndexUnique[K comparable, T any](collection []T, idFunc IdentifierFunc[K, T], onError NonUniqueIndexFunc[T]) (IndexAccessFunc[K, T], error) {
	index, err := uniqueIndex[K, T](collection, idFunc, onError)
	if err != nil {
		return nil, err
	}

	return indexAccessFunc(collection, index), nil
}

func IndexByFunc[K comparable, T any](collection []T, idFunc IdentifierFunc[K, T]) map[K]T {
	indexed := make(map[K]T)
	for _, item := range collection {
		indexed[idFunc(item)] = item
	}

	return indexed
}

func First[T any](collection []T) (T, bool) {
	var needle T
	if len(collection) > 0 {
		needle = collection[0]
		return needle, true
	}
	return needle, false
}

func Last[T any](collection []T) (T, bool) {
	var needle T
	if len(collection) > 0 {
		needle = collection[len(collection)-1]
		return needle, true
	}
	return needle, false
}
