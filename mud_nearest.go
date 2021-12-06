package mud

import (
	"bytes"
	"encoding/binary"

	"github.com/pghq/go-mud/internal"
)

// Nearest to a point
func (g *Graph) Nearest(point []float64, v interface{}, limit ...int) error {
	var buf bytes.Buffer
	_ = binary.Write(&buf, binary.LittleEndian, point)
	key := append([]byte("mud.graph.nearest."), buf.Bytes()...)
	return g.view(key, v, internal.PointQuery(point, limit), g.neighbors.Search)
}
