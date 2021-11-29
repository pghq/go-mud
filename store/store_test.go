package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore_Insert(t *testing.T) {
	s := New()

	t.Run("bad value", func(t *testing.T) {
		_, err := s.Insert("test", func() {})
		assert.NotNil(t, err)
	})

	t.Run("success", func(t *testing.T) {
		item, err := s.Insert("test", "value")
		assert.Nil(t, err)
		assert.NotNil(t, item)

		item, err = s.Insert("test", "value")
		assert.Nil(t, err)
		assert.NotNil(t, item)

		item, err = s.Get(item.Id)
		assert.Nil(t, err)
		assert.NotNil(t, item)

		var value string
		err = item.Value(&value)
		assert.Nil(t, err)
		assert.Equal(t, "value", value)

		another, err := s.Insert("another", "value")
		assert.Nil(t, err)
		assert.NotNil(t, another)

		another, err = s.Get(another.Id)
		assert.Nil(t, err)
		assert.NotNil(t, another)

		assert.NotNil(t, item.Id, another.Id)
	})
}

func TestStore_Get(t *testing.T) {
	s := New()

	t.Run("not found", func(t *testing.T) {
		_, err := s.Get(0)
		assert.NotNil(t, err)
	})
}
