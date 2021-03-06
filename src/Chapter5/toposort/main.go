package main 

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus": {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			// if we didn't visit the item yet
			if !seen[item] {
				// then add it to seen and call recursively
				// visitAll passing the item we need to have
				// before being able to reach the current item
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	// build a slice of all the courses
	sort.String(keys)

	// run the topological sort algorithm
	visitAll(keys)
	return order
}

func main() {
	for i, course := range topoSort(prereqs) {
		// print courses in the order we need to study them
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}