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
	"github.com/pghq/go-ark"

	"github.com/pghq/go-mud/graph"
)

// Classifier is an instance of the KNN classifier.
type Classifier service

// NewClassifier creates a new client instance.
func NewClassifier() *Classifier {
	c := Classifier{
		store: ark.NewInMemory(),
		graph: graph.New(),
	}

	return &c
}

// Wait for graph to be ready
func (c *Classifier) Wait() {
	c.graph.Wait()
}

// service is a shared configuration for all services within the domain.
type service struct {
	graph *graph.Graph
	store *ark.InMemoryStore
}
