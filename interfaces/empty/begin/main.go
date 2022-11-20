// interfaces/empty/begin/main.go
package main

import "fmt"

// define whatAmI which takes in an argument of any type and returns its underlying value's type
// empty interface denoted by the keyword interface with a set of empty braces,
// means there are no type information for the compiler to check against
// but that doesn`t mean there`s no underlying type information
func whatAmI(val interface{}) string{
	return fmt.Sprintf("%v, %T", val, val)
} 
func main() {
	// invoke whatAmI with an int
	fmt.Println(whatAmI(1))

	// invoke whatAmI with a string
	fmt.Println(whatAmI("hello"))
}
