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

//var wildcardRegexpReplacer = strings.NewReplacer(
//	// characters in the wildcard that must
//	// be escaped in the regexp
//	"+", `\+`,
//	"(", `\(`,
//	")", `\)`,
//	"^", `\^`,
//	"$", `\$`,
//	".", `\.`,
//	"{", `\{`,
//	"}", `\}`,
//	"[", `\[`,
//	"]", `\]`,
//	`|`, `\|`,
//	`\`, `\\`,
//	// wildcard characters
//	"*", ".*",
//	"?", ".")

type WildcardQuery struct {
	Wildcard string `json:"wildcard"`
	FieldVal string `json:"field,omitempty"`
}

// NewWildcardQuery creates a new Query which finds
// documents containing terms that match the
// specified wildcard.  In the wildcard pattern '*'
// will match any sequence of 0 or more characters,
// and '?' will match any single character.
func NewWildcardQuery(wildcard string) *WildcardQuery {
	return &WildcardQuery{
		Wildcard: wildcard,
	}
}

func (q *WildcardQuery) SetField(f string) {
	q.FieldVal = f
}

func (q *WildcardQuery) Field() string {
	return q.FieldVal
}

func (q *WildcardQuery) Validate() error {
	return nil // real validation delayed until searcher constructor
}
