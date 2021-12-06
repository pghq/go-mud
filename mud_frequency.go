package mud

import (
	"github.com/pghq/go-mud/internal"
)

// Frequency of occurrences ordered
func (g *Graph) Frequency(v interface{}, limit ...int) error {
	key := []byte("mud.graph.frequency")
	return g.view(key, v, internal.LimitQuery(limit), g.frequencies.Max)
}

// TagFrequency of occurrences ordered
func (g *Graph) TagFrequency(tag string, v interface{}, limit ...int) error {
	key := []byte("mud.graph.frequency." + tag)
	return g.view(key, v, internal.TagQuery(tag, limit), g.frequencies.MaxTag)
}
