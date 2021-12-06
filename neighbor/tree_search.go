package neighbor

import (
	"context"
	"math"

	"github.com/pghq/go-tea"
	"github.com/sjwhitworth/golearn/kdtree"
	"github.com/sjwhitworth/golearn/metrics/pairwise"

	"github.com/pghq/go-mud/internal"
)

// Search for nearest neighbors
func (t *Tree) Search(q internal.Query) ([][]byte, error) {
	if t.kd == nil {
		return nil, tea.NewNoContent("not found")
	}

	t.mutex.RLock()
	defer t.mutex.RUnlock()
	dst := make([]float64, t.size)
	copy(dst, q.Point)

	k := int(math.Sqrt(float64(len(t.data))))
	rows, _, _ := t.kd.Search(k, pairwise.NewEuclidean(), dst)

	end := int(math.Min(float64(q.Limit), float64(len(rows))))
	var keys [][]byte
	for i := 0; i < end; i++ {
		keys = append(keys, t.keys[rows[i]])
	}

	return keys, nil
}

// searchJob builds the kd tree if changes are present
func (t *Tree) searchJob(ctx context.Context) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	select {
	case <-ctx.Done():
		return
	default:
		if !t.changes {
			return
		}

		size := 0
		for _, data := range t.data {
			if size < len(data) {
				size = len(data)
			}
		}

		t.size = size
		dst := make([][]float64, len(t.data))
		for i, data := range t.data {
			dst[i] = make([]float64, size)
			copy(dst[i], data)
		}

		tree := kdtree.New()
		_ = tree.Build(dst)
		t.kd = tree
		t.changes = false
	}

	if len(t.pending) == 0 {
		for {
			select {
			case wg := <-t.groups:
				wg.Done()
			default:
				return
			}
		}
	}
}
