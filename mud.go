// Copyright 2021 PGHQ. All Rights Reserved.
//
// Licensed under the GNU General Public License, Version 3 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mud

import (
	"context"
	"reflect"
	"time"

	"github.com/pghq/go-ark"
	"github.com/pghq/go-tea"

	"github.com/pghq/go-mud/frequency"
	"github.com/pghq/go-mud/internal"
	"github.com/pghq/go-mud/neighbor"
)

const (
	// QueryCacheTTL is the TTL for queries
	QueryCacheTTL = 100 * time.Millisecond
)

// Graph classification service.
type Graph struct {
	neighbors   *neighbor.Tree
	frequencies *frequency.Map
	conn        *ark.KVSConn
}

// NewGraph creates a new graph.
func NewGraph() *Graph {
	dm := ark.Open()
	conn, _ := dm.ConnectKVS(context.Background(), "inmem")

	g := Graph{
		conn:        conn,
		neighbors:   neighbor.NewTree(),
		frequencies: frequency.NewMap(),
	}

	return &g
}

// Wait for plots be processed.
func (g *Graph) Wait() {
	g.neighbors.Wait()
}

// view keys by algorithm
func (g *Graph) view(key []byte, v interface{}, q internal.Query, fn func(q internal.Query) ([][]byte, error)) error {
	return g.conn.Do(context.Background(), func(tx *ark.KVSTxn) error {
		rv := reflect.ValueOf(v)
		if rv.Kind() != reflect.Ptr || rv.IsNil() || !rv.IsValid() {
			return tea.NewError("dst must be a pointer")
		}

		rv = rv.Elem()
		if rv.Kind() != reflect.Slice {
			return tea.NewError("dst must be a pointer to slice")
		}

		var keys [][]byte
		if _, err := tx.Get(key, &keys).Resolve(); err != nil {
			keys, err = fn(q)
			if err != nil {
				return tea.Error(err)
			}
			tx.InsertWithTTL(key, keys, QueryCacheTTL)
		}

		var values []reflect.Value
		for _, key := range keys {
			item := reflect.New(reflect.TypeOf(v).Elem().Elem())
			if _, err := tx.Get(key, &item).Resolve(); err != nil {
				return tea.Error(err)
			}
			values = append(values, item.Elem())
		}

		rv.Set(reflect.Append(rv, values...))
		return nil
	})
}
