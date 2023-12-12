package collection_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/soulcodex/go-collections/collection"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) IncrementAge(incr int) error {
	p.Age += incr

	return nil
}

var people = []*Person{
	{"John", 20},
	{"Jane", 30},
	{"Jack", 40},
}

func TestCollectionFunctional(t *testing.T) {
	t.Run("it filter a collection of integers", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result := collection.Filter[int](items, func(item int) bool {
			return item%2 == 0
		})

		assert.NotNil(t, result)
		assert.Equal(t, 2, len(result))
	})

	t.Run("it filter a collection of struct pointers", func(t *testing.T) {
		result := collection.Filter[*Person](people, func(item *Person) bool {
			return item.Age > 30
		})

		assert.NotNil(t, result)
		assert.Equal(t, 1, len(result))
		assert.Equal(t, "Jack", result[0].Name)
	})

	t.Run("it iterate over a collection of integers", func(t *testing.T) {
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

	t.Run("it iterate over a collection and return an error", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}

		err := collection.Each[int](items, func(item int) error {
			return errors.New("error")
		})

		assert.Error(t, err)
	})

	t.Run("it iterate over a collection of struct pointers", func(t *testing.T) {
		err := collection.Each[*Person](people, func(item *Person) error {
			return item.IncrementAge(10)
		})

		assert.NoError(t, err)
		assert.Equal(t, 30, people[0].Age)
		assert.Equal(t, 40, people[1].Age)
		assert.Equal(t, 50, people[2].Age)
	})

	t.Run("it search an item in a collection", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result, found := collection.Search[int](items, func(item int) bool {
			return item == 3
		})

		assert.Equal(t, 3, result)
		assert.True(t, found)
	})

	t.Run("it search an item in a collection and return nil if not found", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result, found := collection.Search[int](items, func(item int) bool {
			return item == 6
		})

		assert.Equal(t, 0, result)
		assert.False(t, found)
	})

	t.Run("it search an item in a collection of struct pointers", func(t *testing.T) {
		result, found := collection.Search[*Person](people, func(item *Person) bool {
			return item.Name == "Jane"
		})

		assert.True(t, found)
		assert.Equal(t, "Jane", result.Name)
	})

	t.Run("it index a collection of integers", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		accessor, err := collection.IndexUnique[int, int](items, func(item int) int {
			return item
		}, func(item int) error {
			return nil
		})

		assert.NoError(t, err)
		assert.NotNil(t, accessor)

		item, found := accessor(1)
		assert.True(t, found)
		assert.Equal(t, 1, item)
	})

	t.Run("it index a collection and fail on duplicated identifier", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5, 5}
		accessor, err := collection.IndexUnique[int, int](items, func(item int) int {
			return item
		}, func(item int) error {
			return errors.New("error")
		})

		assert.Error(t, err)
		assert.Nil(t, accessor)
	})

	t.Run("it index a collection by func even if there are duplicated elements overriding the key", func(t *testing.T) {
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
		result, found := collection.First[int](items)

		assert.Equal(t, 1, result)
		assert.True(t, found)
	})

	t.Run("it returns the last element of a collection", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result, found := collection.Last[int](items)

		assert.Equal(t, 5, result)
		assert.True(t, found)
	})

	t.Run("it returns the first element of a collection or nil if empty", func(t *testing.T) {
		var items []*int
		result, found := collection.First[*int](items)

		assert.Nil(t, result)
		assert.False(t, found)
	})

	t.Run("it returns the last element of a collection or nil if empty", func(t *testing.T) {
		var items []*int
		result, found := collection.Last[*int](items)

		assert.Nil(t, result)
		assert.False(t, found)
	})
}
