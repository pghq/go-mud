package frequency

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pghq/go-mud/internal"
)

func TestMap_Incr(t *testing.T) {
	t.Parallel()

	t.Run("job timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), 0)
		defer cancel()
		m := NewMap()
		m.incrJob(ctx)
	})
}

func TestMap_Max(t *testing.T) {
	t.Parallel()

	m := NewMap()

	t.Run("no data", func(t *testing.T) {
		t.Run("all", func(t *testing.T) {
			_, err := m.Max(internal.Query{})
			assert.NotNil(t, err)
		})

		t.Run("tags", func(t *testing.T) {
			_, err := m.MaxTag(internal.Query{})
			assert.NotNil(t, err)
		})
	})

	m.Incr([]byte("0"), "tag0").
		Incr([]byte("1"), "tag1").
		Incr([]byte("2"), "tag2").
		Incr([]byte("2"), "tag2").
		Incr([]byte("32"), "tag3", "tag2").
		Incr([]byte("3"), "tag3").
		Incr([]byte("3"), "tag3").
		Incr([]byte("3"), "tag3")

	m.Wait()

	t.Run("success", func(t *testing.T) {
		t.Run("all", func(t *testing.T) {
			keys, err := m.Max(internal.Query{Limit: 2})
			assert.Nil(t, err)
			assert.NotNil(t, keys)
			assert.Len(t, keys, 2)
			assert.Equal(t, [][]byte{[]byte("3"), []byte("2")}, keys)
		})

		t.Run("tags", func(t *testing.T) {
			keys, err := m.MaxTag(internal.Query{Tag: "tag3", Limit: 2})
			assert.Nil(t, err)
			assert.NotNil(t, keys)
			assert.Len(t, keys, 2)
			assert.Equal(t, [][]byte{[]byte("3"), []byte("32")}, keys)
		})
	})
}
