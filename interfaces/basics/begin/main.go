// interfaces/basics/begin/main.go
package main

import "fmt"

// define a humanoid interface with speak and walk methods returning string
type humanoid interface {
	speak()
	walk()
}
// define a person type that implements humanoid interface
type person struct {
	name string
}
func(p person) speak(){
	fmt.Println(p.name + "human speak")
}
func(p person) walk(){
	fmt.Println(p.name + "human walk")
}
// implement the Stringer interface for the person type
func(p person) String() string{
	return p.name + " is name"
}

// define a dog type that can walk but not speak
type dog struct {
	name string
}
func(d dog) walk(){
	fmt.Println(d.name + "dog walk")
}
func(d dog) speak(){
	fmt.Println(d.name + "dog can't speak")
}

func main() {
	// invoke with a person
	p := person{name:"ohdoking"}
	doHumanThings(p)

	// can we invoke with a dog?
	doHumanThings(dog{name : "toto"})

	fmt.Println(p)
}

func doHumanThings(h humanoid) {
	h.speak()
	h.walk()
}
