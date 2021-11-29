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

// Package mud provides a client for ML recommendations.
package mud

import (
	"github.com/pghq/go-mud/graph"
	"github.com/pghq/go-mud/store"
)

// Client allows interaction with services within the domain.
type Client struct {
	common service

	Trends *TrendService
}

// NewClient creates a new client instance.
func NewClient() *Client {
	c := Client{
		common: service{
			store: store.New(),
			graph: graph.New(),
		},
	}

	c.Trends = (*TrendService)(&c.common)

	return &c
}

// service is a shared configuration for all services within the domain.
type service struct {
	graph *graph.Graph
	store *store.Store
}
