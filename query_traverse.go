package main

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type QueryTraverser struct {
	Debug bool
}

func (qt *QueryTraverser) TraverseBooleanQuery(b *types.BoolQuery, v QueryVisitor, path []string) {
	if b == nil {
		return
	}
	v.PreVisitBool(b)
	for _, q := range b.Filter {
		if qt.Debug {
			path = append(path, getJsonTagName(&q, q.Bool))
		}
		qt.TraverseQuery(&q, v, path)
		if qt.Debug {
			path = path[:len(path)-1]
		}
	}
	for _, q := range b.Should {
		if qt.Debug {
			path = append(path, getJsonTagName(&q, q.Bool))
		}
		qt.TraverseQuery(&q, v, path)
		if qt.Debug {
			path = path[:len(path)-1]
		}
	}
	for _, q := range b.Must {
		if qt.Debug {
			path = append(path, getJsonTagName(&q, q.Bool))
		}
		qt.TraverseQuery(&q, v, path)
		if qt.Debug {
			path = path[:len(path)-1]
		}
	}
	for _, q := range b.MustNot {
		if qt.Debug {
			path = append(path, getJsonTagName(&q, q.Bool))
		}
		qt.TraverseQuery(&q, v, path)
		if qt.Debug {
			path = path[:len(path)-1]
		}
	}
	v.PostVisitBool(b)
}

func (qt *QueryTraverser) TraverseCombinedFieldsQuery(cfq *types.CombinedFieldsQuery, v QueryVisitor, path []string) {
	if cfq == nil {
		return
	}
	v.PreVisitCombinedFieldsQuery(cfq)
	v.PostVisitCombinedFieldsQuery(cfq)
}

func (qt *QueryTraverser) TraverseTypeQuery(t *types.TypeQuery, v QueryVisitor, path []string) {
	if t == nil {
		return
	}
	v.PreVisitTypeQuery(t)
	v.PostVisitTypeQuery(t)
}

func (qt *QueryTraverser) TraverseBoostingQuery(b *types.BoostingQuery, v QueryVisitor, path []string) {
	if b == nil {
		return
	}
	v.PreVisitBoostingQuery(b)
	qt.TraverseQuery(b.Positive, v, path)
	qt.TraverseQuery(b.Negative, v, path)
	v.PostVisitBoostingQuery(b)
}

func (qt *QueryTraverser) TraverseCommonTermsQuery(ctq *types.CommonTermsQuery, field string, v QueryVisitor, path []string) {
	if ctq == nil {
		return
	}
	v.PreVisitCommonTermsQuery(ctq, field)
	v.PostVisitCommonTermsQuery(ctq, field)
}

func (qt *QueryTraverser) TraverseTermQuery(t *types.TermQuery, field string, v QueryVisitor, path []string) {
	if t == nil {
		return
	}
	v.PreVisitTerm(t, field)
	v.PostVisitTerm(t, field)
}

func (qt *QueryTraverser) TraverseConstantScoreQuery(csq *types.ConstantScoreQuery, v QueryVisitor, path []string) {
	if csq == nil {
		return
	}
	v.PreVisitConstantScoreQuery(csq)
	qt.TraverseQuery(csq.Filter, v, path)
	v.PostVisitConstantScoreQuery(csq)
}

func (qt *QueryTraverser) TraverseDistanceFeatureQuery(dfq types.DistanceFeatureQuery, v QueryVisitor, path []string) {
	if dfq == nil {
		return
	}
	v.PreVisitDistanceFeatureQuery(dfq)
	v.PostVisitDistanceFeatureQuery(dfq)

}

func (qt *QueryTraverser) TraverseDismaxQuery(dq *types.DisMaxQuery, v QueryVisitor, path []string) {
	if dq == nil {
		return
	}
	v.PreVisitDisMaxQuery(dq)
	for _, q := range dq.Queries {
		qt.TraverseQuery(&q, v, path)
	}
	v.PostVisitDisMaxQuery(dq)
}

func (*QueryTraverser) TraverseExistsQuery(eq *types.ExistsQuery, v QueryVisitor, path []string) {
	if eq == nil {
		return
	}

	v.PreVisitExistsQuery(eq)
	v.PostVisitExistsQuery(eq)
}

func (*QueryTraverser) TraverseMatchQuery(mq *types.MatchQuery, field string, v QueryVisitor, path []string) {
	if mq == nil {
		return
	}
	v.PreVisitMatchQuery(mq, field)
	v.PostVisitMatchQuery(mq, field)
}

func (qt *QueryTraverser) TraverseMatchAllQuery(mq *types.MatchAllQuery, v QueryVisitor, path []string) {
	if mq == nil {
		return
	}
	v.PreVisitMatchAllQuery(mq)
	v.PostVisitMatchAllQuery(mq)
}

func (qt *QueryTraverser) TraverseMatchBoolPrefixQuery(mbpq *types.MatchBoolPrefixQuery, field string, v QueryVisitor, path []string) {
	if mbpq == nil {
		return
	}
	v.PreVisitMatchBoolPrefixQuery(mbpq, field)
	v.PostVisitMatchBoolPrefixQuery(mbpq, field)
}

func (qt *QueryTraverser) TraverseQuery(q *types.Query, v QueryVisitor, path []string) {
	if q == nil {
		return
	}
	if qt.Debug {
		path = append(path, "query")
	}
	v.PreVisitQuery(q)
	if qt.Debug {
		path = append(path, getJsonTagName(q, q.Bool))
	}
	qt.TraverseBooleanQuery(q.Bool, v, path)
	if qt.Debug {
		path = path[:len(path)-1]
	}
	for field, commonTermsQ := range q.Common {
		qt.TraverseCommonTermsQuery(&commonTermsQ, field, v, path)
	}
	qt.TraverseTypeQuery(q.Type, v, path)
	qt.TraverseBoostingQuery(q.Boosting, v, path)
	qt.TraverseDismaxQuery(q.DisMax, v, path)
	qt.TraverseCombinedFieldsQuery(q.CombinedFields, v, path)
	qt.TraverseConstantScoreQuery(q.ConstantScore, v, path)

	qt.TraverseDistanceFeatureQuery(q.DistanceFeature, v, path)

	qt.TraverseExistsQuery(q.Exists, v, path)

	for field, matchQ := range q.Match {
		qt.TraverseMatchQuery(&matchQ, field, v, path)
	}
	qt.TraverseMatchAllQuery(q.MatchAll, v, path)

	for field, matchBoolPrefixQ := range q.MatchBoolPrefix {
		qt.TraverseMatchBoolPrefixQuery(&matchBoolPrefixQ, field, v, path)
	}
	for field, termQ := range q.Term {
		qt.TraverseTermQuery(&termQ, field, v, path)
	}

	v.PostVisitQuery(q)
	if qt.Debug {
		path = path[:len(path)-1]
	}
}
