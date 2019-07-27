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
	"errors"
	"fmt"
	"math"
	"time"
)

func parseDateTime(input string) (time.Time, error) {
	layout := "20060102"
	rv, err := time.Parse(layout, input)
	if err == nil {
		return rv, nil
	}
	return time.Time{}, errors.New("invalid date")
}

// QueryDateTimeFormat controls the format when Marshaling to JSON
var QueryDateTimeFormat = time.RFC3339

type QueryTime struct {
	time.Time
}

var MinRFC3339CompatibleTime time.Time
var MaxRFC3339CompatibleTime time.Time

func init() {
	MinRFC3339CompatibleTime, _ = time.Parse(time.RFC3339, "1677-12-01T00:00:00Z")
	MaxRFC3339CompatibleTime, _ = time.Parse(time.RFC3339, "2262-04-11T11:59:59Z")
}

func queryTimeFromString(t string) (time.Time, error) {
	rv, err := parseDateTime(t)
	if err != nil {
		return time.Time{}, err
	}
	return rv, nil
}

func (t *QueryTime) MarshalJSON() ([]byte, error) {
	tt := time.Time(t.Time)
	return []byte("\"" + tt.Format(QueryDateTimeFormat) + "\""), nil
}

func (t *QueryTime) UnmarshalJSON(data []byte) error {
	var timeString string
	err := json.Unmarshal(data, &timeString)
	if err != nil {
		return err
	}
	t.Time, err = parseDateTime(timeString)
	if err != nil {
		return err
	}
	return nil
}

type DateRangeQuery struct {
	Start          QueryTime `json:"start,omitempty"`
	End            QueryTime `json:"end,omitempty"`
	InclusiveStart *bool     `json:"inclusive_start,omitempty"`
	InclusiveEnd   *bool     `json:"inclusive_end,omitempty"`
	FieldVal       string    `json:"field,omitempty"`
}

// NewDateRangeQuery creates a new Query for ranges
// of date values.
// Date strings are parsed using the DateTimeParser configured in the
//  top-level config.QueryDateTimeParser
// Either, but not both endpoints can be nil.
func NewDateRangeQuery(start, end time.Time) *DateRangeQuery {
	return NewDateRangeInclusiveQuery(start, end, nil, nil)
}

// NewDateRangeInclusiveQuery creates a new Query for ranges
// of date values.
// Date strings are parsed using the DateTimeParser configured in the
//  top-level config.QueryDateTimeParser
// Either, but not both endpoints can be nil.
// startInclusive and endInclusive control inclusion of the endpoints.
func NewDateRangeInclusiveQuery(start, end time.Time, startInclusive, endInclusive *bool) *DateRangeQuery {
	return &DateRangeQuery{
		Start:          QueryTime{start},
		End:            QueryTime{end},
		InclusiveStart: startInclusive,
		InclusiveEnd:   endInclusive,
	}
}

func (q *DateRangeQuery) SetField(f string) {
	q.FieldVal = f
}

func (q *DateRangeQuery) Field() string {
	return q.FieldVal
}

func (q *DateRangeQuery) parseEndpoints() (*float64, *float64, error) {
	min := math.Inf(-1)
	max := math.Inf(1)
	if !q.Start.IsZero() {
		if !isDatetimeCompatible(q.Start) {
			// overflow
			return nil, nil, fmt.Errorf("invalid/unsupported date range, start: %v", q.Start)
		}
		startInt64 := q.Start.UnixNano()
		min = math.Float64frombits(uint64(startInt64))
	}
	if !q.End.IsZero() {
		if !isDatetimeCompatible(q.End) {
			// overflow
			return nil, nil, fmt.Errorf("invalid/unsupported date range, end: %v", q.End)
		}
		endInt64 := q.End.UnixNano()
		max = math.Float64frombits(uint64(endInt64))
	}

	return &min, &max, nil
}

func (q *DateRangeQuery) Validate() error {
	if q.Start.IsZero() && q.End.IsZero() {
		return fmt.Errorf("must specify start or end")
	}
	_, _, err := q.parseEndpoints()
	if err != nil {
		return err
	}
	return nil
}

func isDatetimeCompatible(t QueryTime) bool {
	if QueryDateTimeFormat == time.RFC3339 &&
		(t.Before(MinRFC3339CompatibleTime) || t.After(MaxRFC3339CompatibleTime)) {
		return false
	}

	return true
}
