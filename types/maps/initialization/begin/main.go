// types/maps/initialization/begin/main.go
package main

import "fmt"

type author struct {
	name string
}

func main() {
	// declare a map of string keys and author values
	var authors map[string]author

	// initialize the map with make
	authors = make(map[string]author)

	// add authors to the map
	authors["tb"] = author{name: "Tom Boy"}
	authors["od"] = author{name: "Oh Doking"}

	// print the map with fmt.Printf
	fmt.Println(authors)

	// other way
	var authors2 = map[string]author{
		"tb" : {name: "Tom Boy"},
		"od" : {name: "Oh Doking"},
	}
	fmt.Println(authors2)

	// read a value from the map with a known key
	fmt.Println(authors["od"])
}
