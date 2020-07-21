package main

import (
	"fmt"
	m "./models"
)

func main() {
	var author m.Author = m.Author{"Author one","author1@gmail.com"}
	var Book m.Books = m.Books{2,"Book one","4a4a5a5a58a45as",author}
	fmt.Println(Book.Author.Email)
}