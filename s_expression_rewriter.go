package main

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type Expression struct {
	value    string
	children []*Expression
	eType    string
	path     []string
}

func (s *Expression) Dump(indentation int, data map[string]interface{}) {
	for i := 0; i < indentation; i++ {
		fmt.Print("  ")
	}
	fmt.Print(s.value)
	fmt.Println(" ", s.path)
	s.path = append(s.path, "*")
	printJSON(data, "", s.path)
	for _, child := range s.children {
		child.Dump(indentation+1, data)
	}
}

type SExpressionRewriter struct {
	exprStack []*Expression
}

func NewSExpressionRewriter() *SExpressionRewriter {
	return &SExpressionRewriter{}
}

func (r *SExpressionRewriter) PreVisitBool(b *types.BoolQuery, path []string) {
	boolExpr := &Expression{}
	if len(b.Must) > 0 {
		boolExpr.value += "and"
	}
	if len(b.Should) > 0 {
		boolExpr.value += "or"
	}
	boolExpr.eType = "bool"
	r.exprStack = append(r.exprStack, boolExpr)
	boolExpr.path = path
}
func (r *SExpressionRewriter) PreVisitTerm(t *types.TermQuery, field string, path []string) {
	termExpr := &Expression{}
	termExpr.value = field + " = " + t.Value.(string)
	termExpr.eType = "term"
	r.exprStack = append(r.exprStack, termExpr)
	termExpr.path = path
}
func (r *SExpressionRewriter) PostVisitTerm(*types.TermQuery, string, []string) {}
func (r *SExpressionRewriter) PreVisitQuery(*types.Query, []string)             {}
func (r *SExpressionRewriter) PostVisitQuery(*types.Query, []string)            {}
func (r *SExpressionRewriter) PostVisitBool(*types.BoolQuery, []string) {
	children := []*Expression{}
	i := len(r.exprStack) - 1
	for ; i >= 0; i-- {
		expr := r.exprStack[i]
		if expr.eType == "bool" && i != len(r.exprStack)-1 {
			break
		}
		children = append(children, expr)
	}
	r.exprStack[i].children = children
	r.exprStack = r.exprStack[:i+1]
}
func (r *SExpressionRewriter) PreVisitBoostingQuery(*types.BoostingQuery, []string)                {}
func (r *SExpressionRewriter) PostVisitBoostingQuery(*types.BoostingQuery, []string)               {}
func (r *SExpressionRewriter) PreVisitTypeQuery(*types.TypeQuery, []string)                        {}
func (r *SExpressionRewriter) PostVisitTypeQuery(*types.TypeQuery, []string)                       {}
func (r *SExpressionRewriter) PreVisitCommonTermsQuery(*types.CommonTermsQuery, string, []string)  {}
func (r *SExpressionRewriter) PostVisitCommonTermsQuery(*types.CommonTermsQuery, string, []string) {}
func (r *SExpressionRewriter) PreVisitCombinedFieldsQuery(*types.CombinedFieldsQuery, []string)    {}
func (r *SExpressionRewriter) PostVisitCombinedFieldsQuery(*types.CombinedFieldsQuery, []string)   {}

func (r *SExpressionRewriter) PreVisitConstantScoreQuery(*types.ConstantScoreQuery, []string)     {}
func (r *SExpressionRewriter) PostVisitConstantScoreQuery(*types.ConstantScoreQuery, []string)    {}
func (r *SExpressionRewriter) PreVisitDisMaxQuery(*types.DisMaxQuery, []string)                   {}
func (r *SExpressionRewriter) PostVisitDisMaxQuery(*types.DisMaxQuery, []string)                  {}
func (r *SExpressionRewriter) PreVisitDistanceFeatureQuery(types.DistanceFeatureQuery, []string)  {}
func (r *SExpressionRewriter) PostVisitDistanceFeatureQuery(types.DistanceFeatureQuery, []string) {}
func (r *SExpressionRewriter) PreVisitExistsQuery(*types.ExistsQuery, []string)                   {}
func (r *SExpressionRewriter) PostVisitExistsQuery(*types.ExistsQuery, []string)                  {}
func (r *SExpressionRewriter) PreVisitMatchQuery(*types.MatchQuery, string, []string)             {}
func (r *SExpressionRewriter) PostVisitMatchQuery(*types.MatchQuery, string, []string)            {}
func (r *SExpressionRewriter) PreVisitMatchAllQuery(*types.MatchAllQuery, []string)               {}
func (r *SExpressionRewriter) PostVisitMatchAllQuery(*types.MatchAllQuery, []string)              {}
func (r *SExpressionRewriter) PreVisitMatchBoolPrefixQuery(*types.MatchBoolPrefixQuery, string, []string) {
}
func (r *SExpressionRewriter) PostVisitMatchBoolPrefixQuery(*types.MatchBoolPrefixQuery, string, []string) {
}
