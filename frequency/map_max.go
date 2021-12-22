package frequency

import (
	"math"

	"github.com/pghq/go-tea"

	"github.com/pghq/go-mud/internal"
)

// Max frequency
func (m *Map) Max(q internal.Query) ([][]byte, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if len(m.list) == 0 {
		return nil, tea.ErrNoContent("not found")
	}

	end := int(math.Min(float64(q.Limit), float64(len(m.list))))
	var keys [][]byte
	for i := 0; i < end; i++ {
		node := m.list[i]
		keys = append(keys, node.Key)
	}

	return keys, nil
}

// MaxTag frequency
func (m *Map) MaxTag(q internal.Query) ([][]byte, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	list, _ := m.tags[q.Tag]
	if len(list) == 0 {
		return nil, tea.ErrNoContent("not found")
	}

	end := int(math.Min(float64(q.Limit), float64(len(list))))
	var keys [][]byte
	for i := 0; i < end; i++ {
		node := list[i]
		keys = append(keys, node.Key)
	}

	return keys, nil
}
