package store

import (
	"strconv"

	"github.com/dgraph-io/badger/v3"
	"github.com/pghq/go-tea"
)

// Insert a value
func (s *Store) Insert(key string, v interface{}) (*Item, error) {
	var item *Item
	return item, s.db.Update(func(txn *badger.Txn) error {
		s.mutex.RLock()
		id, present := s.ids[key]
		if present {
			item, _ = s.Get(id)
		} else {
			item = &Item{
				Id: len(s.ids),
			}
		}
		s.mutex.RUnlock()
		if err := item.SetValue(v); err != nil {
			return tea.Error(err)
		}

		err := txn.Set([]byte(strconv.Itoa(item.Id)), item.Bytes())
		if err == nil {
			s.mutex.Lock()
			defer s.mutex.Unlock()
			s.ids[key] = item.Id
		}

		return err
	})
}
