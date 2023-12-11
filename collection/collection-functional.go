package collection

import (
	"fmt"
)

func uniqueIndex[K comparable, T any](collection []T, idFunc IdentifierFunc[K, T]) (map[K]T, error) {
	index, items := make(map[K]struct{}, len(collection)), make(map[K]T, len(collection))

	for _, item := range collection {
		id := idFunc(item)
		if _, ok := index[id]; ok {
			return nil, fmt.Errorf("duplicated identifier")
		}

		index[id] = struct{}{}
		items[id] = item
	}

	return items, nil
}

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

func Search[T any](collection []T, fn FilterFunc[T]) T {
	var needle T

	for _, item := range collection {
		if fn(item) {
			needle = item
			break
		}
	}

	return needle
}

func IndexUnique[K comparable, T any](collection []T, idFunc IdentifierFunc[K, T]) (map[K]T, error) {
	items, err := uniqueIndex[K, T](collection, idFunc)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func IndexByFunc[K comparable, T any](collection []T, idFunc IdentifierFunc[K, T]) map[K]T {
	indexed := make(map[K]T)
	for _, item := range collection {
		indexed[idFunc(item)] = item
	}

	return indexed
}

func First[T any](collection []T) T {
	var needle T
	if len(collection) > 0 {
		needle = collection[0]
		return needle
	}
	return needle
}

func Last[T any](collection []T) T {
	var needle T
	if len(collection) > 0 {
		needle = collection[len(collection)-1]
		return needle
	}
	return needle
}
