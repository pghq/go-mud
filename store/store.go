package store

import (
	"bytes"
	"encoding/gob"
	"sync"

	"github.com/dgraph-io/badger/v3"
	"github.com/pghq/go-tea"
)

// Store implements persistence
type Store struct {
	db    *badger.DB
	mutex sync.RWMutex
	ids   map[string]int
}

// New creates a new store instance
func New() *Store {
	db, _ := badger.Open(badger.DefaultOptions("").WithLogger(nil).WithInMemory(true))
	return &Store{
		db:  db,
		ids: make(map[string]int),
	}
}

// Item is an instance of tagged arbitrary data
type Item struct {
	Id   int
	Data []byte
}

// Bytes encoded item
func (i *Item) Bytes() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	_ = enc.Encode(&i)
	return buf.Bytes()
}

// SetValue for item
func (i *Item) SetValue(v interface{}) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(v); err != nil {
		return tea.Error(err)
	}
	i.Data = buf.Bytes()
	return nil
}

// Value decodes the item
func (i *Item) Value(v interface{}) error {
	dec := gob.NewDecoder(bytes.NewReader(i.Data))
	return dec.Decode(v)
}
