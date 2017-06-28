package main 

import (
	"fmt"
	"strings"
	"os"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus": {"linear algebra"},
	"linear algebra": {"calculus"}, // add cycle

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

func contains(ancestors []string, s string) bool {
	for _, v := range ancestors {
		if v == s {
			return true
		}
	}
	return false
}

func topoSort(m map[string][]string) (order []string, err error) {
	seen := make(map[string]bool)
	var visitAll func(items []string, ancestors []string)

	visitAll = func(items []string, ancestors []string) {
		for _, item := range items {
			// cannot add append(ancestors, item)
			// here, otherwise we might add several ancestors of the
			// same level l and we only need to add one ancestor for
			// each level. That is why we add append(ancestors, item)
			// in the call to visitAll

			// if duplicate found then there is a cycle
			if contains(ancestors, item) {
				// cannot write 
				// return nil, fmt.Errorf("cycle found in: %s", strings.Join(ancestors, " -> "))
				// here, because visitAll doesn't return a value! So we need to set err
				// and return nil, err in the loop key := range m {} for example
				err = fmt.Errorf("cycle found in: %s\n", strings.Join(append(ancestors, item), " -> "))
			}

			// if we didn't visit the item yet
			if !seen[item] {
				// then add it to seen and call recursively
				// visitAll passing the item we need to have
				// before being able to reach the current item
				seen[item] = true
				visitAll(m[item], append(ancestors, item))
				order = append(order, item)
			}
		}
	}

	// run the topological sort algorithm
	// seen map ensure we don't add 2 times the same course!
	for key := range m {
		if err != nil {
			return nil, err
		}
		visitAll([]string{key}, nil)
	}
	return order, nil
}

func main() {
	order, err := topoSort(prereqs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	for i, course := range order {
		// print courses in the order we need to study them
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}