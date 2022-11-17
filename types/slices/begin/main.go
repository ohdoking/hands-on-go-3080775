// types/slices/begin/main.go
package main

import "fmt"

func main() {
	// declare a slice string the make function
	// make works with slices, maps and channels to allocate a size in an optional capacity
	// where applicable for these types. 
	names := make([]string, 0)

	// append 3 names to the slice
	// the appned built in is where some of the magic of slices happens
	names = append(names, "John")
	names = append(names, "Jane")
	names = append(names, "Mark")

	// print the slice
	fmt.Println(names)
	fmt.Printf("%#v", names)
	
	// slice the slice using syntax slice[low:high]
	//slices support a slice operation where we spcify a low and high bound, seperated by a colon
	// high bound is non inclusive
	// specifying one for the low and three for high bounds
	// get us jane and mary
	// [Jane Mary]
	fmt.Println(names[1:3])
	// upper bound is non-inclusive
	// leaving a high bound open gets everything from index one on  
	// [Jane Mary]
	fmt.Println(names[1:])
	// leaving a low bound unspecified get us everything from index zero up to the non-inclusive high bound
	// [John]
	fmt.Println(names[:1])
	// lastly just specifying the colon, get us everything from the slice
	// [John Jane Mary]
	fmt.Println(names[:])
}
