package main

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type AggregationVisitor interface {
	PreVisitAggregations(aggregation *types.Aggregations)
	PostVisitAggregations(aggregation *types.Aggregations)
}
