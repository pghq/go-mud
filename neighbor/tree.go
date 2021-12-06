package neighbor

import (
	"sync"

	"github.com/pghq/go-red"
	"github.com/sjwhitworth/golearn/kdtree"
)

// Tree containing kd tree of points
type Tree struct {
	mutex   sync.RWMutex
	ids     map[string]int
	keys    [][]byte
	data    [][]float64
	kd      *kdtree.Tree
	worker  *red.Worker
	changes bool
	size    int
	pending chan node
	groups  chan *sync.WaitGroup
}

// Wait for pending nodes to be processed
func (t *Tree) Wait() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	t.groups <- &wg
	wg.Wait()
}

// NewTree creates a new graph
func NewTree() *Tree {
	t := Tree{
		pending: make(chan node, 1000),
		groups:  make(chan *sync.WaitGroup),
		ids:     make(map[string]int),
	}
	t.worker = red.NewWorker(t.insertJob, t.searchJob).Quiet()
	go t.worker.Start()
	return &t
}

// node in the tree
type node struct {
	key  []byte
	data []float64
}
