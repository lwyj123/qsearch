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

type RegexpQuery struct {
	Regexp   string `json:"regexp"`
	FieldVal string `json:"field,omitempty"`
}

// NewRegexpQuery creates a new Query which finds
// documents containing terms that match the
// specified regular expression.  The regexp pattern
// SHOULD NOT include ^ or $ modifiers, the search
// will only match entire terms even without them.
func NewRegexpQuery(regexp string) *RegexpQuery {
	return &RegexpQuery{
		Regexp: regexp,
	}
}

func (q *RegexpQuery) SetField(f string) {
	q.FieldVal = f
}

func (q *RegexpQuery) Field() string {
	return q.FieldVal
}

func (q *RegexpQuery) Validate() error {
	return nil // real validation delayed until searcher constructor
}
