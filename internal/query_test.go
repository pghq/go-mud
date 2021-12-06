package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointQuery(t *testing.T) {
	t.Run("not nil", func(t *testing.T) {
		q := PointQuery([]float64{1}, []int{1})
		assert.Equal(t, []float64{1}, q.Point)
		assert.Equal(t, 1, q.Limit)
	})
}

func TestTagQuery(t *testing.T) {
	t.Run("not nil", func(t *testing.T) {
		q := TagQuery("tag", []int{1})
		assert.Equal(t, "tag", q.Tag)
		assert.Equal(t, 1, q.Limit)
	})
}

func TestLimitQuery(t *testing.T) {
	t.Run("not nil", func(t *testing.T) {
		q := LimitQuery([]int{1})
		assert.Equal(t, 1, q.Limit)
	})
}
