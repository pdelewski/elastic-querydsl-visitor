package main

import (
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
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

func printColoredTextWithBackground(text, textColor, backgroundColor string) {
	// ANSI escape code for clearing the line and setting background color
	fmt.Printf("\033[K%s%s%s%s%s\n", backgroundColor, textColor, text, Reset, Reset)
}

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
	indentation := 0
	_ = indentation
	jsonData := queryContent

	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonData), &data)
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
