package mud

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrendService_Commit(t *testing.T) {
	c := NewClassifier()

	t.Run("bad request", func(t *testing.T) {
		err := c.Insert("", nil, nil)
		assert.NotNil(t, err)
	})

	t.Run("bad value", func(t *testing.T) {
		err := c.Insert("foo", func() {}, []float64{2, 3})
		assert.NotNil(t, err)
	})
}

func TestTrendService_Search(t *testing.T) {
	c := NewClassifier()

	t.Run("no data", func(t *testing.T) {
		_, err := c.Search([]float64{7, 3})
		assert.NotNil(t, err)
	})

	_ = c.Insert("foo1", "bar1", []float64{2, 3})
	_ = c.Insert("foo2", "bar2", []float64{5, 4})
	_ = c.Insert("foo3", "bar3", []float64{4, 7})
	_ = c.Insert("foo4", "bar4", []float64{8, 1})
	_ = c.Insert("foo5", "bar5", []float64{7, 2})
	_ = c.Insert("foo6", "bar6", []float64{9, 6})
	c.Wait()

	t.Run("success", func(t *testing.T) {
		c, err := c.Search([]float64{7, 3})
		assert.Nil(t, err)
		assert.NotNil(t, c)

		values := []string{
			"bar5",
			"bar2",
		}

		i := 0
		for c.Next() {
			var value string
			_ = c.Decode(&value)
			assert.Equal(t, values[i], value)
			i++
		}

		assert.Equal(t, len(values), i)
	})
}
