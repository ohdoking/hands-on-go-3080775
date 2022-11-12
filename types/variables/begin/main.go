// types/variables/begin/main.go
package main

import "fmt"

// declare package-level variables of type int
//

// declare package-level variables of type bool and override the default values (also known as "zero")
// we can declare with value then go will recognize type based on value
var x, y, z = true, "blue", 15

func main() {

	// print ints
	// 

	// override the default value of a package-level variable
	// d = 1_000_000
	// fmt.Printf("d: %d\n", d)

	// declare and initialize variables of similar names as booleans but of local scope
	// same name with package level didn't interfere with each other.
	x, y, z := 0, 1.25, "hello"
	fmt.Println("x, y, z:", x, y, z)

	// print the package-level variables x, y, and z
	printXYZ()
}


func printXYZ(){
	fmt.Println("x, y, z:", x, y, z)
}