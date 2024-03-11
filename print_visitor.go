package main

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"strconv"
)

type PrintVisitor struct {
}

func (p *PrintVisitor) PreVisitQuery(q *types.Query) {
	fmt.Println("Pre query")

}

func (p *PrintVisitor) PostVisitQuery(q *types.Query) {
	fmt.Println("Post query")
}

func (p *PrintVisitor) PreVisitBool(b *types.BoolQuery) {
	fmt.Println("Pre bool filter:" + strconv.Itoa(len(b.Filter)))
	fmt.Println("Pre bool must:" + strconv.Itoa(len(b.Must)))
	fmt.Println("Pre bool should:" + strconv.Itoa(len(b.Should)))
}

func (p *PrintVisitor) PostVisitBool(b *types.BoolQuery) {
	fmt.Println("Post bool")
}

func (p *PrintVisitor) PreVisitBoostingQuery(b *types.BoostingQuery) {
	fmt.Println("Pre boostring")
}

func (p *PrintVisitor) PostVisitBoostingQuery(b *types.BoostingQuery) {
	fmt.Println("Post boosting")
}

func (p *PrintVisitor) PreVisitTypeQuery(t *types.TypeQuery) {
	fmt.Println("Pre type type query")
}
func (p *PrintVisitor) PostVisitTypeQuery(t *types.TypeQuery) {
	fmt.Println("Post type type query")
}

func (p *PrintVisitor) PreVisitCommonTermsQuery(ctq *types.CommonTermsQuery) {
	fmt.Println("Pre common terms query")
}
func (p *PrintVisitor) PostVisitCommonTermsQuery(ctq *types.CommonTermsQuery) {
	fmt.Println("Post common terms query")
}

func (p *PrintVisitor) PreVisitCombinedFieldsQuery(ctq *types.CombinedFieldsQuery) {
	fmt.Println("Pre combined fields query")
}
func (p *PrintVisitor) PostVisitCombinedFieldsQuery(ctq *types.CombinedFieldsQuery) {
	fmt.Println("Post combined fields query")
}
func (p *PrintVisitor) PreVisitTerm(t *types.TermQuery, field string) {
	fmt.Println("Pre term query")
}
func (p *PrintVisitor) PostVisitTerm(t *types.TermQuery, field string) {
	fmt.Println("Post term query")
}
