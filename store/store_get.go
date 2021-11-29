package store

import (
	"bytes"
	"encoding/gob"
	"strconv"

	"github.com/dgraph-io/badger/v3"
	"github.com/pghq/go-tea"
)

// Get an item
func (s *Store) Get(id int) (*Item, error) {
	var item *Item
	return item, s.db.View(func(txn *badger.Txn) error {
		i, err := txn.Get([]byte(strconv.Itoa(id)))
		if i != nil {
			err = i.Value(func(b []byte) error {
				dec := gob.NewDecoder(bytes.NewReader(b))
				return dec.Decode(&item)
			})
		}
		if err == badger.ErrKeyNotFound {
			err = tea.Error(err)
		}
		return err
	})
}
