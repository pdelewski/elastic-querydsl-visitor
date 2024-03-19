package main

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type QueryVisitor interface {
	PreVisitQuery(q *types.Query)
	PostVisitQuery(q *types.Query)
	PreVisitBool(b *types.BoolQuery)
	PostVisitBool(b *types.BoolQuery)
	PreVisitBoostingQuery(b *types.BoostingQuery)
	PostVisitBoostingQuery(b *types.BoostingQuery)
	PreVisitTypeQuery(t *types.TypeQuery)
	PostVisitTypeQuery(t *types.TypeQuery)
	PreVisitCommonTermsQuery(ctq *types.CommonTermsQuery, field string)
	PostVisitCommonTermsQuery(ctq *types.CommonTermsQuery, field string)
	PreVisitCombinedFieldsQuery(ctq *types.CombinedFieldsQuery)
	PostVisitCombinedFieldsQuery(ctq *types.CombinedFieldsQuery)
	PreVisitTerm(t *types.TermQuery, field string)
	PostVisitTerm(t *types.TermQuery, field string)
	PreVisitConstantScoreQuery(csq *types.ConstantScoreQuery)
	PostVisitConstantScoreQuery(csq *types.ConstantScoreQuery)
	PreVisitDisMaxQuery(dmq *types.DisMaxQuery)
	PostVisitDisMaxQuery(dmq *types.DisMaxQuery)
	PreVisitDistanceFeatureQuery(dfq types.DistanceFeatureQuery)
	PostVisitDistanceFeatureQuery(dfq types.DistanceFeatureQuery)
	PreVisitExistsQuery(eq *types.ExistsQuery)
	PostVisitExistsQuery(eq *types.ExistsQuery)
	PreVisitMatchQuery(mq *types.MatchQuery, field string)
	PostVisitMatchQuery(mq *types.MatchQuery, field string)
	PreVisitMatchAllQuery(mq *types.MatchAllQuery)
	PostVisitMatchAllQuery(mq *types.MatchAllQuery)
	PreVisitMatchBoolPrefixQuery(mq *types.MatchBoolPrefixQuery, field string)
	PostVisitMatchBoolPrefixQuery(mq *types.MatchBoolPrefixQuery, field string)
}
