package main

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func TraverseBooleanQuery(b *types.BoolQuery, v QueryVisitor) {
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

func TraverseCombinedFieldsQuery(cfq *types.CombinedFieldsQuery, v QueryVisitor) {
	if cfq == nil {
		return
	}
	v.PreVisitCombinedFieldsQuery(cfq)
	v.PostVisitCombinedFieldsQuery(cfq)
}

func TraverseTypeQuery(t *types.TypeQuery, v QueryVisitor) {
	if t == nil {
		return
	}
	v.PreVisitTypeQuery(t)
	v.PostVisitTypeQuery(t)
}

func TraverseBoostingQuery(b *types.BoostingQuery, v QueryVisitor) {
	if b == nil {
		return
	}
	v.PreVisitBoostingQuery(b)
	TraverseQuery(b.Positive, v)
	TraverseQuery(b.Negative, v)
	v.PostVisitBoostingQuery(b)
}

func TraverseCommonTermsQuery(ctq *types.CommonTermsQuery, field string, v QueryVisitor) {
	if ctq == nil {
		return
	}
	v.PreVisitCommonTermsQuery(ctq, field)
	v.PostVisitCommonTermsQuery(ctq, field)
}

func TraverseTermQuery(t *types.TermQuery, field string, v QueryVisitor) {
	if t == nil {
		return
	}
	v.PreVisitTerm(t, field)
	v.PostVisitTerm(t, field)
}

func TraverseConstantScoreQuery(csq *types.ConstantScoreQuery, v QueryVisitor) {
	if csq == nil {
		return
	}
	v.PreVisitConstantScoreQuery(csq)
	TraverseQuery(csq.Filter, v)
	v.PostVisitConstantScoreQuery(csq)
}

func TraverseDistanceFeatureQuery(dfq types.DistanceFeatureQuery, v QueryVisitor) {
	if dfq == nil {
		return
	}
	v.PreVisitDistanceFeatureQuery(dfq)
	v.PostVisitDistanceFeatureQuery(dfq)

}

func TraverseDismaxQuery(dq *types.DisMaxQuery, v QueryVisitor) {
	if dq == nil {
		return
	}
	v.PreVisitDisMaxQuery(dq)
	for _, q := range dq.Queries {
		TraverseQuery(&q, v)
	}
	v.PostVisitDisMaxQuery(dq)
}

func TraverseExistsQuery(eq *types.ExistsQuery, v QueryVisitor) {
	if eq == nil {
		return
	}

	v.PreVisitExistsQuery(eq)
	v.PostVisitExistsQuery(eq)
}

func TraverseMatchQuery(mq *types.MatchQuery, field string, v QueryVisitor) {
	if mq == nil {
		return
	}
	v.PreVisitMatchQuery(mq, field)
	v.PostVisitMatchQuery(mq, field)
}

func TraverseMatchAllQuery(mq *types.MatchAllQuery, v QueryVisitor) {
	if mq == nil {
		return
	}
	v.PreVisitMatchAllQuery(mq)
	v.PostVisitMatchAllQuery(mq)
}

func TraverseMatchBoolPrefixQuery(mbpq *types.MatchBoolPrefixQuery, field string, v QueryVisitor) {
	if mbpq == nil {
		return
	}
	v.PreVisitMatchBoolPrefixQuery(mbpq, field)
	v.PostVisitMatchBoolPrefixQuery(mbpq, field)
}

func TraverseQuery(q *types.Query, v QueryVisitor) {
	if q == nil {
		return
	}
	v.PreVisitQuery(q)

	TraverseBooleanQuery(q.Bool, v)
	for field, commonTermsQ := range q.Common {
		TraverseCommonTermsQuery(&commonTermsQ, field, v)
	}
	TraverseTypeQuery(q.Type, v)
	TraverseBoostingQuery(q.Boosting, v)
	TraverseDismaxQuery(q.DisMax, v)
	TraverseCombinedFieldsQuery(q.CombinedFields, v)
	TraverseConstantScoreQuery(q.ConstantScore, v)

	TraverseDistanceFeatureQuery(q.DistanceFeature, v)

	TraverseExistsQuery(q.Exists, v)

	for field, matchQ := range q.Match {
		TraverseMatchQuery(&matchQ, field, v)
	}
	TraverseMatchAllQuery(q.MatchAll, v)

	for field, matchBoolPrefixQ := range q.MatchBoolPrefix {
		TraverseMatchBoolPrefixQuery(&matchBoolPrefixQ, field, v)
	}
	for field, termQ := range q.Term {
		TraverseTermQuery(&termQ, field, v)
	}

	v.PostVisitQuery(q)
}
