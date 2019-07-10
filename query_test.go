package query_test

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
