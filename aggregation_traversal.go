package main

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

func TraverseAggregations(aggregation *types.Aggregations, v AggregationVisitor) {
	if aggregation == nil {
		return
	}
	v.PreVisitAggregations(aggregation)
	v.PostVisitAggregations(aggregation)
}
