// flow-control/loop-range/begin/main.go
package main

import "fmt"

func main() {
	// initialize a slice of ints
	nums := []int{100, 200, 300}

	// use for-range to iterate over the slice and print each value
	// capturing an index in a number, using range with numbs
	// again range returns both the index and the value currently being iterated on
	// so we cappture both of those with i and n
	for i, n := range nums {
		fmt.Println(i, n)
	}

	// declare a map of strings to ints
	m := map[string]int{"item":1, "item2":2, "item3":3}

	// use for-range to iterate over the map and print each key/value pair
	for k, v := range m {
		fmt.Println(k, v)
	}
}
