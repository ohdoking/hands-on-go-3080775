// flow-contorl/if-else/begin/main.go
package main

import (
	"fmt"
)

// parseOddsEvens returns two slices, one with the odd numbers and one with the even numbers
func parseOddsEvens(ints []int) (odds []int, evens []int) {
	// use a for-range loop to iterate over the incoming slice
	// _ : blank identifier, to effectively discard the index that's coming back from calling the range keyword
	for _, i := range ints {
		// if i%2 == 0 {
		// setting up an inline variable right inside of the if statement and we're going to use the module operator to determine if the number is even
		// and we interrogate the even boolean
		if even := i%2 == 0; even {
			evens = append(evens, i)
		} else {
			odds = append(odds, i)
		}
	}
	// use the modulo operator to check if the number is odd or even and add it to the appropriate slice
	return
}

func main() {
	// invoke and print the results of parseOddsEvens
	odds, evens := parseOddsEvens([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(odds)
	fmt.Println(evens)
}
