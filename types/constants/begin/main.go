// types/constants/begin/main.go
package main

// idiomatic Go nudges us towards using lower case for const 

// declare a constant representing pi
const pi - 3.14159
// declare a rune constant for the letter 'a'
// rune type is actullay an alias for int32 type
const a rune = 'a'

func main() {
	// print the value and types of pi and a
	// fmt.Printf("pi: %v - %T\n", pi, pi)
	// fmt.Printf("a: %c - %T\n", a, a)

	// use unicode package to confirm that the rune is a letter
	// unicode.IsLetter(a)
}
