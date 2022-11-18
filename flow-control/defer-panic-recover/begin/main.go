// flow-control/defer-panic-recover/begin/main.go
package main

import "fmt"

func cleanup(msg string) {
	fmt.Println(msg)
}
// defer is like try catch family, but it's in languages that support exception handling
// the defer mechanism can also help when it comes

// recover is a build-in function that allows us to regain control of a panicking thread
// like we have here in our example

// as a parting piece of advice, 
// don't use the defer panic recover mechanism as you would the try-catch family in other language
// it looks similar but Go encourages us to handle our erros like values hence why it doesn't support exceptions
//
func main() {
	// defer function calls
	defer cleanup("first")
	defer cleanup("second")

	fmt.Println("working in main...")
	// defer recovery
	defer func (){
		if err := recover(); err != nil {
			fmt.Println("recovered from panic:", err)
		}
	}()

	// panic
	panic("something bad happened!")
}
