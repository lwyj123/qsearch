# qsearch
query search utils

## Usage
```Golang
import "github.com/lwyj123/qsearch"

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

```
