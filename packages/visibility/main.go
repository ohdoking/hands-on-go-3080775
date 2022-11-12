// packages/visibility/main.go
package main

import "fmt"
// function with lowercase makes it unusable outside of the package
func main() {
	fmt.Println("Hello Gopher!")
	// fmt.newPrinter()
}
