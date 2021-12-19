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
	"github.com/pghq/go-ark/db"
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
	mapper      *ark.Mapper
}

// New creates a new graph.
func New(opts ...GraphOption) *Graph {
	g := Graph{
		mapper:      ark.New(),
		neighbors:   neighbor.NewTree(),
		frequencies: frequency.NewMap(),
	}

	for _, opt := range opts {
		opt(&g)
	}

	return &g
}

// Wait for plots be processed.
func (g *Graph) Wait() {
	g.neighbors.Wait()
}

// view keys by algorithm
func (g *Graph) view(key []byte, v interface{}, q internal.Query, fn func(q internal.Query) ([][]byte, error)) error {
	return g.mapper.Do(context.Background(), func(tx db.Txn) error {
		rv := reflect.ValueOf(v)
		if rv.Kind() != reflect.Ptr || rv.IsNil() || !rv.IsValid() {
			return tea.NewError("dst must be a pointer")
		}

		rv = rv.Elem()
		if rv.Kind() != reflect.Slice {
			return tea.NewError("dst must be a pointer to slice")
		}

		var keys [][]byte
		if err := tx.Get("", string(key), &keys); err != nil {
			keys, err = fn(q)
			if err != nil {
				return tea.Error(err)
			}
			_ = tx.Insert("", string(key), keys, db.TTL(QueryCacheTTL))
		}

		var values []reflect.Value
		for _, key := range keys {
			item := reflect.New(reflect.TypeOf(v).Elem().Elem())
			if err := tx.Get("", string(key), item.Interface()); err != nil {
				return tea.Error(err)
			}
			values = append(values, item.Elem())
		}

		rv.Set(reflect.Append(rv, values...))
		return nil
	}, db.BatchReadSize(2))
}

// GraphOption to configure custom graph
type GraphOption func(g *Graph)

// Mapper sets a custom data mapper
func Mapper(o *ark.Mapper) GraphOption {
	return func(g *Graph) {
		g.mapper = o
	}
}
