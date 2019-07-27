package qsearch_test

import (
	"testing"

	"github.com/lwyj123/qsearch"
)

func TestParseQuery(t *testing.T) {
	queryString := query.NewQueryStringQuery("key word ping:echo date:>=20190618 date:<20190625")
	q, err := queryString.Parse()
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("q: %#v", q)
		t.Logf("q.Must: %#v", q.(*query.BooleanQuery).Must)
		t.Logf("q.Should.Disjuncts: %#v", q.(*query.BooleanQuery).Should.(*query.DisjunctionQuery).Disjuncts)
		t.Logf("q.MustNot: %#v", q.(*query.BooleanQuery).MustNot)

		t.Logf("%#v", q.(*query.BooleanQuery).Should.(*query.DisjunctionQuery).Disjuncts[0])
		t.Logf("%#v", q.(*query.BooleanQuery).Should.(*query.DisjunctionQuery).Disjuncts[1])
		t.Logf("%#v", q.(*query.BooleanQuery).Should.(*query.DisjunctionQuery).Disjuncts[2])
		t.Logf("%#v", q.(*query.BooleanQuery).Should.(*query.DisjunctionQuery).Disjuncts[3])
		t.Logf("%#v", q.(*query.BooleanQuery).Should.(*query.DisjunctionQuery).Disjuncts[4])
	}
}

func TestParseQuery_2(t *testing.T) {
	queryString := query.NewQueryStringQuery("this AND that OR thus")
	q, err := queryString.Parse()
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("q: %#v", q)
		t.Logf("q.Must: %#v", q.(*query.BooleanQuery).Must)
		t.Logf("q.Should: %#v", q.(*query.BooleanQuery).Should.(*query.DisjunctionQuery))
		t.Logf("q.MustNot: %#v", q.(*query.BooleanQuery).MustNot)
		t.Logf("%#v", q.(*query.BooleanQuery).Should.(*query.DisjunctionQuery).Disjuncts[0])
	}
}

/*


type listAuthorQuery struct {
	StartDate   int
	EndDate     int
	MediaSearch *string
	Owners      *[]string
	Category    *string
}

func parseQuery(qs string) (*listAuthorQuery, error) {
	var result listAuthorQuery
	queryString := query.NewQueryStringQuery(qs)
	q, err := queryString.Parse()
	if err != nil {
		return nil, err
	}
	var owners = make([]string, 0)
	for _, queryItem := range q.(*query.BooleanQuery).Should.(*query.DisjunctionQuery).Disjuncts {
		switch queryType := queryItem.(type) {
		case *query.MatchQuery:
			if queryType.Field() == "owner" {
				owners = append(owners, queryType.Match)
			} else if queryType.Field() == "" {
				result.MediaSearch = &queryType.Match
			} else if queryType.Field() == "category" {
				result.Category = &queryType.Match
			}
		case *query.NumericRangeQuery:
			if queryType.FieldVal == "date" {
				if queryType.Min != nil {
					if *queryType.InclusiveMin {
						// date:>=20060102
						result.StartDate = int(*queryType.Min)
					} else {
						// date:>20060102
						dateInclusive, err := date.DateAddOne(int(*queryType.Min))
						if err != nil {
							return nil, err
						}
						result.StartDate = dateInclusive
					}
				}
				if queryType.Max != nil {
					if *queryType.InclusiveMax {
						// date:<=20060102
						result.EndDate = int(*queryType.Max)
					} else {
						// date:<20060102
						dateInclusive, err := date.DateMinusOne(int(*queryType.Max))
						if err != nil {
							return nil, err
						}
						result.EndDate = dateInclusive
					}
				}
			}
		default:
			logs.Debugf("not match %#v", queryType)
		}
	}
	if len(owners) != 0 {
		result.Owners = &owners
	}
	return &result, nil
}


*/
