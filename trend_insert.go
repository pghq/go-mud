package mud

import (
	"github.com/pghq/go-tea"
)

// Insert a value with tags
func (s *TrendService) Insert(key string, value interface{}, data []float64) error {
	if key == "" || len(data) == 0 {
		return tea.NewError("bad request")
	}

	item, err := s.store.Insert(key, value)
	if err != nil {
		return tea.Error(err)
	}

	s.graph.Plot(item.Id, data)
	return nil
}
