package graph

import (
	"context"
)

// PlotJob plots pending points
func (g *Graph) PlotJob(ctx context.Context) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	for {
		select {
		case <-ctx.Done():
			return
		case point := <-g.pending:
			g.changes = true
			if point.id >= len(g.data) {
				g.data = append(g.data, []float64{})
			}
			g.data[point.id] = point.data
		default:
			return
		}
	}
}

// Plot adds a point to be plotted later
func (g *Graph) Plot(id int, data []float64) *Graph {
	g.pending <- NewPoint(id, data)
	return g
}
