package main

import (
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"strconv"
)

type Bool struct {
	*types.BoolQuery
}

type Query struct {
	*types.Query
	bool *Bool
}

type Visitor interface {
	PreVisitQuery(q *types.Query)
	PostVisitQuery(q *types.Query)
	PreVisitBool(b *types.BoolQuery)
	PostVisitBool(b *types.BoolQuery)
}

type Traverser interface {
	TraverseQuery(q *Query)
	TraverseBool(b *Bool)
}

func (q *Query) Accept(v Traverser) {
	v.TraverseQuery(q)
}

func (b *Bool) Accept(v Traverser) {
	v.TraverseBool(b)
}

type BasicTraverser struct {
	visitor Visitor
}

func (p *BasicTraverser) TraverseBool(b *Bool) {
	p.visitor.PreVisitBool(b.BoolQuery)
	p.visitor.PostVisitBool(b.BoolQuery)
}

func (t *BasicTraverser) TraverseQuery(q *Query) {
	t.visitor.PreVisitQuery(q.Query)
	q.bool.Accept(t)
	t.visitor.PostVisitQuery(q.Query)
}

type PrintVisitor struct {
}

func (p *PrintVisitor) PreVisitQuery(q *types.Query) {
	fmt.Println("Pre query")

}

func (p *PrintVisitor) PostVisitQuery(q *types.Query) {
	fmt.Println("Post query")
}

func (p *PrintVisitor) PreVisitBool(b *types.BoolQuery) {
	fmt.Println("Pre bool:" + strconv.Itoa(len(b.Filter)))

}

func (p *PrintVisitor) PostVisitBool(b *types.BoolQuery) {
	fmt.Println("Post bool")
}

func main() {
	fmt.Println("uql driver")

	body := []byte(`{
    "query": {
        "bool": {
            "filter": [
                {
                    "range": {
                        "@timestamp": {
                            "lt": "now - 15m"
                        }
                    }
                }
            ]
        }
    },
    "runtime_mappings": {},
    "size": 100,
    "track_total_hits": true
}`)
	query := &types.Query{}
	_ = query
	err := json.Unmarshal(body, query)
	if err != nil {
		fmt.Println(err)
	}
	q := &Query{query, &Bool{BoolQuery: query.Bool}}
	v := &BasicTraverser{visitor: &PrintVisitor{}}
	q.Accept(v)
}
