package main

import (
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"strconv"
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

func printJSONHelper(data map[string]interface{}, indent string, currentPath []string, matchPath []string, matcher func([]string, []string) bool) {
	for key, value := range data {
		switch v := value.(type) {
		case map[string]interface{}:
			currentPath = append(currentPath, key)
			if matcher(currentPath, matchPath) {
				fmt.Println(Red+indent, key, ": {"+Reset)
			} else {
				fmt.Println(indent, key, ": {")
			}
			printJSONHelper(v, indent+"    ", currentPath, matchPath, matcher)
			currentPath = currentPath[:len(currentPath)-1]
			fmt.Println(indent, " }")
		case []interface{}:
			currentPath = append(currentPath, key)
			if matcher(currentPath, matchPath) {
				fmt.Println(Red+indent, key, ": ["+Reset)
			} else {
				fmt.Println(indent, key, ": [")
			}

			currentPath = currentPath[:len(currentPath)-1]
			for i, item := range v {
				if nested, ok := item.(map[string]interface{}); ok {
					fmt.Println(indent + "    " + " {")
					currentPath = append(currentPath, key)
					currentPath = append(currentPath, strconv.Itoa(i))
					printJSONHelper(nested, indent+"        ", currentPath, matchPath, matcher)
					currentPath = currentPath[:len(currentPath)-1]
					currentPath = currentPath[:len(currentPath)-1]
					fmt.Println(indent + "    " + " },")
				} else {
					fmt.Println(indent+"    ", item)
				}
			}
			fmt.Println(indent, " ]")
		default:
			currentPath = append(currentPath, key)
			if matcher(currentPath, matchPath) {
				fmt.Println(Red+indent, key, ":", value, Reset)
			} else {
				fmt.Println(indent, key, ":", value)
			}
			currentPath = currentPath[:len(currentPath)-1]
		}
	}
}

func isSubSlice(slice1, slice2 []string) bool {
	// If the second slice is empty, it's considered a sub-slice
	if len(slice2) == 0 {
		return true
	}

	// Iterate over both slices
	for i := 0; i < len(slice2) && i < len(slice1); i++ {
		// If the elements at corresponding indices don't match, return false
		if slice2[i] == "*" {
			continue
		}
		if slice1[i] != slice2[i] {
			return false
		}
	}

	// If the loop completes without returning false, return true
	return true
}

func matcher(path []string, matchPath []string) bool {
	if isSubSlice(path, matchPath) {
		return true
	}
	return false
}

func printJSON(data map[string]interface{}, indent string, matchPath []string) {
	fmt.Println(indent, "{")
	currentPath := []string{}
	printJSONHelper(data, indent+"    ", currentPath, matchPath, matcher)
	fmt.Println(indent, "}")
}

func printJsonMain() {
	jsonData := queryContent

	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("JSON Content:")
	matchPath := []string{"query", "bool", "must", "*", "term"}

	printJSON(data, "", matchPath)
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
	jsonData := queryContent

	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	rewriter.exprStack[0].Dump(indentation, data)
}

func main() {
	fmt.Println("uql driver")
	traverse()
}
