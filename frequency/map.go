package frequency

import (
	"sync"

	"github.com/pghq/go-red"
)

// Map is an ordered frequency map
type Map struct {
	mutex   sync.RWMutex
	nodes   map[string]*node
	tags    map[string][]*node
	list    []*node
	worker  *red.Worker
	groups  chan *sync.WaitGroup
	pending chan node
}

// NewMap creates a new ordered map instance
func NewMap() *Map {
	m := Map{
		nodes:   make(map[string]*node),
		tags:    make(map[string][]*node),
		pending: make(chan node, 1000),
		groups:  make(chan *sync.WaitGroup),
	}

	m.worker = red.NewWorker(m.incrJob).Quiet()
	go m.worker.Start()

	return &m
}

// Wait for pending jobs to be processed
func (m *Map) Wait() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	m.groups <- &wg
	wg.Wait()
}

// node is the internal representation of frequency data
type node struct {
	Key       []byte
	Tags      []string
	Frequency int
	position  int
	positions map[string]int
}
