package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// takes all arguments
	args := os.Args[1:]

	if args[0] == "list" {
		ListAllTheBooks(BookStorage())
	} else if args[0] == "search" {
		bookName := strings.Join(args[1:], " ")
		SearchBook(bookName)
	} else {
		fmt.Println("Wrong command please try again with correct commands -> ")
	}

}

// ListAllTheBooks printing book list
func ListAllTheBooks(books []string) {
	for _, book := range books {
		fmt.Println(book)
	}
}

// SearchBook Search given prefix book name in book storage.
func SearchBook(searchBookName string) {
	books := BookStorage()
	searchedBooks := make([]string, 0)
	for _, bookName := range books {
		if strings.Contains(strings.ToLower(bookName), strings.ToLower(searchBookName)) {
			searchedBooks = append(searchedBooks, bookName)
		}
	}
	ListAllTheBooks(searchedBooks)
}

// BookStorage Store all books
// return bookList list of books which is string array.
func BookStorage() (bookList []string) {
	books := []string{"In Search of Lost Time",
		"Ulysses",
		"Don Quixote",
		"One Hundred Years of Solitude",
		"The Great Gatsby",
		"Moby Dick",
		"War and Peace",
		"Hamlet",
		"The Odyssey",
		"Madame Bovary",
		"The Divine Comedy",
		"Lolita",
		"The Brothers Karamazov",
		"Crime and Punishment",
		"The Catcher in the Rye",
		"Pride and Prejudice",
		"The Adventures of Huckleberry Finn",
		"Anna Karenina",
		"Alice's Adventures in Wonderland",
		"The Iliad",
		"To the Lighthouse",
		"Catch-22",
		"Heart of Darkness",
		"The Sound and the Fury",
		"Nineteen Eighty Four",
		"Great Expectations",
		"One Thousand and One Nights",
		"The Grapes of Wrath",
		"Absalom, Absalom!",
		"Invisible Man",
		"To Kill a Mockingbird",
		"The Trial",
		"The Red and the Black",
		"Middlemarch",
		"Gulliver's Travels",
		"Beloved",
		"Mrs. Dalloway",
		"The Stories of Anton Chekhov",
		"The Stranger",
		"Jane Eyre",
		"The Aeneid",
		"Collected Fiction",
		"The Sun Also Rises",
		"David Copperfield",
		"Tristram Shandy",
		"Leaves of Grass",
		"The Magic Mountain",
		"A Portrait of the Artist as a Young Man",
		"Midnight's Children",
		"Oedipus the King",
		"Candide",
		"The Lord of the Rings",
		"The Idiot",
		"A Passage to India",
		"The Old Man and the Sea",
		"Things Fall Apart",
		"For Whom the Bell Tolls",
		"The Complete Stories of Franz Kafka",
		"The Metamorphosis",
		"The Portrait of a Lady",
		"Frankenstein",
		"Pale Fire",
		"Antigone",
		"As I Lay Dying",
		"The Color Purple"}
	return books
}
