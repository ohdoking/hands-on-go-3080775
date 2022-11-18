// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// handle any panics that might occur anywhere in this function
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Erorr recover! : %v\n", err)
		}
	}()

	// use os package to read the file specified as a command line argument
	file, err := os.ReadFile("challenges/data.txt")
	if err != nil {
		panic(err)
	}

	// convert the bytes to a string
	fileString := string(file[:])

	// initialize a map to store the counts
	counts := make(map[string]int)

	// use the standard library utility package that can help us split the string into words
	//

	// capture the length of the words slice
	for _, c := range fileString {
		cString := string(c)
		if val, ok := counts[cString]; ok {
			counts[cString] = val + 1
		} else {
			counts[cString] = 1
		}
	}

	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	//

	// dump the map to the console using the spew package
	spew.Dump(counts)
}
