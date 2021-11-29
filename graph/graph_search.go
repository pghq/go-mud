package graph

import (
	"context"
	"math"

	"github.com/pghq/go-tea"
	"github.com/sjwhitworth/golearn/kdtree"
	"github.com/sjwhitworth/golearn/metrics/pairwise"
)

const (
	// DefaultLimit is the default max search results
	DefaultLimit int = 50
)

// SearchJob builds the kd tree if changes are present
func (g *Graph) SearchJob(ctx context.Context) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	select {
	case <-ctx.Done():
		return
	default:
		if !g.changes {
			return
		}

		size := 0
		for _, data := range g.data {
			if size < len(data) {
				size = len(data)
			}
		}

		g.size = size
		dst := make([][]float64, len(g.data))
		for i, data := range g.data {
			dst[i] = make([]float64, size)
			copy(dst[i], data)
		}

		tree := kdtree.New()
		_ = tree.Build(dst)
		g.tree = tree
		g.changes = false
	}

	for {
		select {
		case wg := <-g.wait:
			wg.Done()
		default:
			return
		}
	}
}

// Search for nearest neighbors
func (g *Graph) Search(data []float64, opts ...SearchOption) ([]int, error) {
	if g.tree == nil {
		return nil, tea.NewNoContent()
	}

	g.mutex.RLock()
	defer g.mutex.RUnlock()
	dst := make([]float64, g.size)
	copy(dst, data)

	conf := SearchConfig{
		limit: DefaultLimit,
	}

	for _, opt := range opts {
		opt.Apply(&conf)
	}

	k := int(math.Sqrt(float64(len(g.data))))
	rows, _, _ := g.tree.Search(k, pairwise.NewEuclidean(), dst)

	page := conf.limit * conf.page
	if page > len(rows) {
		return nil, tea.NewError("out of bounds")
	}

	limit := int(math.Min(float64(conf.limit), float64(len(rows))))
	return rows[page:limit], nil
}

// SearchConfig is a configuration for searching KNN
type SearchConfig struct {
	page  int
	limit int
}

// SearchOption is an option for the SearchConfig
type SearchOption interface {
	Apply(conf *SearchConfig)
}

type page int

func (p page) Apply(conf *SearchConfig) {
	if conf != nil {
		conf.page = int(p)
	}
}

// Page configures the search page
func Page(num int) SearchOption {
	return page(num)
}

type limit int

func (l limit) Apply(conf *SearchConfig) {
	if conf != nil {
		conf.limit = int(l)
	}
}

// Limit configures the search limit
func Limit(num int) SearchOption {
	return limit(num)
}
