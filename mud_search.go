package mud

import (
	"github.com/pghq/go-ark"
	"github.com/pghq/go-tea"

	"github.com/pghq/go-mud/graph"
)

// Cursor for search results
type Cursor struct {
	row   int
	ids   []int
	store *ark.InMemoryStore
}

// Next advances the iterator
func (c *Cursor) Next() bool {
	c.row += 1
	return c.row < len(c.ids)
}

// Decode the current value
func (c *Cursor) Decode(v interface{}) error {
	id := c.ids[c.row]
	item, err := c.store.Get(id)
	if item != nil {
		err = item.Value(v)
	}
	return err
}

// NewCursor creates a new forecast instance
func NewCursor(ids []int, store *ark.InMemoryStore) *Cursor {
	return &Cursor{
		row:   -1,
		ids:   ids,
		store: store,
	}
}

// Search for the KNN to point
func (c *Classifier) Search(data []float64, opts ...graph.SearchOption) (*Cursor, error) {
	ids, err := c.graph.Search(data, opts...)
	if err != nil {
		return nil, tea.Error(err)
	}

	return NewCursor(ids, c.store), nil
}
