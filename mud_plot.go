package mud

import (
	"context"

	"github.com/pghq/go-ark"
	"github.com/pghq/go-tea"
)

// Plot point
func (g *Graph) Plot(key []byte, value interface{}, point []float64, tags ...string) error {
	if key == nil || len(point) == 0 {
		return tea.Err("bad request")
	}

	return g.db.Do(context.Background(), func(tx ark.Txn) error {
		if err := tx.Insert("", string(key), value); err != nil {
			return tea.Stack(err)
		}

		g.neighbors.Insert(key, point)
		g.frequencies.Incr(key, tags...)
		return nil
	})
}
