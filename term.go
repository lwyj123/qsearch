//  Copyright (c) 2014 Couchbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package query

type TermQuery struct {
	Term     string `json:"term"`
	FieldVal string `json:"field,omitempty"`
}

// NewTermQuery creates a new Query for finding an
// exact term match in the index.
func NewTermQuery(term string) *TermQuery {
	return &TermQuery{
		Term: term,
	}
}

func (q *TermQuery) SetField(f string) {
	q.FieldVal = f
}

func (q *TermQuery) Field() string {
	return q.FieldVal
}
