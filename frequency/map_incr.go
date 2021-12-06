package frequency

import (
	"context"
)

// Incr key
func (m *Map) Incr(key []byte, tags ...string) *Map {
	m.pending <- node{Key: key, Tags: tags}
	return m
}

// incrJob increments pending nodes
func (m *Map) incrJob(ctx context.Context) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for {
		select {
		case <-ctx.Done():
			return
		case n := <-m.pending:
			data, present := m.nodes[string(n.Key)]
			if !present {
				data = &node{
					Key:       n.Key,
					Tags:      n.Tags,
					positions: make(map[string]int),
				}

				data.position = len(m.list)
				m.list = append(m.list, data)
				for _, tag := range n.Tags {
					data.positions[tag] = len(m.tags[tag])
					m.tags[tag] = append(m.tags[tag], data)
				}
			}

			data.Frequency += 1
			for i := data.position - 1; i >= 0; i-- {
				next := m.list[i]
				if data.Frequency > next.Frequency {
					m.list[data.position], m.list[next.position] = m.list[next.position], m.list[data.position]
					data.position, next.position = next.position, data.position
				}
			}

			for _, tag := range n.Tags {
				list := m.tags[tag]
				for i := data.positions[tag] - 1; i >= 0; i-- {
					next := list[i]
					if data.Frequency > next.Frequency {
						list[data.positions[tag]], list[next.positions[tag]] = list[next.positions[tag]], list[data.positions[tag]]
						data.positions[tag], next.positions[tag] = next.positions[tag], data.positions[tag]
					}
				}
			}

			m.nodes[string(n.Key)] = data
		default:
			for {
				select {
				case wg := <-m.groups:
					wg.Done()
				default:
					return
				}
			}
		}
	}
}
