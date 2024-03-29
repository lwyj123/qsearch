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
)

type MatchAllQuery struct {
}

// NewMatchAllQuery creates a Query which will
// match all documents in the index.
func NewMatchAllQuery() *MatchAllQuery {
	return &MatchAllQuery{}
}

func (q *MatchAllQuery) MarshalJSON() ([]byte, error) {
	tmp := map[string]interface{}{
		"match_all": map[string]interface{}{},
	}
	return json.Marshal(tmp)
}
