package collection_test

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/soulcodex/go-collections/collection"
)

type User struct {
	Id   string
	Name string
}

type UsersCollection struct {
	collection.Collection[*User]
}

func NewUsersCollection(users ...*User) *UsersCollection {
	return &UsersCollection{Collection: collection.NewCollection[*User](users...)}
}

var randomUsers = []*User{
	{Id: "1", Name: "John"},
	{Id: "2", Name: "Jane"},
	{Id: "3", Name: "Doe"},
	{Id: "4", Name: "Jack"},
	{Id: "5", Name: "Rick"},
	{Id: "6", Name: "Karl"},
}

func TestCollection(t *testing.T) {
	t.Run("it get the first item", func(t *testing.T) {
		users := NewUsersCollection(randomUsers...)
		user := users.First()

		assert.NotNil(t, user)
		assert.IsTypef(t, &User{}, user, "user should be a pointer to User")
		assert.Equal(t, "1", user.Id)
	})

	t.Run("it get the last item", func(t *testing.T) {
		users := NewUsersCollection(randomUsers...)
		user := users.Last()

		assert.NotNil(t, user)
		assert.IsTypef(t, &User{}, user, "user should be a pointer to User")
		assert.Equal(t, "6", user.Id)
	})

	t.Run("it get the items", func(t *testing.T) {
		users := NewUsersCollection(randomUsers...)
		items := users.Items()

		assert.NotNil(t, items)
		assert.IsTypef(t, []*User{}, items, "items should be a slice of pointers to User")
		assert.Equal(t, 6, users.Count())
		assert.Equal(t, 6, len(items))
	})

	t.Run("it check if the collection is empty", func(t *testing.T) {
		users := NewUsersCollection()
		assert.True(t, users.Empty())
	})

	t.Run("it check if the collection is not empty", func(t *testing.T) {
		users := NewUsersCollection(randomUsers...)
		assert.False(t, users.Empty())
	})

	t.Run("it filter the collection", func(t *testing.T) {
		users := NewUsersCollection(randomUsers...)
		filtered := users.Filter(func(user *User) bool {
			return user.Id == "1" || user.Id == "2"
		})
		users = NewUsersCollection(filtered.Items()...)

		assert.NotNil(t, filtered)
		assert.IsTypef(t, &UsersCollection{}, users, "users should be a pointer to UsersCollection")
		assert.Equal(t, 2, filtered.Count())
		assert.Equal(t, 2, len(filtered.Items()))
	})

	t.Run("it iterate over each item on the collection", func(t *testing.T) {
		users := NewUsersCollection(randomUsers...)
		var names []string

		err := users.Each(func(user *User) error {
			names = append(names, user.Name)
			return nil
		})

		assert.Nil(t, err)
		assert.Equal(t, users.Count(), len(names))
	})

	t.Run("it iterate over each item on the collection and return an error", func(t *testing.T) {
		errorMsg := "there is a user  with id equal to 3"
		users := NewUsersCollection(randomUsers...)

		err := users.Each(func(user *User) error {
			if user.Id == "3" {
				return errors.New(errorMsg)
			}

			user.Name = fmt.Sprintf("Mr. %s", user.Name)
			return nil
		})

		assert.NotNil(t, err)
		assert.Equal(t, errorMsg, err.Error())
	})

	t.Run("it search an item on the collection", func(t *testing.T) {
		users := NewUsersCollection(randomUsers...)
		user := users.Search(func(user *User) bool {
			return user.Id == "3"
		})

		assert.NotNil(t, user)
		assert.IsTypef(t, &User{}, user, "user should be a pointer to User")
		assert.Equal(t, "3", user.Id)
		assert.Equal(t, "Doe", user.Name)
	})
}
