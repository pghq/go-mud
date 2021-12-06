package mud

import (
	"context"

	"github.com/pghq/go-ark"
	"github.com/pghq/go-tea"
)

// Plot point
func (g *Graph) Plot(key []byte, value interface{}, point []float64, tags ...string) error {
	if key == nil || len(point) == 0 {
		return tea.NewError("bad request")
	}

	return g.conn.Do(context.Background(), func(tx *ark.KVSTxn) error {
		_, err := tx.Insert(key, value).Resolve()
		g.neighbors.Insert(key, point)
		g.frequencies.Incr(key, tags...)
		return err
	})
}
