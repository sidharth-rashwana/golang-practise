package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func updatePages(book *Book, pages int) {
	// your code goes here
}

func main() {

	/*
		Create 3 Book Structs with the following data:

		Book 1:
		Title: "The Great Gatsby"
		Author: "F. Scott Fitzgerald"
		Pages: 180

		Book 2
		Title: "To Kill a Mockingbird"
		Author: "Harper Lee"
		Pages: 281

		Book 3
		Title: "Pride and Prejudice"
		Author: "Jane Austen"
		Pages: 279
	*/
	var book_one = Book{Title: "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Pages:  180}
	var book_two = Book{Title: "To Kill a Mockingbird",
		Author: "Harper Lee",
		Pages:  281}
	var book_three = Book{Title: "Pride and Prejudice",
		Author: "Jane Austen",
		Pages:  279}
	// your code for creating struct objects goes here

	/*
		Update the information for Books as following:

		Book 1: Updates Page Count to 210
		Book 2: Updates Page Count to 250
		Book 3: Updates Page Count to 295

	*/
	book_one.Pages = 210
	book_two.Pages = 250
	book_three.Pages = 295

	// your code for function calls to updatePages goes here

	fmt.Println(book_one)
	fmt.Println(book_two)
	fmt.Println(book_three)
	/*
		Print all the struct objects
		fmt.Println(book)
	*/

	// your code for printing objects goes here
}
