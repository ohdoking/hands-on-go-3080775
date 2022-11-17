// challenges/types-composite/begin/main.go
package main

import "fmt"

// define an author type with a name
type author struct {
	name string
}

// define a book type with a title and author
type book struct {
	author
	title string
}

// define a library type to track a list of books
type library struct {
	books map[string][]book
}

// define addBook to add a book to the library
func (l library) addBook(b book){
	l.books[b.name] = append(l.books[b.name], b)
}

// define a lookupByAuthor function to find books by an author's name
func (l library) lookupByAuthor(authorName string) []book{
	result, ok := l.books[authorName]
	if !ok {
		// exception handling
	}
	return result
}

func main() {
	// create a new library
	l := library{books: make(map[string][]book)}

	// add 2 books to the library by the same author
	l.addBook(book{title: "title1", author : author{name : "Ohdoking"}})
	l.addBook(book{title: "title2", author : author{name : "Ohdoking"}})

	// add 1 book to the library by a different author
	l.addBook(book{title: "title3", author : author{name : "Lisa"}})

	// dump the library with spew
	//

	// lookup books by known author in the library
	books := l.lookupByAuthor("Ohdoking")

	// print out the first book's title and its author's name
	fmt.Println(books[0])

}
