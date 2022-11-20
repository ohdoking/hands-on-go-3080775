// interfaces/type-assertions/begin/main.go
package main

import "fmt"

func main() {
	// perform a type assertion
	var i any = "vv"
	printValue(i)
	
	i = 2
	printValue(i)
}

func printValue(i any){
	// perform a type assertion and handle the error
	if _, ok := i.(string) ; !ok {
		fmt.Printf("%v is not a string\n", i)
	} else{
		fmt.Println(i.(string))
	}
}
