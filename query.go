package query

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Query interface {
	//Field() string
}

// A FieldableQuery represents a Query which can be restricted
// to a single field.
type FieldableQuery interface {
	Query
	SetField(f string)
	Field() string
}

// A ValidatableQuery represents a Query which can be validated
// prior to execution.
type ValidatableQuery interface {
	Query
	Validate() error
}

func parseQuerySyntax(query string) (rq Query, err error) {
	if query == "" {
		return NewMatchNoneQuery(), nil
	}
	lex := newLexerWrapper(newQueryStringLex(strings.NewReader(query)))
	doParse(lex)

	if len(lex.errs) > 0 {
		return nil, fmt.Errorf(strings.Join(lex.errs, "\n"))
	}
	return lex.query, nil
}

func doParse(lex *lexerWrapper) {
	defer func() {
		r := recover()
		if r != nil {
			lex.errs = append(lex.errs, fmt.Sprintf("parse error: %v", r))
		}
	}()

	yyParse(lex)
}

const (
	queryShould = iota
	queryMust
	queryMustNot
)

type lexerWrapper struct {
	lex   yyLexer
	errs  []string
	query *BooleanQuery
}

func newLexerWrapper(lex yyLexer) *lexerWrapper {
	return &lexerWrapper{
		lex:   lex,
		query: NewBooleanQueryForQueryString(nil, nil, nil),
	}
}

func (l *lexerWrapper) Lex(lval *yySymType) int {
	return l.lex.Lex(lval)
}

func (l *lexerWrapper) Error(s string) {
	l.errs = append(l.errs, s)
}

// ParseQuery deserializes a JSON representation of
// a Query object.
func ParseQuery(input []byte) (Query, error) {
	var tmp map[string]interface{}
	err := json.Unmarshal(input, &tmp)
	if err != nil {
		return nil, err
	}
	_, isMatchQuery := tmp["match"]
	_, hasFuzziness := tmp["fuzziness"]
	if hasFuzziness && !isMatchQuery {
		var rv FuzzyQuery
		err := json.Unmarshal(input, &rv)
		if err != nil {
			return nil, err
		}
		return &rv, nil
	}
	_, isTermQuery := tmp["term"]
	if isTermQuery {
		var rv TermQuery
		err := json.Unmarshal(input, &rv)
		if err != nil {
			return nil, err
		}
		return &rv, nil
	}
	if isMatchQuery {
		var rv MatchQuery
		err := json.Unmarshal(input, &rv)
		if err != nil {
			return nil, err
		}
		return &rv, nil
	}
	_, isMatchPhraseQuery := tmp["match_phrase"]
	if isMatchPhraseQuery {
		var rv MatchPhraseQuery
		err := json.Unmarshal(input, &rv)
		if err != nil {
			return nil, err
		}
		return &rv, nil
	}
	_, hasMust := tmp["must"]
	_, hasShould := tmp["should"]
	_, hasMustNot := tmp["must_not"]
	if hasMust || hasShould || hasMustNot {
		var rv BooleanQuery
		err := json.Unmarshal(input, &rv)
		if err != nil {
			return nil, err
		}
		return &rv, nil
	}
	_, hasTerms := tmp["terms"]
	if hasTerms {
		var rv PhraseQuery
		err := json.Unmarshal(input, &rv)
		if err != nil {
			// now try multi-phrase
			var rv2 MultiPhraseQuery
			err = json.Unmarshal(input, &rv2)
			if err != nil {
				return nil, err
			}
			return &rv2, nil
		}
		return &rv, nil
	}
	_, hasConjuncts := tmp["conjuncts"]
	if hasConjuncts {
		var rv ConjunctionQuery
		err := json.Unmarshal(input, &rv)
		if err != nil {
			return nil, err
		}
		return &rv, nil
	}
	_, hasDisjuncts := tmp["disjuncts"]
	if hasDisjuncts {
		var rv DisjunctionQuery
		err := json.Unmarshal(input, &rv)
		if err != nil {
			return nil, err
		}
		return &rv, nil
	}

	_, hasSyntaxQuery := tmp["query"]
	if hasSyntaxQuery {
		var rv QueryStringQuery
		err := json.Unmarshal(input, &rv)
		if err != nil {
			return nil, err
		}
		return &rv, nil
	}
	_, hasMin := tmp["min"].(float64)
	_, hasMax := tmp["max"].(float64)
	if hasMin || hasMax {
		var rv NumericRangeQuery
		err := json.Unmarshal(input, &rv)
		if err != nil {
			return nil, err
		}
		return &rv, nil
	}
	_, hasMinStr := tmp["min"].(string)
	_, hasMaxStr := tmp["max"].(string)
	if hasMinStr || hasMaxStr {
		var rv TermRangeQuery
		err := json.Unmarshal(input, &rv)
		if err != nil {
			return nil, err
		}
		return &rv, nil
	}
	_, hasStart := tmp["start"]
	_, hasEnd := tmp["end"]
	if hasStart || hasEnd {
		var rv DateRangeQuery
		err := json.Unmarshal(input, &rv)
		if err != nil {
			return nil, err
		}
		return &rv, nil
	}
	_, hasPrefix := tmp["prefix"]
	if hasPrefix {
		var rv PrefixQuery
		err := json.Unmarshal(input, &rv)
		if err != nil {
			return nil, err
		}
		return &rv, nil
	}
	_, hasRegexp := tmp["regexp"]
	if hasRegexp {
		var rv RegexpQuery
		err := json.Unmarshal(input, &rv)
		if err != nil {
			return nil, err
		}
		return &rv, nil
	}
	_, hasWildcard := tmp["wildcard"]
	if hasWildcard {
		var rv WildcardQuery
		err := json.Unmarshal(input, &rv)
		if err != nil {
			return nil, err
		}
		return &rv, nil
	}
	_, hasMatchAll := tmp["match_all"]
	if hasMatchAll {
		var rv MatchAllQuery
		err := json.Unmarshal(input, &rv)
		if err != nil {
			return nil, err
		}
		return &rv, nil
	}
	_, hasMatchNone := tmp["match_none"]
	if hasMatchNone {
		var rv MatchNoneQuery
		err := json.Unmarshal(input, &rv)
		if err != nil {
			return nil, err
		}
		return &rv, nil
	}
	_, hasBool := tmp["bool"]
	if hasBool {
		var rv BoolFieldQuery
		err := json.Unmarshal(input, &rv)
		if err != nil {
			return nil, err
		}
		return &rv, nil
	}
	return nil, fmt.Errorf("unknown query type")
}
