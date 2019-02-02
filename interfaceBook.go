package main

import "fmt"

type Book struct {
	Title  string
	Author string
}

func (b *Book) String() string {
	return fmt.Sprintf("%s : %s", b.Title, b.Author)
}

func main() {

	var s fmt.Stringer = &Book{Title: "Golang", Author: "Mike"}
	println(s.String())
}
