package neighbor

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pghq/go-mud/internal"
)

func TestTree_InsertJob(t *testing.T) {
	t.Parallel()

	t.Run("timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), 0)
		defer cancel()
		tr := NewTree()
		tr.insertJob(ctx)
	})
}

func TestTree_SearchJob(t *testing.T) {
	t.Parallel()

	t.Run("timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), 0)
		defer cancel()
		tr := NewTree()
		tr.searchJob(ctx)
	})
}

func TestTree_Search(t *testing.T) {
	t.Parallel()

	tr := NewTree()

	t.Run("no data", func(t *testing.T) {
		_, err := tr.Search(internal.Query{})
		assert.NotNil(t, err)
	})

	tr.Insert([]byte("0"), []float64{2, 3}).
		Insert([]byte("1"), []float64{5, 4}).
		Insert([]byte("2"), []float64{4, 7}).
		Insert([]byte("3"), []float64{8, 1}).
		Insert([]byte("4"), []float64{7, 2}).
		Insert([]byte("5"), []float64{9, 6})

	tr.Wait()

	t.Run("success", func(t *testing.T) {
		keys, err := tr.Search(internal.Query{Point: []float64{7, 3}, Limit: 2})
		assert.Nil(t, err)
		assert.NotNil(t, keys)
		assert.Len(t, keys, 2)
		assert.Equal(t, [][]byte{[]byte("4"), []byte("1")}, keys)
	})
}
