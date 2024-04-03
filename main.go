package main

import (
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const queryContent = `{
  "query": {
    "bool": {
      "must": [
        { "term": { "field1": "value1" }},
        {
          "bool": {
            "should": [
              { "term": { "field2": "value2" }},
              { "term": { "field3": "value3" }}
            ]
          }
        }
      ]
    }
  }
}`

func traverse() {
	body := []byte(queryContent)
	query := &types.Query{}
	_ = query
	err := json.Unmarshal(body, query)
	if err != nil {
		fmt.Println(err)
	}

	rewriter := NewSExpressionRewriter()
	path := []string{}
	queryTraverser := QueryTraverser{Debug: true, PathMatched: false}
	queryTraverser.TraverseQuery(query, rewriter, path)

	var data map[string]interface{}
	err = json.Unmarshal([]byte(queryContent), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	rewriter.exprStack[0].DumpPrefix(data)
}

func main() {
	fmt.Println("uql driver")
	traverse()
}
