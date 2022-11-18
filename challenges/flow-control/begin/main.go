// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

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
	file, err := os.ReadFile(os.Args[0])
	if err != nil {
		panic(err)
	}

	// convert the bytes to a string
	fileString := string(file[:])

	// initialize a map to store the counts
	analysis := make(map[string]int)

	// use the standard library utility package that can help us split the string into words
	words := strings.Fields(fileString)

	// capture the length of the words slice
	analysis["words"] = len(words)

	// capture the length of the words slice
	for _, c := range fileString {
		if unicode.IsLetter(c) {
			analysis["letters"]++
		} else if unicode.IsNumber(c) {
			analysis["numbers"]++
		} else {
			analysis["symbols"]++
		}
	}

	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	//

	// dump the map to the console using the spew package
	spew.Dump(analysis)
}
