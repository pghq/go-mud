package graph

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLearn_PlotJob(t *testing.T) {
	t.Run("timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), 0)
		defer cancel()
		l := New()
		l.PlotJob(ctx)
	})
}

func TestLearn_SearchJob(t *testing.T) {
	t.Run("timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), 0)
		defer cancel()
		l := New()
		l.SearchJob(ctx)
	})
}

func TestLearn_Search(t *testing.T) {
	l := New()

	t.Run("no data", func(t *testing.T) {
		_, err := l.Search(nil)
		assert.NotNil(t, err)
	})

	l.Plot(0, []float64{2, 3}).
		Plot(1, []float64{5, 4}).
		Plot(2, []float64{4, 7}).
		Plot(3, []float64{8, 1}).
		Plot(4, []float64{7, 2}).
		Plot(5, []float64{9, 6})

	l.Wait()

	t.Run("out of bounds", func(t *testing.T) {
		_, err := l.Search([]float64{7, 3}, Page(1))
		assert.NotNil(t, err)
	})

	t.Run("success", func(t *testing.T) {
		rows, err := l.Search([]float64{7, 3}, Limit(2))
		assert.Nil(t, err)
		assert.NotNil(t, rows)
		assert.Len(t, rows, 2)
		assert.Equal(t, 4, rows[0])
		assert.Equal(t, 1, rows[1])
	})
}
