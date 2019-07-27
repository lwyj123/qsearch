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

package qsearch

import (
	"encoding/json"
	"fmt"
)

type PhraseQuery struct {
	Terms []string `json:"terms"`
	Field string   `json:"field,omitempty"`
}

// NewPhraseQuery creates a new Query for finding
// exact term phrases in the index.
// The provided terms must exist in the correct
// order, at the correct index offsets, in the
// specified field. Queried field must have been indexed with
// IncludeTermVectors set to true.
func NewPhraseQuery(terms []string, field string) *PhraseQuery {
	return &PhraseQuery{
		Terms: terms,
		Field: field,
	}
}

func (q *PhraseQuery) Validate() error {
	if len(q.Terms) < 1 {
		return fmt.Errorf("phrase query must contain at least one term")
	}
	return nil
}

func (q *PhraseQuery) UnmarshalJSON(data []byte) error {
	type _phraseQuery PhraseQuery
	tmp := _phraseQuery{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	q.Terms = tmp.Terms
	q.Field = tmp.Field
	return nil
}
