package main

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

func TraverseBooleanQuery(b *types.BoolQuery, v Visitor) {
	if b == nil {
		return
	}
	v.PreVisitBool(b)
	for _, q := range b.Filter {
		TraverseQuery(&q, v)
	}
	for _, q := range b.Should {
		TraverseQuery(&q, v)
	}
	for _, q := range b.Must {
		TraverseQuery(&q, v)
	}
	for _, q := range b.MustNot {
		TraverseQuery(&q, v)
	}
	v.PostVisitBool(b)
}

func TraverseCombinedFieldsQuery(cfq *types.CombinedFieldsQuery, v Visitor) {
	if cfq == nil {
		return
	}
	v.PreVisitCombinedFieldsQuery(cfq)
	v.PostVisitCombinedFieldsQuery(cfq)
}

func TraverseTypeQuery(t *types.TypeQuery, v Visitor) {
	if t == nil {
		return
	}
	v.PreVisitTypeQuery(t)
	v.PostVisitTypeQuery(t)
}

func TraverseBoostingQuery(b *types.BoostingQuery, v Visitor) {
	if b == nil {
		return
	}
	v.PreVisitBoostingQuery(b)
	TraverseQuery(b.Positive, v)
	TraverseQuery(b.Negative, v)
	v.PostVisitBoostingQuery(b)
}

func TraverseCommonTermsQuery(ctq *types.CommonTermsQuery, v Visitor) {
	if ctq == nil {
		return
	}
	v.PreVisitCommonTermsQuery(ctq)
	v.PostVisitCommonTermsQuery(ctq)
}

func TraverseTermQuery(t *types.TermQuery, field string, v Visitor) {
	if t == nil {
		return
	}
	v.PreVisitTerm(t, field)
	v.PostVisitTerm(t, field)
}

func TraverseQuery(q *types.Query, v Visitor) {
	if q == nil {
		return
	}
	v.PreVisitQuery(q)
	TraverseBooleanQuery(q.Bool, v)
	TraverseBoostingQuery(q.Boosting, v)
	TraverseTypeQuery(q.Type, v)
	for _, commonTermsQ := range q.Common {
		TraverseCommonTermsQuery(&commonTermsQ, v)
	}
	TraverseCombinedFieldsQuery(q.CombinedFields, v)
	for field, termQ := range q.Term {
		TraverseTermQuery(&termQ, field, v)
	}

	v.PostVisitQuery(q)
}
