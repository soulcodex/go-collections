package collection_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/soulcodex/go-collections/collection"
)

type UsersUniqueCollection struct {
	collection.UniqueCollection[string, *User]
}

func NewUsersUniqueCollection(users ...*User) (*UsersUniqueCollection, error) {
	items, err := collection.NewUniqueCollection[string, *User](func(user *User) string {
		return user.Id
	}, users...)

	if err != nil {
		return nil, err
	}

	return &UsersUniqueCollection{UniqueCollection: items}, nil
}

func TestUniqueCollection(t *testing.T) {
	t.Run("it get the item by id", func(t *testing.T) {
		users, _ := NewUsersUniqueCollection(randomUsers...)
		user := users.Item("1")

		assert.NotNil(t, user)
		assert.IsTypef(t, &User{}, user, "user should be a pointer to User")
		assert.Equal(t, "1", user.Id)
	})

	t.Run("it get the items", func(t *testing.T) {
		users, _ := NewUsersUniqueCollection(randomUsers...)
		items := users.Items()

		assert.NotNil(t, items)
		assert.IsTypef(t, []*User{}, items, "items should be a slice of pointers to User")
		assert.Equal(t, 6, users.Count())
		assert.Equal(t, 6, len(items))
	})

	t.Run("it filter the unique collection items", func(t *testing.T) {
		users, _ := NewUsersUniqueCollection(randomUsers...)
		items, _ := users.Filter(func(user *User) bool {
			return user.Id == "1" || user.Id == "2"
		})
		users, _ = NewUsersUniqueCollection(items.Items()...)

		assert.NotNil(t, items)
		assert.IsType(t, collection.UniqueCollection[string, *User]{}, items)
		assert.IsType(t, []*User{}, users.Items())
		assert.Equal(t, 2, len(users.Items()))
	})
}
