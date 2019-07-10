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

type QueryStringQuery struct {
	Query string `json:"query"`
}

// NewQueryStringQuery creates a new Query used for
// finding documents that satisfy a query string.  The
// query string is a small query language for humans.
func NewQueryStringQuery(query string) *QueryStringQuery {
	return &QueryStringQuery{
		Query: query,
	}
}

func (q *QueryStringQuery) Parse() (Query, error) {
	return parseQuerySyntax(q.Query)
}

func (q *QueryStringQuery) Validate() error {
	newQuery, err := parseQuerySyntax(q.Query)
	if err != nil {
		return err
	}
	if newQuery, ok := newQuery.(ValidatableQuery); ok {
		return newQuery.Validate()
	}
	return nil
}
