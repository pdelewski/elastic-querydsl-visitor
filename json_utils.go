package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// PrintJSON prints the JSON data with the given match path highlighted in red
// parameters:
// data: the JSON data to print
// currentPath: the current path in the JSON data
// indent: the initial indentation
// matchPath: the path to highlight
// returns: void
func printJSONHelper(data map[string]interface{},
	indent string,
	currentPath []string,
	matchPath []string,
	matcher func([]string, []string) bool) {
	for key, value := range data {
		switch v := value.(type) {
		case map[string]interface{}:
			currentPath = append(currentPath, key)
			matched := matcher(currentPath, matchPath)
			if matched {
				fmt.Println(Red+indent, key, ": {"+Reset)
			} else {
				fmt.Println(indent, key, ": {")
			}
			printJSONHelper(v, indent+"    ", currentPath, matchPath, matcher)
			currentPath = currentPath[:len(currentPath)-1]
			if matched {
				fmt.Println(Red+indent, " }"+Reset)
			} else {
				fmt.Println(indent, "}")
			}
		case []interface{}:
			currentPath = append(currentPath, key)
			matched := matcher(currentPath, matchPath)
			if matched {
				fmt.Println(Red+indent, key, ": ["+Reset)
			} else {
				fmt.Println(indent, key, ": [")
			}

			currentPath = currentPath[:len(currentPath)-1]
			for i, item := range v {
				if nested, ok := item.(map[string]interface{}); ok {
					fmt.Println(indent + "    " + " {")
					currentPath = append(currentPath, key, strconv.Itoa(i))
					printJSONHelper(nested, indent+"        ", currentPath, matchPath, matcher)
					currentPath = currentPath[:len(currentPath)-2]
					fmt.Println(indent + "    " + " },")
				} else {
					fmt.Println(indent+"    ", item)
				}
			}
			if matched {
				fmt.Println(Red+indent, " ]"+Reset)
			} else {
				fmt.Println(indent, " ]")
			}
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
	fmt.Println(indent, " {")
	currentPath := []string{}
	printJSONHelper(data, indent+"    ", currentPath, matchPath, matcher)
	fmt.Println(indent, " }")
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
