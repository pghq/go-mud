package neighbor

import (
	"context"
)

// Insert node into the tree
func (t *Tree) Insert(key []byte, data []float64) *Tree {
	t.pending <- node{key: key, data: data}
	return t
}

// insertJob appends pending nodes
func (t *Tree) insertJob(ctx context.Context) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	for {
		select {
		case <-ctx.Done():
			return
		case point := <-t.pending:
			t.changes = true
			key := string(point.key)
			i, present := t.ids[key]
			if !present {
				i = len(t.data)
				t.data = append(t.data, []float64{})
				t.ids[key] = i
				t.keys = append(t.keys, point.key)
			}

			t.data[i] = point.data
		default:
			return
		}
	}
}
