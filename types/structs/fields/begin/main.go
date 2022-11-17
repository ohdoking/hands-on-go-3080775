// types/structs/fields/begin/main.go
package main

import "fmt"

// define a struct type for author
//
type author struct {
	first, last string
}

func main() {
	// intialize author
	//
	a := author{
		first: "Doking",
		last:  "Oh",
	}

	// print the author
	// https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/
	fmt.Printf("%#v\n", a)
}
