package graph

import (
	"sync"

	"github.com/pghq/go-red"
	"github.com/sjwhitworth/golearn/kdtree"
)

// Graph containing kd tree of points
type Graph struct {
	mutex   sync.RWMutex
	data    [][]float64
	tree    *kdtree.Tree
	worker  *red.Worker
	changes bool
	size    int
	pending chan *Point
	wait    chan *sync.WaitGroup
}

// Wait for pending messages to be processed
func (g *Graph) Wait() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	g.wait <- &wg
	wg.Wait()
}

// New creates a new graph
func New() *Graph {
	g := Graph{
		pending: make(chan *Point, 1000),
		wait:    make(chan *sync.WaitGroup),
	}
	g.worker = red.NewWorker(g.PlotJob, g.SearchJob).Quiet()
	go g.worker.Start()
	return &g
}

// Point on the graph
type Point struct {
	id   int
	data []float64
}

// NewPoint creates a new point
func NewPoint(id int, data []float64) *Point {
	return &Point{
		id:   id,
		data: data,
	}
}
