package main

import (
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func main() {
	fmt.Println("uql driver")

	body := []byte(`{
  "query": {
    "bool": {
      "must": [
        { "term": { "field1": "value1" }},
        {
          "bool": {
            "should": [
              { "term": { "field2": "value2" }},
              { "term": { "field2": "value3" }}
            ]
          }
        }
      ]
    }
  }
}`)
	query := &types.Query{}
	_ = query
	err := json.Unmarshal(body, query)
	if err != nil {
		fmt.Println(err)
	}
	//TraverseQuery(query, &PrintVisitor{})
	rewriter := NewSExpressionRewriter()
	TraverseQuery(query, rewriter)
	fmt.Print(rewriter.output)
}
