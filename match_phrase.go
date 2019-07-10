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

type MatchPhraseQuery struct {
	MatchPhrase string `json:"match_phrase"`
	FieldVal    string `json:"field,omitempty"`
	Analyzer    string `json:"analyzer,omitempty"`
}

// NewMatchPhraseQuery creates a new Query object
// for matching phrases in the index.
// An Analyzer is chosen based on the field.
// Input text is analyzed using this analyzer.
// Token terms resulting from this analysis are
// used to build a search phrase.  Result documents
// must match this phrase. Queried field must have been indexed with
// IncludeTermVectors set to true.
func NewMatchPhraseQuery(matchPhrase string) *MatchPhraseQuery {
	return &MatchPhraseQuery{
		MatchPhrase: matchPhrase,
	}
}

func (q *MatchPhraseQuery) SetField(f string) {
	q.FieldVal = f
}

func (q *MatchPhraseQuery) Field() string {
	return q.FieldVal
}

//func tokenStreamToPhrase(tokens analysis.TokenStream) [][]string {
//    firstPosition := int(^uint(0) >> 1)
//    lastPosition := 0
//    for _, token := range tokens {
//        if token.Position < firstPosition {
//            firstPosition = token.Position
//        }
//        if token.Position > lastPosition {
//            lastPosition = token.Position
//        }
//    }
//    phraseLen := lastPosition - firstPosition + 1
//    if phraseLen > 0 {
//        rv := make([][]string, phraseLen)
//        for _, token := range tokens {
//            pos := token.Position - firstPosition
//            rv[pos] = append(rv[pos], string(token.Term))
//        }
//        return rv
//    }
//    return nil
//}
