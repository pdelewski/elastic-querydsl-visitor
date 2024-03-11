package main

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type Visitor interface {
	PreVisitQuery(q *types.Query)
	PostVisitQuery(q *types.Query)
	PreVisitBool(b *types.BoolQuery)
	PostVisitBool(b *types.BoolQuery)
	PreVisitBoostingQuery(b *types.BoostingQuery)
	PostVisitBoostingQuery(b *types.BoostingQuery)
	PreVisitTypeQuery(t *types.TypeQuery)
	PostVisitTypeQuery(t *types.TypeQuery)
	PreVisitCommonTermsQuery(ctq *types.CommonTermsQuery)
	PostVisitCommonTermsQuery(ctq *types.CommonTermsQuery)
	PreVisitCombinedFieldsQuery(ctq *types.CombinedFieldsQuery)
	PostVisitCombinedFieldsQuery(ctq *types.CombinedFieldsQuery)
	PreVisitTerm(t *types.TermQuery, field string)
	PostVisitTerm(t *types.TermQuery, field string)
}
