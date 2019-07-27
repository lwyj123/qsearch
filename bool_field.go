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

type BoolFieldQuery struct {
	Bool     bool   `json:"bool"`
	FieldVal string `json:"field,omitempty"`
}

// NewBoolFieldQuery creates a new Query for boolean fields
func NewBoolFieldQuery(val bool) *BoolFieldQuery {
	return &BoolFieldQuery{
		Bool: val,
	}
}

func (q *BoolFieldQuery) SetField(f string) {
	q.FieldVal = f
}

func (q *BoolFieldQuery) Field() string {
	return q.FieldVal
}
