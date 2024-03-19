package main

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type AggregationRewriter struct {
	output string
}

func NewAggregationRewriter() *AggregationRewriter {
	return &AggregationRewriter{}
}

func (r *AggregationRewriter) PreVisitAggregations(aggregation *types.Aggregations)  {}
func (r *AggregationRewriter) PostVisitAggregations(aggregation *types.Aggregations) {}
