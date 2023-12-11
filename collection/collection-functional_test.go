package collection_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/soulcodex/go-collections/collection"
)

func TestCollectionFunctional(t *testing.T) {
	t.Run("it filter a collection", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result := collection.Filter[int](items, func(item int) bool {
			return item%2 == 0
		})

		assert.NotNil(t, result)
		assert.Equal(t, 2, len(result))
	})

	t.Run("it iterate over a collection", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		var result []int

		err := collection.Each[int](items, func(item int) error {
			result = append(result, item*2)
			return nil
		})

		assert.NoError(t, err)
		assert.Equal(t, 5, len(result))
		assert.Equal(t, 2, result[0])
		assert.Equal(t, 4, result[1])
	})

	t.Run("it search an item in a collection", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result := collection.Search[int](items, func(item int) bool {
			return item == 3
		})

		assert.Equal(t, 3, result)
	})

	t.Run("it search an item in a collection and return nil if not found", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result := collection.Search[int](items, func(item int) bool {
			return item == 6
		})

		assert.Equal(t, 0, result)
	})

	t.Run("it index a collection", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result, err := collection.IndexUnique[int, int](items, func(item int) int {
			return item
		})

		assert.NoError(t, err)
		assert.Equal(t, 5, len(result))
		assert.Equal(t, 1, result[1])
		assert.Equal(t, 2, result[2])
		assert.Equal(t, 3, result[3])
		assert.Equal(t, 4, result[4])
		assert.Equal(t, 5, result[5])
	})

	t.Run("it index a collection and fail on duplicated identifier", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5, 5}
		result, err := collection.IndexUnique[int, int](items, func(item int) int {
			return item
		})

		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("it index a collection by func regardless if there are duplicated elements", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5, 5}
		result := collection.IndexByFunc[int, int](items, func(item int) int {
			return item
		})

		assert.NotNil(t, result)
		assert.Equal(t, 5, len(result))
		assert.Equal(t, 1, result[1])
		assert.Equal(t, 5, result[5])
	})

	t.Run("it returns the first element of a collection", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result := collection.First[int](items)

		assert.Equal(t, 1, result)
	})

	t.Run("it returns the last element of a collection", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result := collection.Last[int](items)

		assert.Equal(t, 5, result)
	})

	t.Run("it returns the first element of a collection or nil if empty", func(t *testing.T) {
		var items []*int
		result := collection.First[*int](items)

		assert.Nil(t, result)
	})

	t.Run("it returns the last element of a collection or nil if empty", func(t *testing.T) {
		var items []*int
		result := collection.Last[*int](items)

		assert.Nil(t, result)
	})
}
