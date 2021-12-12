package mud

import (
	"testing"

	"github.com/pghq/go-ark"
	"github.com/stretchr/testify/assert"
)

func TestGraph_Plot(t *testing.T) {
	t.Parallel()

	g := New(Mapper(ark.New()))

	t.Run("bad request", func(t *testing.T) {
		err := g.Plot(nil, nil, nil)
		assert.NotNil(t, err)
	})

	t.Run("bad value", func(t *testing.T) {
		err := g.Plot([]byte("foo"), func() {}, []float64{2, 3})
		assert.NotNil(t, err)
	})
}

func TestGraph_Nearest(t *testing.T) {
	t.Parallel()

	g := New()

	t.Run("no data", func(t *testing.T) {
		var values []string
		err := g.Nearest([]float64{7, 3}, &values)
		assert.NotNil(t, err)
	})

	_ = g.Plot([]byte("foo1"), "bar1", []float64{2, 3})
	_ = g.Plot([]byte("foo2"), "bar2", []float64{5, 4})
	_ = g.Plot([]byte("foo3"), "bar3", []float64{4, 7})
	_ = g.Plot([]byte("foo4"), "bar4", []float64{8, 1})
	_ = g.Plot([]byte("foo5"), "bar5", []float64{7, 2})
	_ = g.Plot([]byte("foo6"), "bar6", []float64{9, 6})
	g.Wait()

	t.Run("not a slice", func(t *testing.T) {
		var value int
		err := g.Nearest([]float64{7, 3}, &value)
		assert.NotNil(t, err)
	})

	t.Run("slice of bad values", func(t *testing.T) {
		var values []int
		err := g.Nearest([]float64{7, 3}, &values)
		assert.NotNil(t, err)
	})

	t.Run("success", func(t *testing.T) {
		var values []string
		err := g.Nearest([]float64{7, 3}, &values)
		assert.Nil(t, err)

		values = nil
		err = g.Nearest([]float64{7, 3}, &values)
		assert.Nil(t, err)
		assert.NotNil(t, values)
		assert.Len(t, values, 2)
		assert.Equal(t, "bar5", values[0])
		assert.Equal(t, "bar2", values[1])
	})
}

func TestGraph_Frequency(t *testing.T) {
	t.Parallel()

	g := New()

	t.Run("nil value", func(t *testing.T) {
		err := g.Frequency(nil)
		assert.NotNil(t, err)
	})

	t.Run("nil value tag", func(t *testing.T) {
		err := g.TagFrequency("", nil)
		assert.NotNil(t, err)
	})
}
