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

import (
	"encoding/json"
	"fmt"
)

type DisjunctionQuery struct {
	Disjuncts       []Query `json:"disjuncts"`
	Min             float64 `json:"min"`
	queryStringMode bool
}

// NewDisjunctionQuery creates a new compound Query.
// Result documents satisfy at least one Query.
func NewDisjunctionQuery(disjuncts []Query) *DisjunctionQuery {
	return &DisjunctionQuery{
		Disjuncts: disjuncts,
	}
}

func (q *DisjunctionQuery) AddQuery(aq ...Query) {
	q.Disjuncts = append(q.Disjuncts, aq...)
	//for _, aaq := range aq {
	//	q.Disjuncts = append(q.Disjuncts, aaq)
	//}
}

func (q *DisjunctionQuery) SetMin(m float64) {
	q.Min = m
}

func (q *DisjunctionQuery) Validate() error {
	if int(q.Min) > len(q.Disjuncts) {
		return fmt.Errorf("disjunction query has fewer than the minimum number of clauses to satisfy")
	}
	for _, q := range q.Disjuncts {
		if q, ok := q.(ValidatableQuery); ok {
			err := q.Validate()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (q *DisjunctionQuery) UnmarshalJSON(data []byte) error {
	tmp := struct {
		Disjuncts []json.RawMessage `json:"disjuncts"`
		Min       float64           `json:"min"`
	}{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	q.Disjuncts = make([]Query, len(tmp.Disjuncts))
	for i, term := range tmp.Disjuncts {
		query, err := ParseQuery(term)
		if err != nil {
			return err
		}
		q.Disjuncts[i] = query
	}
	q.Min = tmp.Min
	return nil
}
