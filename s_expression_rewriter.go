package main

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type SExpressionRewriter struct {
	intendation int
	output      string
}

func NewSExpressionRewriter() *SExpressionRewriter {
	return &SExpressionRewriter{intendation: 0}
}

func (r *SExpressionRewriter) PreVisitQuery(q *types.Query) {

}
func (r *SExpressionRewriter) PostVisitQuery(q *types.Query) {

}
func (r *SExpressionRewriter) PreVisitBool(b *types.BoolQuery) {
	if len(b.Must) > 0 {
		for i := 0; i < r.intendation; i++ {
			r.output += "  "
		}
		r.output += "add\n"
		r.intendation++
	}
	if len(b.Should) > 0 {
		for i := 0; i < r.intendation; i++ {
			r.output += "  "
		}
		r.output += "or\n"
		r.intendation++
	}
}
func (r *SExpressionRewriter) PostVisitBool(b *types.BoolQuery) {

}
func (r *SExpressionRewriter) PreVisitBoostingQuery(b *types.BoostingQuery) {

}
func (r *SExpressionRewriter) PostVisitBoostingQuery(b *types.BoostingQuery) {

}
func (r *SExpressionRewriter) PreVisitTypeQuery(t *types.TypeQuery) {

}
func (r *SExpressionRewriter) PostVisitTypeQuery(t *types.TypeQuery) {

}
func (r *SExpressionRewriter) PreVisitCommonTermsQuery(ctq *types.CommonTermsQuery) {

}
func (r *SExpressionRewriter) PostVisitCommonTermsQuery(ctq *types.CommonTermsQuery) {

}
func (r *SExpressionRewriter) PreVisitCombinedFieldsQuery(ctq *types.CombinedFieldsQuery) {

}
func (r *SExpressionRewriter) PostVisitCombinedFieldsQuery(ctq *types.CombinedFieldsQuery) {

}

func (r *SExpressionRewriter) PreVisitTerm(t *types.TermQuery, field string) {
	for i := 0; i < r.intendation; i++ {
		r.output += "  "
	}

	r.output += field + " = " + t.Value.(string)
	r.output += "\n"
}

func (r *SExpressionRewriter) PostVisitTerm(t *types.TermQuery, field string) {

}
